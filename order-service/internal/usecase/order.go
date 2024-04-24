package usecase

import (
	"order-service/internal/entity"
	"order-service/internal/infrastructure/repository"

	// "order-service/internal/pkg/otlp"
	"context"
	"time"
)

// const (
// 	serviceNameUser = "contentService"
// 	spanNameUser    = "contentUsecase"
// )

type Order interface {
	CreateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error)
	GetOrder(ctx context.Context, params map[string]int64) (*entity.Order, error)
	GetOrders(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Order, error)
	UpdateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error)
	DeleteOrder(ctx context.Context, id int64) error
}

type newOrderService struct {
	BaseUseCase
	repo       repository.Order
	ctxTimeout time.Duration
}

func NewOrderService(ctxTimeout time.Duration, repo repository.Order) newOrderService {
	return newOrderService{
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

func (u newOrderService) CreateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	// ctx, span := otlp.Start(ctx, serviceNameUser, spanNameUser+"Create")
	// defer span.End()

	createdOrder, err := u.repo.CreateOrder(ctx, order)
	if err != nil {
		return &entity.Order{}, nil
	}

	return createdOrder, nil
}

func (u newOrderService) GetOrder(ctx context.Context, params map[string]int64) (*entity.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	// ctx, span := otlp.Start(ctx, serviceNameUser, spanNameUser+"Get")
	// defer span.End()

	return u.repo.GetOrder(ctx, params)
}

func (u newOrderService) GetOrders(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	// ctx, span := otlp.Start(ctx, serviceNameUser, spanNameUser+"List")
	// defer span.End()

	return u.repo.GetOrders(ctx, limit, offset, filter)
}

func (u newOrderService) UpdateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	// ctx, span := otlp.Start(ctx, serviceNameUser, spanNameUser+"Update")
	// defer span.End()

	updatedOrder, err := u.repo.UpdateOrder(ctx, order)
	if err != nil {
		return &entity.Order{}, nil
	}

	return updatedOrder, nil
}

func (u newOrderService) DeleteOrder(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	// ctx, span := otlp.Start(ctx, serviceNameUser, spanNameUser+"Delete")
	// defer span.End()

	return u.repo.DeleteOrder(ctx, id)
}