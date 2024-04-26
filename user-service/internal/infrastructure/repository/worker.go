package repository

import (
	"context"
	"user-service/internal/entity"
)

type Workers interface {
	Create(ctx context.Context, kyc *entity.Worker) error
	Get(ctx context.Context, params map[string]string) (*entity.Worker, error)
	Update(ctx context.Context, kyc *entity.Worker) error
	Delete(ctx context.Context, guid string) error
	List(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Worker, error)
	CheckField(ctx context.Context, field, value string) (bool, error)
}
