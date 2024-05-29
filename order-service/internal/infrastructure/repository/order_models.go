package repository

import (
	"context"
	"projects/order-service/internal/entity"
)

type Order interface {
	CreateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error)
	GetOrder(ctx context.Context, id string) (*entity.Order, error)
	GetOrders(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.GetAllOrdersResponse, error)
	UpdateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error)
	DeleteOrder(ctx context.Context, id string) error
}
