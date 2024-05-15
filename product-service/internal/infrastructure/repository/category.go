package repository

import (
	"context"
	"wcrm/product-service/internal/entity"
)

type Category interface {
	CreateCategory(ctx context.Context, kyc *entity.Category) (*entity.Category, error)
	GetCategory(ctx context.Context, params map[string]string) (*entity.Category, error)
	ListCategory(ctx context.Context, limit, offset uint64, filter map[string]string) (*entity.AllCategory, error)
	UpdateCategory(ctx context.Context, kyc *entity.Category) (*entity.Category, error)
	DeleteCategory(ctx context.Context, id string) (*entity.CheckResponse, error)
}
