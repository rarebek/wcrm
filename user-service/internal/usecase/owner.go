package usecase

import (
	"context"
	"time"
	"user-service/internal/entity"
	"user-service/internal/infrastructure/repository"
	"user-service/internal/pkg/otlp"
)

const (
	serviceNameOwner = "ownerService"
	spanNameOwner    = "ownerUsecase"
)

type Owner interface {
	Create(ctx context.Context, owner *entity.Owner) (*entity.Owner, error)
	Get(ctx context.Context, params map[string]string) (*entity.Owner, error)
	Update(ctx context.Context, owner *entity.Owner) (*entity.Owner, error)
	Delete(ctx context.Context, guid string) (*entity.CheckResponse, error)
	List(ctx context.Context, limit, offset uint64, filter map[string]string) (*entity.AllOwners, error)
	CheckField(ctx context.Context, field, value string) (bool, error)
}

type ownerService struct {
	BaseUseCase
	repo       repository.Owners
	ctxTimeout time.Duration
}

func NewOwnerService(ctxTimeout time.Duration, repo repository.Owners) ownerService {
	return ownerService{
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

func (u ownerService) Create(ctx context.Context, owner *entity.Owner) (*entity.Owner, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameOwner, spanNameOwner+"Create")
	defer span.End()

	u.beforeRequest(&owner.Id, &owner.CreatedAt, &owner.UpdatedAt)

	return u.repo.Create(ctx, owner)
}

func (u ownerService) Get(ctx context.Context, params map[string]string) (*entity.Owner, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameOwner, spanNameOwner+"Get")
	defer span.End()

	return u.repo.Get(ctx, params)
}

func (u ownerService) Update(ctx context.Context, owner *entity.Owner) (*entity.Owner, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameOwner, spanNameOwner+"Update")
	defer span.End()

	u.beforeRequest(nil, nil, &owner.UpdatedAt)

	return u.repo.Update(ctx, owner)
}

func (u ownerService) Delete(ctx context.Context, guid string) (*entity.CheckResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameOwner, spanNameOwner+"Delete")
	defer span.End()

	res, err := u.repo.Delete(ctx, guid)

	return res, err
}

func (u ownerService) List(ctx context.Context, limit, offset uint64, filter map[string]string) (*entity.AllOwners, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameOwner, spanNameOwner+"List")
	defer span.End()

	return u.repo.List(ctx, limit, offset, filter)
}

func (u ownerService) CheckField(ctx context.Context, field, value string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameOwner, spanNameOwner+"CheckField")
	defer span.End()

	return u.repo.CheckField(ctx, field, value)
}