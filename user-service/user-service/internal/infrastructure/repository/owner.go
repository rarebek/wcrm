package repository

import (
	"context"
	"user-service/internal/entity"
)

type Owners interface {
	CreateOwner(ctx context.Context, kyc *entity.Owner) error
	GetOwner(ctx context.Context, params map[string]string) (*entity.Owner, error)
	UpdateOwner(ctx context.Context, kyc *entity.Owner) error
	DeleteOwner(ctx context.Context, guid string) error
	ListOwner(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Owner, error)
	CheckFieldOwner(ctx context.Context, field, value string) (bool, error)
}
