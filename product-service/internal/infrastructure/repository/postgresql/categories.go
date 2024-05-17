package postgresql

import (
	"context"
	"fmt"
	"log"
	"wcrm/product-service/internal/entity"

	"wcrm/product-service/internal/pkg/otlp"
	"wcrm/product-service/internal/pkg/postgres"

	"github.com/Masterminds/squirrel"
)

const (
	categoryTableName      = "categories"
	categoryServiceName    = "categoryService"
	categorySpanRepoPrefix = "categoryRepo"
)

type categoryRepo struct {
	tableName string
	db        *postgres.PostgresDB
}

func NewCategoryRepo(db *postgres.PostgresDB) *categoryRepo {
	return &categoryRepo{
		tableName: categoryTableName,
		db:        db,
	}
}

func (p *categoryRepo) categorySelectQueryPrefix() squirrel.SelectBuilder {
	return p.db.Sq.Builder.
		Select(
			"id",
			"owner_id",
			"name",
			"image",
			"created_at",
			"updated_at",
		).From(p.tableName)
}
func (p categoryRepo) CreateCategory(ctx context.Context, category *entity.Category) (*entity.Category, error) {
	ctx, span := otlp.Start(ctx, categoryServiceName, categorySpanRepoPrefix+"Create")
	defer span.End()

	data := map[string]any{
		"id":         category.Id,
		"name":       category.Name,
		"owner_id":   category.OwnerId,
		"image":      category.Image,
		"created_at": category.CreatedAt,
		"updated_at": category.UpdatedAt,
	}
	query, args, err := p.db.Sq.Builder.Insert(p.tableName).SetMap(data).ToSql()
	if err != nil {
		return &entity.Category{}, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "create"))
	}

	query += " RETURNING id, owner_id, name, image, created_at, updated_at"

	row := p.db.QueryRow(ctx, query, args...)

	var createdCategory entity.Category

	err = row.Scan(&createdCategory.Id,
		&createdCategory.OwnerId,
		&createdCategory.Name,
		&createdCategory.Image,
		&createdCategory.CreatedAt,
		&createdCategory.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &createdCategory, nil

}
func (p categoryRepo) GetCategory(ctx context.Context, params map[string]string) (*entity.Category, error) {
	ctx, span := otlp.Start(ctx, categoryServiceName, categorySpanRepoPrefix+"Get")
	defer span.End()

	var (
		category entity.Category
	)

	queryBuilder := p.categorySelectQueryPrefix()
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
		&category.Id,
		&category.OwnerId,
		&category.Name,
		&category.Image,
		&category.CreatedAt,
		&category.UpdatedAt,
	); err != nil {
		log.Println(err)
		return nil, p.db.Error(err)
	}

	return &category, nil
}
func (p categoryRepo) ListCategory(ctx context.Context, limit, offset uint64, filter map[string]string) (*entity.AllCategory, error) {
	ctx, span := otlp.Start(ctx, categoryServiceName, categorySpanRepoPrefix+"List")
	defer span.End()

	queryBuilder := p.categorySelectQueryPrefix()

	if limit != 0 {
		queryBuilder = queryBuilder.Where(p.db.Sq.Equal("owner_id", filter["owner_id"])).Limit(limit).Offset(offset)
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
	query = `SELECT COUNT(*) FROM categories where owner_id = $1`
	err = p.db.QueryRow(ctx, query, filter["owner_id"]).Scan(&count)
	if err != nil {
		return nil, p.db.Error(err)
	}

	categories := []entity.Category{}
	for rows.Next() {
		var category entity.Category
		if err := rows.Scan(
			&category.Id,
			&category.OwnerId,
			&category.Name,
			&category.Image,
			&category.CreatedAt,
			&category.UpdatedAt,
		); err != nil {
			return nil, p.db.Error(err)
		}
		categories = append(categories, category)
	}

	return &entity.AllCategory{
		Categories: categories,
		Count:      count,
	}, nil
}

func (p categoryRepo) UpdateCategory(ctx context.Context, category *entity.Category) (*entity.Category, error) {
	ctx, span := otlp.Start(ctx, categoryServiceName, categorySpanRepoPrefix+"Update")
	defer span.End()


	clauses := map[string]any{
		"name":       category.Name,
		"image":      category.Image,
		"updated_at": category.UpdatedAt,
	}

	query, args, err := p.db.Sq.Builder.
		Update(p.tableName).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", category.Id)).
		ToSql()
	if err != nil {
		return &entity.Category{}, p.db.ErrSQLBuild(err, p.tableName+" update")
	}

	query += " RETURNING id, owner_id, name, image, created_at, updated_at"

	row := p.db.QueryRow(ctx, query, args...)

	var updatedCategory entity.Category

	err = row.Scan(&updatedCategory.Id,
		&updatedCategory.OwnerId,
		&updatedCategory.Name,
		&updatedCategory.Image,
		&updatedCategory.CreatedAt,
		&updatedCategory.UpdatedAt)

	if err != nil {
		return &entity.Category{}, err
	}
	return &updatedCategory, nil
}
func (p categoryRepo) DeleteCategory(ctx context.Context, id string) (*entity.CheckResponse, error) {
	ctx, span := otlp.Start(ctx, categoryServiceName, categorySpanRepoPrefix+"Delete")
	defer span.End()

	sqlStr, args, err := p.db.Sq.Builder.
		Delete(p.tableName).
		Where(p.db.Sq.Equal("id", id)).
		ToSql()
	if err != nil {
		return &entity.CheckResponse{Check: false}, p.db.ErrSQLBuild(err, p.tableName+" delete")
	}

	commandTag, err := p.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return &entity.CheckResponse{Check: false}, p.db.Error(err)
	}

	if commandTag.RowsAffected() == 0 {
		return &entity.CheckResponse{Check: false}, p.db.Error(fmt.Errorf("no sql rows"))
	}

	return &entity.CheckResponse{Check: true}, nil
}
