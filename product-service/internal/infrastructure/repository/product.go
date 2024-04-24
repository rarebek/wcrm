package repository

import (
	"context"
	"wcrm/product-service/internal/entity"
)

type Product interface {
	CreateProduct(ctx context.Context, kyc *entity.Product) (*entity.Product, error)
	GetProduct(ctx context.Context, params map[string]int64) (*entity.Product, error)
	ListProduct(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Product, error)
	UpdateProduct(ctx context.Context, kyc *entity.Product) (*entity.Product, error)
	DeleteProduct(ctx context.Context, id int64) (*entity.CheckResponse, error)
}