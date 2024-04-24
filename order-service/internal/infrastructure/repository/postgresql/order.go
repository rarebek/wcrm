package postgresql

import (
	"context"
	"fmt"

	// "order-service/genproto/Order"
	"order-service/internal/entity"

	// "order-service/internal/pkg/otlp"
	"order-service/internal/pkg/postgres"

	"github.com/Masterminds/squirrel"
	"github.com/k0kubun/pp"
)

const (
	orderTableName      = "orders"
	orderServiceName    = "orderService"
	orderSpanRepoPrefix = "orderRepo"
)

type orderRepo struct {
	tableName string
	db        *postgres.PostgresDB
}

func NewOrderRepo(db *postgres.PostgresDB) *orderRepo {
	return &orderRepo{
		tableName: orderTableName,
		db:        db,
	}
}

func (p *orderRepo) orderSelectQueryPrefix() squirrel.SelectBuilder {
	return p.db.Sq.Builder.
		Select(
			"id",
			"worker_id",
			"product_id",
			"tax",
			"discount",
			"total_price",
			"created_at",
			"updated_at",
		).From(p.tableName)
}

func (p orderRepo) CreateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error) {
	// ctx, span := otlp.Start(ctx, userServiceName, userSpanRepoPrefix+"Create")
	// defer span.End()
	data := map[string]any{
		"id":          order.Id,
		"worker_id":   order.WorkerId,
		"product_id":  order.ProductId,
		"tax":         order.Tax,
		"discount":    order.Discount,
		"total_price": order.TotalPrice,
		"created_at":  order.CreatedAt,
		"updated_at":  order.UpdatedAt,
	}
	query, args, err := p.db.Sq.Builder.Insert(p.tableName).SetMap(data).ToSql()
	if err != nil {
		return &entity.Order{}, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "create"))
	}

	query += "RETURNING id, worker_id, product_id, tax, discount, total_price, created_at, updated_at"

	row := p.db.QueryRow(ctx, query, args...)

	var created entity.Order

	err = row.Scan(
		&created.Id,
		&created.WorkerId,
		&created.ProductId,
		&created.Tax,
		&created.Discount,
		&created.TotalPrice,
		&created.CreatedAt,
		&created.UpdatedAt)

	if err != nil {
		return &entity.Order{}, err
	}
	return &created, nil
}

func (p orderRepo) GetOrder(ctx context.Context, params map[string]int64) (*entity.Order, error) {
	// ctx, span := otlp.Start(ctx, userServiceName, userSpanRepoPrefix+"Get")
	// defer span.End()

	var (
		order entity.Order
	)

	queryBuilder := p.orderSelectQueryPrefix()

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
		&order.Id,
		&order.WorkerId,
		&order.ProductId,
		&order.Tax,
		&order.Discount,
		&order.TotalPrice,
		&order.CreatedAt,
		&order.UpdatedAt,
	); err != nil {
		return nil, p.db.Error(err)
	}

	return &order, nil
}

func (p orderRepo) GetOrders(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Order, error) {
	// ctx, span := otlp.Start(ctx, userServiceName, userSpanRepoPrefix+"List")
	// defer span.End()

	// fmt.Println(filter)

	// for key, value := range filter {
	// 	if key == "type_id" || key == "lang" || key == "status" {
	// 		queryBuilder = queryBuilder.Where(p.db.Sq.Equal(key, value))
	// 		continue
	// 	}
	// 	if key == "created_at" {
	// 		queryBuilder = queryBuilder.Where("created_at=?", value)
	// 		continue
	// 	}
	// }

	var (
		orders []*entity.Order
	)
	queryBuilder := p.orderSelectQueryPrefix()

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(limit).Offset(offset)
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "getOrders"))
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, p.db.Error(err)
	}
	defer rows.Close()
	orders = make([]*entity.Order, 0)
	for rows.Next() {
		var order entity.Order
		if err = rows.Scan(
			&order.Id,
			&order.WorkerId,
			&order.ProductId,
			&order.Tax,
			&order.Discount,
			&order.TotalPrice,
			&order.CreatedAt,
			&order.UpdatedAt,
		); err != nil {
			return nil, p.db.Error(err)
		}
		orders = append(orders, &order)
	}

	return orders, nil
}

func (p orderRepo) UpdateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error) {
	// ctx, span := otlp.Start(ctx, userServiceName, userSpanRepoPrefix+"Update")
	// defer span.End()

	clauses := map[string]any{
		"tax":         order.Tax,
		"discount":    order.Discount,
		"total_price": order.TotalPrice,
		"updated_at":  order.UpdatedAt,
	}

	query, args, err := p.db.Sq.Builder.
		Update(p.tableName).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", order.Id)).
		ToSql()
	if err != nil {
		return &entity.Order{}, p.db.ErrSQLBuild(err, p.tableName+" update")
	}

	query += " RETURNING id, worker_id, product_id, tax, discount, total_price, created_at, updated_at"

	row := p.db.QueryRow(ctx, query, args...)

	var updated entity.Order

	err = row.Scan(
		&updated.Id,
		&updated.WorkerId,
		&updated.ProductId,
		&updated.Tax,
		&updated.Discount,
		&updated.TotalPrice,
		&updated.CreatedAt,
		&updated.UpdatedAt)

	pp.Println(updated)

	if err != nil {
		return &entity.Order{}, err
	}
	return &updated, nil
}

func (p orderRepo) DeleteOrder(ctx context.Context, id int64) error {
	// ctx, span := otlp.Start(ctx, userServiceName, userSpanRepoPrefix+"Delete")
	// defer span.End()

	sqlStr, args, err := p.db.Sq.Builder.
		Delete(p.tableName).
		Where(p.db.Sq.Equal("id", id)).
		ToSql()
	if err != nil {
		return p.db.ErrSQLBuild(err, p.tableName+" delete")
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
