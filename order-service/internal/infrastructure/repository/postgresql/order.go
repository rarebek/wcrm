package postgresql

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	// "projects/order-service/genproto/order"
	"projects/order-service/internal/entity"

	"projects/order-service/internal/pkg/otlp"
	"projects/order-service/internal/pkg/postgres"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
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
			"table_number",
			"worker_id",
			"products",
			"tax",
			"discount",
			"total_price",
			"created_at",
			"updated_at",
		).From(p.tableName)
}

func (p *orderRepo) CreateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error) {
	ctx, span := otlp.Start(ctx, orderServiceName, orderSpanRepoPrefix+"Create")
	span.SetAttributes(attribute.String("create", "order"))
	defer span.End()
	pp.Println(order)

	// Convert Products to JSONB
	productsJSON, err := json.Marshal(order.Products)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal products: %w", err)
	}

	data := map[string]interface{}{
		"id":           uuid.NewString(),
		"table_number": order.TableNumber,
		"worker_id":    order.WorkerId,
		"products":     productsJSON,
		"tax":          order.Tax,
		"discount":     order.Discount,
		"total_price":  order.TotalPrice,
		"created_at":   time.Now(),
		"updated_at":   time.Now(),
	}

	query, args, err := p.db.Sq.Builder.Insert(p.tableName).
		SetMap(data).
		Suffix("RETURNING id, table_number, worker_id, products, tax, total_price, created_at, updated_at").
		ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "create"))
	}

	row := p.db.QueryRow(ctx, query, args...)
	var createdOrder entity.Order

	var productsJSON2 []byte

	err = row.Scan(
		&createdOrder.Id,
		&createdOrder.TableNumber,
		&createdOrder.WorkerId,
		&productsJSON,
		&createdOrder.Tax,
		&createdOrder.TotalPrice,
		&createdOrder.CreatedAt,
		&createdOrder.UpdatedAt,
	)
	if err != nil {
		log.Println(err)
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "create"))
	}

	err = json.Unmarshal(productsJSON2, &createdOrder.Products)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to unmarshal products: %w", err)
	}

	pp.Println("CREATED ORDER: ", createdOrder)

	return &createdOrder, nil
}

func (p orderRepo) GetOrder(ctx context.Context, id string) (*entity.Order, error) {
	ctx, span := otlp.Start(ctx, orderServiceName, orderSpanRepoPrefix+"Get")
	span.SetAttributes(attribute.String("get", "order"))
	defer span.End()

	var order entity.Order

	queryBuilder := p.orderSelectQueryPrefix()
	sql, args, err := queryBuilder.Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.opTableName, "get"))
	}

	var productsJSON []byte // Declaring productsJSON again to scan into it

	if err = p.db.QueryRow(ctx, sql, args...).Scan(
		&order.Id,
		&order.WorkerId,
		&productsJSON, // Scanning into productsJSON
		&order.Tax,
		&order.Discount,
		&order.TotalPrice,
		&order.CreatedAt,
		&order.UpdatedAt,
	); err != nil {
		return nil, p.db.Error(err)
	}

	// Unmarshal productsJSON into the Products field of order
	err = json.Unmarshal(productsJSON, &order.Products)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal products: %w", err)
	}

	query := `SELECT full_name from workers WHERE id = $1;`
	row := p.db.QueryRow(ctx, query, order.WorkerId).Scan(&order.WorkerName)
	if row == pgx.ErrNoRows {
		return nil, fmt.Errorf("worker with id %s not found", order.WorkerId)
	}

	return &order, nil
}

func (p orderRepo) GetOrders(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.GetAllOrdersResponse, error) {
	ctx, span := otlp.Start(ctx, orderServiceName, orderSpanRepoPrefix+"Gets")
	span.SetAttributes(attribute.String("gets", "order"))
	defer span.End()

	var orders []*entity.GetAllOrdersResponse
	queryBuilder := p.orderSelectQueryPrefix()

	if limit != 0 {
		queryBuilder = queryBuilder.Limit(limit).Offset(offset)
	}

	queryBuilder = queryBuilder.Where(squirrel.Eq{
		"worker_id": filter["worker_id"],
	})
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, fmt.Sprintf("%s %s", p.tableName, "getOrders"))
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, p.db.Error(err)
	}
	defer rows.Close()

	for rows.Next() {
		var order entity.Order
		var productsJSON []byte
		if err = rows.Scan(
			&order.Id,
			&order.TableNumber,
			&order.WorkerId,
			&productsJSON,
			&order.Tax,
			&order.Discount,
			&order.TotalPrice,
			&order.CreatedAt,
			&order.UpdatedAt,
		); err != nil {
			return nil, p.db.Error(err)
		}

		// Unmarshal productsJSON into the Products field of order
		err = json.Unmarshal(productsJSON, &order.Products)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal products: %w", err)
		}

		// Create a new instance of GetAllOrdersResponse
		response := &entity.GetAllOrdersResponse{
			Orders: []entity.Order{order},
		}

		// Append WorkerName to the response
		query := `SELECT full_name from workers where id = $1;`
		if err := p.db.QueryRow(ctx, query, filter["worker_id"]).Scan(&response.WorkerName); err != nil {
			return nil, p.db.Error(err)
		}

		orders = append(orders, response)
	}

	query = `SELECT full_name from workers where id = $1;`
	p.db.QueryRow(ctx, query, filter["worker_id"]).Scan(&orders[0].WorkerName)

	return orders, nil
}

func (p orderRepo) UpdateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error) {
	ctx, span := otlp.Start(ctx, orderServiceName, orderSpanRepoPrefix+"Update")
	span.SetAttributes(attribute.String("update", "order"))
	defer span.End()

	productIdsJSON, err := json.Marshal(order.Products)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal products: %w", err)
	}

	clauses := map[string]interface{}{
		"worker_id":   order.WorkerId,
		"products":    productIdsJSON,
		"tax":         order.Tax,
		"discount":    order.Discount,
		"total_price": order.TotalPrice,
		"updated_at":  order.UpdatedAt,
	}

	query, args, err := p.db.Sq.Builder.
		Update(p.tableName).
		SetMap(clauses).
		Where(p.db.Sq.Equal("id", order.Id)).
		Suffix("RETURNING id, worker_id, products, tax, discount, total_price, created_at, updated_at").
		ToSql()
	if err != nil {
		return nil, p.db.ErrSQLBuild(err, p.tableName+" update")
	}

	row := p.db.QueryRow(ctx, query, args...)

	var updated entity.Order
	var productsJSON []byte // Declaring productsJSON again to scan into it

	err = row.Scan(
		&updated.Id,
		&updated.WorkerId,
		&productsJSON, // Scanning into productsJSON
		&updated.Tax,
		&updated.Discount,
		&updated.TotalPrice,
		&updated.CreatedAt,
		&updated.UpdatedAt,
	)
	if err != nil {
		return nil, p.db.Error(err)
	}

	// Unmarshal productsJSON into the Products field of updated
	err = json.Unmarshal(productsJSON, &updated.Products)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal products: %w", err)
	}

	return &updated, nil
}

func (p orderRepo) DeleteOrder(ctx context.Context, id string) error {
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
