package postgresql

import (
	"context"
	"fmt"
	"time"
	"user-service/internal/entity"
	"user-service/internal/pkg/otlp"
	"user-service/internal/pkg/postgres"

	"github.com/Masterminds/squirrel"
)

const (
	workersTableName      = "workers"
	workersServiceName    = "workerService"
	workersSpanRepoPrefix = "workersRepo"
)

type workersRepo struct {
	tableName string
	db        *postgres.PostgresDB
}

func NewWorkersRepo(db *postgres.PostgresDB) *workersRepo {
	return &workersRepo{
		tableName: workersTableName,
		db:        db,
	}
}

func (p *workersRepo) workersSelectQueryPrefix() squirrel.SelectBuilder {
	return p.db.Sq.Builder.
		Select(
			"id",
			"full_name",
			"login_key",
			"password",
			"owner_id",
			"created_at",
			"updated_at",
		).From(p.tableName)
}

func (p workersRepo) Create(ctx context.Context, worker *entity.Worker) (*entity.Worker, error) {
	ctx, span := otlp.Start(ctx, workersServiceName, workersSpanRepoPrefix+"Create")
	defer span.End()

	data := map[string]any{
		"id":         worker.Id,
		"full_name":  worker.FullName,
		"login_key":  worker.LoginKey,
		"password":   worker.Password,
		"owner_id":   worker.OwnerId,
		"created_at": worker.CreatedAt,
		"updated_at": worker.UpdatedAt,
	}
	query, args, err := p.db.Sq.Builder.Insert(p.tableName).SetMap(data).ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "create"))
	}

	query += "RETURNING id, full_name, login_key, password, owner_id, created_at, updated_at"

	row := p.db.QueryRow(ctx, query, args...)

	if err = row.Scan(&worker.Id, &worker.FullName, &worker.LoginKey, &worker.Password, &worker.OwnerId, &worker.CreatedAt, &worker.UpdatedAt); err != nil {
		return nil, err
	}

	return worker, nil
}

func (p workersRepo) Get(ctx context.Context, params map[string]string) (*entity.Worker, error) {
	ctx, span := otlp.Start(ctx, workersServiceName, workersSpanRepoPrefix+"Create")
	defer span.End()

	var (
		worker entity.Worker
	)

	queryBuilder := p.workersSelectQueryPrefix()

	for key, value := range params {
		if key == "id" {
			queryBuilder = queryBuilder.Where(p.db.Sq.Equal(key, value)).Where("deleted_at IS NULL")
		}
		if key == "login_key" {
			queryBuilder = queryBuilder.Where(p.db.Sq.Equal(key, value)).Where("deleted_at IS NULL")
		}
	}
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "get"))
	}
	if err = p.db.QueryRow(ctx, query, args...).Scan(
		&worker.Id,
		&worker.FullName,
		&worker.LoginKey,
		&worker.Password,
		&worker.OwnerId,
		&worker.CreatedAt,
		&worker.UpdatedAt,
	); err != nil {
		return nil, p.db.Error(err)
	}

	return &worker, nil
}

func (p workersRepo) Update(ctx context.Context, workers *entity.Worker) (*entity.Worker, error) {
	ctx, span := otlp.Start(ctx, workersServiceName, workersSpanRepoPrefix+"Create")
	defer span.End()
	clauses := map[string]any{
		"full_name":  workers.FullName,
		"login_key":  workers.LoginKey,
		"password":   workers.Password,
		"owner_id":   workers.OwnerId,
		"updated_at": workers.UpdatedAt,
	}
	sqlStr, args, err := p.db.Sq.Builder.
		Update(p.tableName).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", workers.Id)).
		Where("deleted_at IS NULL").
		ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, p.tableName+" update")
	}

	sqlStr += " RETURNING id, full_name, login_key, password, owner_id, created_at, updated_at"

	row := p.db.QueryRow(ctx, sqlStr, args...)
	var resWorker entity.Worker
	if err = row.Scan(&resWorker.Id, &resWorker.FullName, &resWorker.LoginKey, &resWorker.Password, &resWorker.OwnerId, &resWorker.CreatedAt, &resWorker.UpdatedAt); err != nil {
		return nil, err
	}

	return &resWorker, nil
}

// For soft delete
func (p workersRepo) Delete(ctx context.Context, guid string) (*entity.CheckResponse, error) {
	ctx, span := otlp.Start(ctx, workersServiceName, workersSpanRepoPrefix+"Create")
	defer span.End()
	data := map[string]any{
		"deleted_at": time.Now(),
	}

	query, args, err := p.db.Sq.Builder.
		Update(p.tableName).
		SetMap(data).
		Where(p.db.Sq.Equal("id", guid)).
		Where("deleted_at IS NULL").
		ToSql()
	if err != nil {
		return &entity.CheckResponse{Check: false}, nil
	}

	commandTag, err := p.db.Exec(ctx, query, args...)
	if err != nil {
		return &entity.CheckResponse{Check: false}, nil
	}

	if commandTag.RowsAffected() == 0 {
		return &entity.CheckResponse{Check: false}, nil
	}

	return &entity.CheckResponse{Check: true}, nil
}

func (p workersRepo) List(ctx context.Context, limit uint64, offset uint64, filter map[string]string) (*entity.AllWorker, error) {
	ctx, span := otlp.Start(ctx, workersServiceName, workersSpanRepoPrefix+"Create")
	defer span.End()
	fmt.Println(filter["owner_id"])

	queryBuilder := p.workersSelectQueryPrefix()

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(limit).Offset(offset).Where(
			squirrel.And{
				squirrel.Eq{
					"deleted_at": nil,
				},
				squirrel.Eq{
					"owner_id": filter["owner_id"],
				},
			},
		)
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "list"))
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, p.db.Error(err)
	}
	defer rows.Close()

	var count int
	query = `SELECT COUNT(*) FROM workers WHERE owner_id = $1`
	err = p.db.QueryRow(ctx, query, filter["owner_id"]).Scan(&count)
	if err != nil {
		return nil, p.db.Error(err)
	}

	workers := []entity.Worker{}

	for rows.Next() {
		var worker entity.Worker
		if err = rows.Scan(
			&worker.Id,
			&worker.FullName,
			&worker.LoginKey,
			&worker.Password,
			&worker.OwnerId,
			&worker.CreatedAt,
			&worker.UpdatedAt,
		); err != nil {
			return nil, p.db.Error(err)
		}
		workers = append(workers, worker)
	}

	return &entity.AllWorker{
		Workers: workers,
		Count:   count,
	}, nil
}

func (p workersRepo) CheckField(ctx context.Context, field, value string) (bool, error) {
	query := fmt.Sprintf(
		`SELECT count(1) 
		FROM workers WHERE %s = $1 
		AND deleted_at IS NULL`, field)

	var isExists int

	row := p.db.QueryRow(ctx, query, value)
	if err := row.Scan(&isExists); err != nil {
		return true, err
	}

	if isExists == 0 {
		return false, nil
	}

	return true, nil
}
