package repository

import (
	"context"
	"user-service/internal/entity"
)

type Workers interface {
	CreateWorker(ctx context.Context, kyc *entity.Worker) error
	GetWorker(ctx context.Context, params map[string]string) (*entity.Worker, error)
	UpdateWorker(ctx context.Context, kyc *entity.Worker) error
	DeleteWorker(ctx context.Context, guid string) error
	ListWorker(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Worker, error)
}
