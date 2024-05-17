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
	productTableName      = "products"
	productServiceName    = "productService"
	productSpanRepoPrefix = "productRepo"
)

type productRepo struct {
	tableName string
	db        *postgres.PostgresDB
}

func NewProductRepo(db *postgres.PostgresDB) *productRepo {
	return &productRepo{
		tableName: productTableName,
		db:        db,
	}
}
func (p *productRepo) productSelectQueryPrefix() squirrel.SelectBuilder {
	return p.db.Sq.Builder.
		Select(
			"id",
			"owner_id",
			"title",
			"description",
			"price",
			"discount",
			"picture",
			"created_at",
			"updated_at",
		).From(p.tableName)
}

func (p productRepo) CreateProduct(ctx context.Context, product *entity.ProductWithCategoryId) (*entity.Product, error) {
	ctx, span := otlp.Start(ctx, productServiceName, productSpanRepoPrefix+"Create")

	defer span.End()
	data := map[string]any{
		"id":          product.Id,
		"title":       product.Title,
		"owner_id":    product.OwnerId,
		"description": product.Description,
		"price":       product.Price,
		"discount":    product.Discount,
		"picture":     product.Picture,
		"created_at":  product.CreatedAt,
		"updated_at":  product.UpdatedAt,
	}
	query, args, err := p.db.Sq.Builder.Insert(p.tableName).SetMap(data).ToSql()
	if err != nil {
		return &entity.Product{}, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "create"))
	}

	query += " RETURNING id, owner_id, title, description, price, discount, picture, created_at, updated_at"

	row := p.db.QueryRow(ctx, query, args...)

	var createdProduct entity.Product

	err = row.Scan(&createdProduct.Id,
		&createdProduct.OwnerId,
		&createdProduct.Title,
		&createdProduct.Description,
		&createdProduct.Price,
		&createdProduct.Discount,
		&createdProduct.Picture,
		&createdProduct.CreatedAt,
		&createdProduct.UpdatedAt,
	)

	if err != nil {
		log.Println(err)

		return &entity.Product{}, err
	}

	data = map[string]any{
		"product_id":  createdProduct.Id,
		"category_id": product.CategoryId,
		"created_at":  product.CreatedAt,
		"updated_at":  product.UpdatedAt,
	}
	query, args, err = p.db.Sq.Builder.Insert("categories_products").SetMap(data).ToSql()

	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", "categories_products", "create"))
	}

	_, err = p.db.Exec(ctx, query, args...)

	if err != nil {
		// (err)
		log.Println(err)
		return nil, err
	}

	return &createdProduct, nil
}
func (p productRepo) GetProduct(ctx context.Context, params map[string]string) (*entity.Product, error) {
	ctx, span := otlp.Start(ctx, productServiceName, productSpanRepoPrefix+"Get")
	defer span.End()

	var (
		product entity.Product
	)

	queryBuilder := p.productSelectQueryPrefix()

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
		&product.Id,
		&product.OwnerId,
		&product.Title,
		&product.Description,
		&product.Price,
		&product.Discount,
		&product.Picture,
		&product.CreatedAt,
		&product.UpdatedAt,
	); err != nil {
		log.Println(err)
		return nil, p.db.Error(err)
	}

	return &product, nil
}
func (p productRepo) ListProduct(ctx context.Context, limit, offset uint64, filter map[string]string) (*entity.AllProduct, error) {
	ctx, span := otlp.Start(ctx, productServiceName, productSpanRepoPrefix+"List")
	defer span.End()

	queryBuilder := p.productSelectQueryPrefix()

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
	query = `SELECT COUNT(*) FROM products where owner_id = $1`
	err = p.db.QueryRow(ctx, query, filter["owner_id"]).Scan(&count)
	if err != nil {
		return nil, p.db.Error(err)
	}

	products := []entity.Product{}
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(
			&product.Id,
			&product.OwnerId,
			&product.Title,
			&product.Description,
			&product.Price,
			&product.Discount,
			&product.Picture,
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			return nil, p.db.Error(err)
		}
		products = append(products, product)
	}

	return &entity.AllProduct{
		Products: products,
		Count:    count,
	}, nil
}
func (p productRepo) UpdateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	ctx, span := otlp.Start(ctx, productServiceName, productSpanRepoPrefix+"Update")
	defer span.End()

	clauses := map[string]any{
		"owner_id":    product.OwnerId,
		"title":       product.Title,
		"description": product.Description,
		"price":       product.Price,
		"discount":    product.Discount,
		"picture":     product.Picture,
		"updated_at":  product.UpdatedAt,
	}

	query, args, err := p.db.Sq.Builder.
		Update(p.tableName).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", product.Id)).
		ToSql()
	if err != nil {
		return &entity.Product{}, p.db.ErrSQLBuild(err, p.tableName+" update")
	}

	query += " RETURNING id, owner_id, title, description, price, discount, picture, created_at, updated_at"

	row := p.db.QueryRow(ctx, query, args...)

	var updatedProduct entity.Product

	err = row.Scan(&updatedProduct.Id,
		&updatedProduct.OwnerId,
		&updatedProduct.Title,
		&updatedProduct.Description,
		&updatedProduct.Price,
		&updatedProduct.Discount,
		&updatedProduct.Picture,
		&updatedProduct.CreatedAt,
		&updatedProduct.UpdatedAt)

	if err != nil {
		return &entity.Product{}, err
	}
	return &updatedProduct, nil
}
func (p productRepo) DeleteProduct(ctx context.Context, id string) (*entity.CheckResponse, error) {
	ctx, span := otlp.Start(ctx, productServiceName, productSpanRepoPrefix+"Delete")
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
func (p productRepo) SearchProduct(ctx context.Context, page, offset int64, title string, ownerId string) (*entity.AllProduct, error) {

	ctx, span := otlp.Start(ctx, productServiceName, productSpanRepoPrefix+"SearchProduct")
	defer span.End()

	var (
		products []entity.Product
	)

	queryBuilder := p.productSelectQueryPrefix()

	query, _, err := queryBuilder.Where("title ilike " + "'%" + title + "%' AND owner_id = $3 LIMIT $1 OFFSET $2").ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "search_product"))
	}

	rows, err := p.db.Query(ctx, query, page, offset, ownerId)
	if err != nil {
		return nil, p.db.Error(err)
	}
	defer rows.Close()

	var count int
	query = `SELECT COUNT(*) FROM products where owner_id = $1`
	err = p.db.QueryRow(ctx, query, ownerId).Scan(&count)
	if err != nil {
		return nil, p.db.Error(err)
	}

	products = []entity.Product{}

	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(
			&product.Id,
			&product.OwnerId,
			&product.Title,
			&product.Description,
			&product.Price,
			&product.Discount,
			&product.Picture,
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			return nil, p.db.Error(err)
		}
		products = append(products, product)
	}

	return &entity.AllProduct{
		Products: products,
		Count:    count,
	}, nil
}
func (p productRepo) GetAllProductByCategoryId(ctx context.Context, limit, offset uint64, CategoryId string) (*entity.AllProduct, error) {
	ctx, span := otlp.Start(ctx, productServiceName, productSpanRepoPrefix+"GetAllProductByCategoryId")
	defer span.End()

	queryBuilder := p.db.Sq.Builder.
		Select(
			"product_id",
		).From("categories_products")

	if limit != 0 {
		queryBuilder = queryBuilder.Where(p.db.Sq.Equal("category_id", CategoryId)).Limit(limit).Offset(offset)
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "GetAllProductByCategoryId"))
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, p.db.Error(err)
	}
	defer rows.Close()

	products := []entity.Product{}

	for rows.Next() {
		var req string
		if err := rows.Scan(
			&req,
		); err != nil {
			return nil, p.db.Error(err)
		}
		filter := map[string]string{"id": req}
		product, err := p.GetProduct(ctx, filter)
		if err != nil {
			return nil, p.db.Error(err)
		}
		products = append(products, *product)
	}

	var count int
	query = `SELECT COUNT(*) FROM categories_products WHERE category_id = $1`
	err = p.db.QueryRow(ctx, query, CategoryId).Scan(&count)
	if err != nil {
		return nil, p.db.Error(err)
	}

	return &entity.AllProduct{
		Products: products,
		Count:    count,
	}, nil
}
