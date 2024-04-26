package postgresql

import (
	"context"
	"fmt"

	// "projects/order-service/genproto/order"
	"projects/order-service/internal/entity"

	"projects/order-service/internal/pkg/logger"
	"projects/order-service/internal/pkg/otlp"
	"projects/order-service/internal/pkg/postgres"

	"github.com/Masterminds/squirrel"
	"github.com/k0kubun/pp"
	"go.opentelemetry.io/otel/attribute"
)

const (
	or_proTableName     = "orders_products"
	orderTableName      = "orders"
	orderServiceName    = "orderService"
	orderSpanRepoPrefix = "orderRepo"
)

type orderRepo struct {
	tableName   string
	db          *postgres.PostgresDB
	opTableName string
}

func NewOrderRepo(db *postgres.PostgresDB) *orderRepo {
	return &orderRepo{
		tableName:   orderTableName,
		db:          db,
		opTableName: or_proTableName,
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
	ctx, span := otlp.Start(ctx, orderServiceName, orderSpanRepoPrefix+"Create")
	span.SetAttributes(attribute.String("create", "order"))
	defer span.End()

	data := map[string]any{
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

	query += " RETURNING id, worker_id, product_id, tax, discount, total_price, created_at, updated_at"
	row := p.db.QueryRow(ctx, query, args...)

	var createdOrder entity.Order

	err = row.Scan(
		&createdOrder.Id,
		&createdOrder.WorkerId,
		&createdOrder.ProductId,
		&createdOrder.Tax,
		&createdOrder.Discount,
		&createdOrder.TotalPrice,
		&createdOrder.CreatedAt,
		&createdOrder.UpdatedAt,
	)

	if err != nil {
		return &entity.Order{}, err
	}

	op_data := map[string]any{
		"order_id":   createdOrder.Id,
		"product_id": createdOrder.ProductId,
	}

	opquery, op_args, err := p.db.Sq.Builder.Insert(p.opTableName).SetMap(op_data).ToSql()
	if err != nil {
		return &entity.Order{}, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.opTableName, "opcreate"))
	}

	// opquery += " RETURNING id, order_id, product_id"

	// op_row := p.db.QueryRow(ctx, opquery, op_args...)

	_, err = p.db.Exec(ctx, opquery, op_args...)
	if err != nil {
		logger.Error(err)
	}

	// var op_order entity.OPModel

	// errr = op_row.Scan(
	// 	&op_order.Id,
	// 	&op_order.OrderId,
	// 	&op_order.ProductID,
	// )

	return &createdOrder, nil
}

func (p orderRepo) GetOrder(ctx context.Context, params map[string]int64) (*entity.Order, error) {
	ctx, span := otlp.Start(ctx, orderServiceName, orderSpanRepoPrefix+"Get")
	span.SetAttributes(attribute.String("get", "order"))
	defer span.End()

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
	ctx, span := otlp.Start(ctx, orderServiceName, orderSpanRepoPrefix+"Gets")
	span.SetAttributes(attribute.String("gets", "order"))
	defer span.End()

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
	ctx, span := otlp.Start(ctx, orderServiceName, orderSpanRepoPrefix+"Update")
	span.SetAttributes(attribute.String("update", "order"))
	defer span.End()

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
	ctx, span := otlp.Start(ctx, orderServiceName, orderSpanRepoPrefix+"Delete")
	span.SetAttributes(attribute.String("delete", "order"))
	defer span.End()

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
