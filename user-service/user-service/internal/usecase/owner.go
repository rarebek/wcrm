package usecase

import (
	"context"
	"time"
	"user-service/internal/entity"
	"user-service/internal/infrastructure/repository"
)

type Owner interface {
	CreateOwner(ctx context.Context, owner *entity.Owner) (string, error)
	GetOwner(ctx context.Context, params map[string]string) (*entity.Owner, error)
	UpdateOwner(ctx context.Context, owner *entity.Owner) error
	DeleteOwner(ctx context.Context, guid string) error
	ListOwner(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Owner, error)
	CheckFieldOwner(ctx context.Context, field, value string) (bool, error)
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

func (u ownerService) CreateOwner(ctx context.Context, owner *entity.Owner) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	u.beforeRequest(&owner.Id, &owner.CreatedAt, &owner.UpdatedAt)

	return owner.Id, u.repo.CreateOwner(ctx, owner)
}

func (u ownerService) GetOwner(ctx context.Context, params map[string]string) (*entity.Owner, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.GetOwner(ctx, params)
}

func (u ownerService) UpdateOwner(ctx context.Context, owner *entity.Owner) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	u.beforeRequest(nil, nil, &owner.UpdatedAt)

	return u.repo.UpdateOwner(ctx, owner)
}

func (u ownerService) DeleteOwner(ctx context.Context, guid string) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.DeleteOwner(ctx, guid)
}

func (u ownerService) ListOwner(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Owner, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.ListOwner(ctx, limit, offset, filter)
}

func (u ownerService) CheckFieldOwner(ctx context.Context, field, value string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.CheckFieldOwner(ctx, field, value)
}