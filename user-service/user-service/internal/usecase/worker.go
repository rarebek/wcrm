package usecase

import (
	"context"
	"time"
	"user-service/internal/entity"
	"user-service/internal/infrastructure/repository"
)

type Worker interface {
	CreateWorker(ctx context.Context, worker *entity.Worker) (string, error)
	GetWorker(ctx context.Context, params map[string]string) (*entity.Worker, error)
	UpdateWorker(ctx context.Context, worker *entity.Worker) error
	DeleteWorker(ctx context.Context, guid string) error
	ListWorker(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Worker, error)
	CheckFieldWorker(ctx context.Context, field, value string) (bool, error)
}

type workerService struct {
	BaseUseCase
	repo       repository.Workers
	ctxTimeout time.Duration
}

func NewWorkerService(ctxTimeout time.Duration, repo repository.Workers) workerService {
	return workerService{
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

func (u workerService) CreateWorker(ctx context.Context, worker *entity.Worker) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	u.beforeRequest(&worker.Id, &worker.CreatedAt, &worker.UpdatedAt)

	return worker.Id, u.repo.CreateWorker(ctx, worker)
}

func (u workerService) GetWorker(ctx context.Context, params map[string]string) (*entity.Worker, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.GetWorker(ctx, params)
}

func (u workerService) UpdateWorker(ctx context.Context, worker *entity.Worker) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	u.beforeRequest(nil, nil, &worker.UpdatedAt)

	return u.repo.UpdateWorker(ctx, worker)
}

func (u workerService) DeleteWorker(ctx context.Context, guid string) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.DeleteWorker(ctx, guid)
}

func (u workerService) ListWorker(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Worker, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.ListWorker(ctx, limit, offset, filter)
}

func (u workerService) CheckFieldWorker(ctx context.Context, field, value string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.CheckFieldWorker(ctx, field, value)
}