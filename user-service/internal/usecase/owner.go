package usecase

import (
	"context"
	"time"
	"user-service/internal/entity"
	"user-service/internal/infrastructure/repository"
)

type Owner interface {
	Create(ctx context.Context, owner *entity.Owner) (*entity.Owner, error)
	Get(ctx context.Context, params map[string]string) (*entity.Owner, error)
	Update(ctx context.Context, owner *entity.Owner) (*entity.Owner, error)
	Delete(ctx context.Context, guid string) (*entity.Owner, error)
	List(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Owner, error)
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

	u.beforeRequest(&owner.Id, &owner.CreatedAt, &owner.UpdatedAt)

	owner, err := u.repo.Create(ctx, owner)
	if err != nil {
		return nil, err
	}

	return owner, nil
}

func (u ownerService) Get(ctx context.Context, params map[string]string) (*entity.Owner, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.Get(ctx, params)
}

func (u ownerService) Update(ctx context.Context, owner *entity.Owner) (*entity.Owner, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	u.beforeRequest(nil, nil, &owner.UpdatedAt)

	return u.repo.Update(ctx, owner)
}

func (u ownerService) Delete(ctx context.Context, guid string) (*entity.Owner, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.Delete(ctx, guid)
}

func (u ownerService) List(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Owner, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.List(ctx, limit, offset, filter)
}

func (u ownerService) CheckField(ctx context.Context, field, value string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.CheckField(ctx, field, value)
}
