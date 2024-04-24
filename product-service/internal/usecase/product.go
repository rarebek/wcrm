package usecase

import (
	"wcrm/product-service/internal/entity"
	"wcrm/product-service/internal/infrastructure/repository"

	// "wcrm/product-service/internal/pkg/otlp"
	"context"
	"time"
)

// const (
// 	serviceNameUser = "contentService"
// 	spanNameUser    = "contentUsecase"
// )

type Product interface {
	CreateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error)
	GetProduct(ctx context.Context, params map[string]int64) (*entity.Product, error)
	ListProduct(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Product, error)
	UpdateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error)
	DeleteProduct(ctx context.Context, id int64) (*entity.CheckResponse, error)
}

type newsService struct {
	BaseUseCase
	repo       repository.Product
	ctxTimeout time.Duration
	// client 
}

func NewProductService(ctxTimeout time.Duration, repo repository.Product) newsService {
	return newsService{
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

func (u newsService) CreateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	// ctx, span := otlp.Start(ctx, serviceNameUser, spanNameUser+"Create")
	// defer span.End()

	createdProduct, err := u.repo.CreateProduct(ctx, product)
	if err != nil {
		return &entity.Product{}, nil
	}

	return createdProduct, nil
}

func (u newsService) GetProduct(ctx context.Context, params map[string]int64) (*entity.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	// ctx, span := otlp.Start(ctx, serviceNameUser, spanNameUser+"Get")
	// defer span.End()

	return u.repo.GetProduct(ctx, params)
}

func (u newsService) ListProduct(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	// ctx, span := otlp.Start(ctx, serviceNameUser, spanNameUser+"List")
	// defer span.End()

	return u.repo.ListProduct(ctx, limit, offset, filter)
}

func (u newsService) UpdateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	// ctx, span := otlp.Start(ctx, serviceNameUser, spanNameUser+"Update")
	// defer span.End()

	updatedProduct, err := u.repo.UpdateProduct(ctx, product)
	if err != nil {
		return &entity.Product{}, nil
	}

	return updatedProduct, nil
}

func (u newsService) DeleteProduct(ctx context.Context, id int64) (*entity.CheckResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	// ctx, span := otlp.Start(ctx, serviceNameUser, spanNameUser+"Delete")
	// defer span.End()

	deleteProduct, err := u.repo.DeleteProduct(ctx, id)

	return deleteProduct, err
}
