package usecase

import (
	"context"
	"time"
	"user-service/internal/entity"
	"user-service/internal/infrastructure/repository"
)

type Geolocation interface {
	CreateGeolocation(ctx context.Context, geolocation *entity.Geolocation) (int64, error)
	GetGeolocation(ctx context.Context, params map[string]int64) (*entity.Geolocation, error)
	UpdateGeolocation(ctx context.Context, geolocation *entity.Geolocation) error
	DeleteGeolocation(ctx context.Context, guid int64) error
	ListGeolocation(ctx context.Context, filter map[string]string) ([]*entity.Geolocation, error)
}

type geolocationService struct {
	BaseUseCase
	repo       repository.Geolocations
	ctxTimeout time.Duration
}

func NewGeolocationService(ctxTimeout time.Duration, repo repository.Geolocations) geolocationService {
	return geolocationService{
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

func (u geolocationService) CreateGeolocation(ctx context.Context, geolocation *entity.Geolocation) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return geolocation.Id, u.repo.CreateGeolocation(ctx, geolocation)
}

func (u geolocationService) GetGeolocation(ctx context.Context, params map[string]int64) (*entity.Geolocation, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.GetGeolocation(ctx, params)
}

func (u geolocationService) UpdateGeolocation(ctx context.Context, geolocation *entity.Geolocation) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.UpdateGeolocation(ctx, geolocation)
}

func (u geolocationService) DeleteGeolocation(ctx context.Context, guid int64) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.DeleteGeolocation(ctx, guid)
}

func (u geolocationService) ListGeolocation(ctx context.Context, filter map[string]string) ([]*entity.Geolocation, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.ListGeolocation(ctx, filter)
}