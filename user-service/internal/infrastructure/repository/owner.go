package repository

import (
	"context"
	"user-service/internal/entity"
)

type Owners interface {
	Create(ctx context.Context, kyc *entity.Owner) (*entity.Owner, error)
	Get(ctx context.Context, params map[string]string) (*entity.Owner, error)
	Update(ctx context.Context, kyc *entity.Owner) (*entity.Owner, error)
	Delete(ctx context.Context, guid string) (*entity.CheckResponse, error)
	List(ctx context.Context, limit, offset uint64, filter map[string]string) (*entity.AllOwners, error)
	CheckField(ctx context.Context, field, value string) (bool, error)
}
