package usecase

import (
	"context"
	"time"
	"user-service/internal/entity"
	"user-service/internal/infrastructure/repository"
	"user-service/internal/pkg/otlp"
)

const (
	serviceNameWorker = "workerService"
	spanNameWorker    = "workerUsecase"
)

type Worker interface {
	Create(ctx context.Context, worker *entity.Worker) (*entity.Worker, error)
	Get(ctx context.Context, params map[string]string) (*entity.Worker, error)
	Update(ctx context.Context, worker *entity.Worker) (*entity.Worker, error)
	Delete(ctx context.Context, guid string) (*entity.CheckResponse, error)
	List(ctx context.Context, limit, offset uint64, filter map[string]string) (*entity.AllWorker, error)
	CheckField(ctx context.Context, field, value string) (bool, error)
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

func (u workerService) Create(ctx context.Context, worker *entity.Worker) (*entity.Worker, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameWorker, spanNameWorker+"Create")
	defer span.End()

	u.beforeRequest(&worker.Id, &worker.CreatedAt, &worker.UpdatedAt)

	return u.repo.Create(ctx, worker)
}

func (u workerService) Get(ctx context.Context, params map[string]string) (*entity.Worker, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameWorker, spanNameWorker+"Get")
	defer span.End()

	return u.repo.Get(ctx, params)
}

func (u workerService) Update(ctx context.Context, worker *entity.Worker) (*entity.Worker, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameWorker, spanNameWorker+"Update")
	defer span.End()

	u.beforeRequest(nil, nil, &worker.UpdatedAt)

	return u.repo.Update(ctx, worker)
}

func (u workerService) Delete(ctx context.Context, guid string) (*entity.CheckResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	req, err := u.repo.Delete(ctx, guid)

	return req, err
}

func (u workerService) List(ctx context.Context, limit, offset uint64, filter map[string]string) (*entity.AllWorker, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameWorker, spanNameWorker+"List")
	defer span.End()

	return u.repo.List(ctx, limit, offset, filter)
}

func (u workerService) CheckField(ctx context.Context, field, value string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameWorker, spanNameWorker+"CheckField")
	defer span.End()

	return u.repo.CheckField(ctx, field, value)
}
