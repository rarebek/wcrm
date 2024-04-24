package repository

import (
	"context"
	"order-service/internal/entity"
)

type Order interface {
	CreateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error)
	GetOrder(ctx context.Context, params map[string]int64) (*entity.Order, error)
	GetOrders(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Order, error)
	UpdateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error)
	DeleteOrder(ctx context.Context, id int64) error
}
