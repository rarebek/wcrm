package postgresql

import (
	"context"
	"fmt"
	"time"
	"user-service/internal/entity"
	"user-service/internal/pkg/otlp"
	"user-service/internal/pkg/postgres"

	"github.com/Masterminds/squirrel"
	"github.com/k0kubun/pp"
)

const (
	ownersTableName      = "owners"
	ownersServiceName    = "ownerService"
	ownersSpanRepoPrefix = "ownerRepo"
)

type ownersRepo struct {
	tableName string
	db        *postgres.PostgresDB
}

func NewOwnersRepo(db *postgres.PostgresDB) *ownersRepo {
	return &ownersRepo{
		tableName: ownersTableName,
		db:        db,
	}
}

func (p *ownersRepo) ownersSelectQueryPrefix() squirrel.SelectBuilder {
	return p.db.Sq.Builder.
		Select(
			"id",
			"full_name",
			"company_name",
			"email",
			"password",
			"avatar",
			"tax",
			"refresh_token",
			"created_at",
			"updated_at",
		).From(p.tableName)
}

func (p ownersRepo) Create(ctx context.Context, owner *entity.Owner) (*entity.Owner, error) {
	ctx, span := otlp.Start(ctx, ownersServiceName, ownersSpanRepoPrefix+"Create")
	defer span.End()

	data := map[string]any{
		"id":            owner.Id,
		"full_name":     owner.FullName,
		"company_name":  owner.CompanyName,
		"email":         owner.Email,
		"password":      owner.Password,
		"avatar":        owner.Avatar,
		"tax":           owner.Tax,
		"refresh_token": owner.RefreshToken,
		"created_at":    owner.CreatedAt,
		"updated_at":    owner.UpdatedAt,
	}
	query, args, err := p.db.Sq.Builder.Insert(p.tableName).SetMap(data).ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "create"))
	}
	query += "RETURNING id, full_name, company_name, email, password, avatar, tax, refresh_token, created_at, updated_at"

	row := p.db.QueryRow(ctx, query, args...)

	if err = row.Scan(&owner.Id, &owner.FullName, &owner.CompanyName, &owner.Email, &owner.Password, &owner.Avatar, &owner.Tax, &owner.RefreshToken, &owner.CreatedAt, &owner.UpdatedAt); err != nil {
		return nil, err
	}

	return owner, nil
}

func (p ownersRepo) Get(ctx context.Context, params map[string]string) (*entity.Owner, error) {
	ctx, span := otlp.Start(ctx, ownersServiceName, ownersSpanRepoPrefix+"Get")
	defer span.End()
	var (
		owner entity.Owner
	)

	queryBuilder := p.ownersSelectQueryPrefix()

	for key, value := range params {
		if key == "id" {
			queryBuilder = queryBuilder.Where(p.db.Sq.Equal(key, value)).Where("deleted_at IS NULL")
		}
		if key == "email" {
			queryBuilder = queryBuilder.Where(p.db.Sq.Equal(key, value)).Where("deleted_at IS NULL")
		}

	}
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "get"))
	}
	if err = p.db.QueryRow(ctx, query, args...).Scan(
		&owner.Id,
		&owner.FullName,
		&owner.CompanyName,
		&owner.Email,
		&owner.Password,
		&owner.Avatar,
		&owner.Tax,
		&owner.RefreshToken,
		&owner.CreatedAt,
		&owner.UpdatedAt,
	); err != nil {
		return nil, p.db.Error(err)
	}

	return &owner, nil
}

func (p ownersRepo) Update(ctx context.Context, owners *entity.Owner) (*entity.Owner, error) {
	ctx, span := otlp.Start(ctx, ownersServiceName, ownersSpanRepoPrefix+"Update")
	defer span.End()
	clauses := map[string]any{
		"full_name":     owners.FullName,
		"company_name":  owners.CompanyName,
		"email":         owners.Email,
		"password":      owners.Password,
		"avatar":        owners.Avatar,
		"tax":           owners.Tax,
		"refresh_token": owners.RefreshToken,
		"updated_at":    owners.UpdatedAt,
	}
	sqlStr, args, err := p.db.Sq.Builder.
		Update(p.tableName).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", owners.Id)).
		Where("deleted_at IS NULL").
		ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, p.tableName+" update")
	}

	sqlStr += " RETURNING id, full_name, company_name, email, password, avatar, tax, refresh_token, created_at, updated_at"

	row := p.db.QueryRow(ctx, sqlStr, args...)
	var resOwner entity.Owner
	if err = row.Scan(&resOwner.Id, &resOwner.FullName, &resOwner.CompanyName, &resOwner.Email, &resOwner.Password, &resOwner.Avatar, &resOwner.Tax, &resOwner.RefreshToken, &resOwner.CreatedAt, &resOwner.UpdatedAt); err != nil {
		return nil, err
	}
	return &resOwner, nil
}

// For soft delete
func (p ownersRepo) Delete(ctx context.Context, guid string) (*entity.CheckResponse, error) {
	ctx, span := otlp.Start(ctx, ownersServiceName, ownersSpanRepoPrefix+"Delete")
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

func (p ownersRepo) List(ctx context.Context, limit uint64, offset uint64, filter map[string]string) (*entity.AllOwners, error) {
	ctx, span := otlp.Start(ctx, ownersServiceName, ownersSpanRepoPrefix+"List")
	defer span.End()

	queryBuilder := p.ownersSelectQueryPrefix()

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(limit).Offset(offset).Where("deleted_at IS NULL")
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
	query = `SELECT COUNT(*) FROM owners WHERE deleted_at IS NULL`
	err = p.db.QueryRow(ctx, query).Scan(&count)
	if err != nil {
		return nil, p.db.Error(err)
	}

	pp.Println(count)

	owners := []entity.Owner{}
	for rows.Next() {
		var owner entity.Owner
		if err = rows.Scan(
			&owner.Id,
			&owner.FullName,
			&owner.CompanyName,
			&owner.Email,
			&owner.Password,
			&owner.Avatar,
			&owner.Tax,
			&owner.RefreshToken,
			&owner.CreatedAt,
			&owner.UpdatedAt,
		); err != nil {
			return nil, p.db.Error(err)
		}
		owners = append(owners, owner)
	}

	return &entity.AllOwners{
		Owners: owners,
		Count:  count,
	}, nil
}

func (p ownersRepo) CheckField(ctx context.Context, field, value string) (bool, error) {
	query := fmt.Sprintf(
		`SELECT count(1) 
		FROM owners WHERE %s = $1 
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
