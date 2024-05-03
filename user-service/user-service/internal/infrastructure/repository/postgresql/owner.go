package postgresql

import (
	"context"
	"fmt"
	"time"
	"user-service/internal/entity"
	"user-service/internal/pkg/postgres"

	"github.com/Masterminds/squirrel"
)

const (
	ownersTableName      = "owners"
	ownersServiceName    = "ownerService"
	ownersSpanRepoPrefix = "ownersRepo"
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
			"created_at",
			"updated_at",
		).From(p.tableName)
}

func (p ownersRepo) CreateOwner(ctx context.Context, owner *entity.Owner) error {
	data := map[string]any{
		"id":           owner.Id,
		"full_name":    owner.FullName,
		"company_name": owner.CompanyName,
		"email":        owner.Email,
		"password":     owner.Password,
		"avatar":       owner.Avatar,
		"tax":          owner.Tax,
		"created_at":   owner.CreatedAt,
		"updated_at":   owner.UpdatedAt,
	}
	query, args, err := p.db.Sq.Builder.Insert(p.tableName).SetMap(data).ToSql()
	if err != nil {
		return p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "create"))
	}

	_, err = p.db.Exec(ctx, query, args...)
	if err != nil {
		return p.db.Error(err)
	}

	return nil
}

func (p ownersRepo) GetOwner(ctx context.Context, params map[string]string) (*entity.Owner, error) {
	var (
		owner entity.Owner
	)

	queryBuilder := p.ownersSelectQueryPrefix()

	for key, value := range params {
		if key == "id" {
			queryBuilder = queryBuilder.Where(p.db.Sq.Equal(key, value))
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
		&owner.CreatedAt,
		&owner.UpdatedAt,
	); err != nil {
		return nil, p.db.Error(err)
	}

	return &owner, nil
}

func (p ownersRepo) UpdateOwner(ctx context.Context, owners *entity.Owner) error {
	clauses := map[string]any{
		"full_name":    owners.FullName,
		"company_name": owners.CompanyName,
		"email":        owners.Email,
		"password":     owners.Password,
		"avatar":       owners.Avatar,
		"tax":          owners.Tax,
		"updated_at":   owners.UpdatedAt,
	}
	sqlStr, args, err := p.db.Sq.Builder.
		Update(p.tableName).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", owners.Id)).
		ToSql()
	if err != nil {
		return p.db.ErrSQLBuild(err, p.tableName+" update")
	}

	commandTag, err := p.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return p.db.Error(err)
	}

	if commandTag.RowsAffected() == 0 {
		return p.db.Error(fmt.Errorf("no sql rows"))
	}

	return nil
}

// For soft delete
func (p ownersRepo) DeleteOwner(ctx context.Context, guid string) error {
	data := map[string]any{
		"deleted_at": time.Now(),
	}

	query, args, err := p.db.Sq.Builder.
		Update(p.tableName).
		SetMap(data).
		Where(p.db.Sq.Equal("id", guid)).
		ToSql()
	if err != nil {
		return p.db.ErrSQLBuild(err, p.tableName+" delete")
	}

	commandTag, err := p.db.Exec(ctx, query, args...)
	if err != nil {
		return p.db.Error(err)
	}

	if commandTag.RowsAffected() == 0 {
		return p.db.Error(fmt.Errorf("no sql rows"))
	}

	return nil
}

func (p ownersRepo) ListOwner(ctx context.Context, limit uint64, offset uint64, filter map[string]string) ([]*entity.Owner, error) {
	var (
		owners []*entity.Owner
	)
	queryBuilder := p.ownersSelectQueryPrefix()

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(limit).Offset(offset)
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
	owners = make([]*entity.Owner, 0)
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
			&owner.CreatedAt,
			&owner.UpdatedAt,
		); err != nil {
			return nil, p.db.Error(err)
		}
		owners = append(owners, &owner)
	}

	return owners, nil
}

func (p ownersRepo) CheckFieldOwner(ctx context.Context, field, value string) (bool, error) {
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

	// query, args, err := p.db.Sq.Builder.
	// 	Select(p.tableName).
	// 	Where(p.db.Sq.Equal(field, value)).
	// 	ToSql()
	// if err != nil {
	// 	fmt.Println("\x1b[32m Error 1\x1b[0m")
	// 	return true, p.db.ErrSQLBuild(err, p.tableName+" CheckFieldOwner")
	// }

	// commandTag, err := p.db.Exec(ctx, query, args...)
	// if err != nil {
	// 	fmt.Println("\x1b[32m Error 2\x1b[0m")
	// 	return true, p.db.Error(err)
	// }

	// if commandTag.RowsAffected() == 0 {
	// 	fmt.Println("\x1b[32m Error 3\x1b[0m")
	// 	return false, nil
	// }

	// return true, nil
}
