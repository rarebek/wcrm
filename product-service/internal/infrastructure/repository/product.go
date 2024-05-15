package repository

import (
	"context"
	"wcrm/product-service/internal/entity"
)

type Product interface {
	CreateProduct(ctx context.Context, kyc *entity.ProductWithCategoryId) (*entity.Product, error)
	GetProduct(ctx context.Context, params map[string]string) (*entity.Product, error)
	ListProduct(ctx context.Context, limit, offset uint64, filter map[string]string) (*entity.AllProduct, error)
	UpdateProduct(ctx context.Context, kyc *entity.Product) (*entity.Product, error)
	DeleteProduct(ctx context.Context, id string) (*entity.CheckResponse, error)
	SearchProduct(ctx context.Context, limit, offset int64, title string, ownerId string) (*entity.AllProduct, error)
	GetAllProductByCategoryId(ctx context.Context, limit, offset uint64, id string) (*entity.AllProduct, error)
}
