package usecase

import (
	"context"
	"time"
	"user-service/internal/entity"
	"user-service/internal/infrastructure/repository"
)

type Geolocation interface {
	Create(ctx context.Context, geolocation *entity.Geolocation) (int64, error)
	Get(ctx context.Context, params map[string]int64) (*entity.Geolocation, error)
	Update(ctx context.Context, geolocation *entity.Geolocation) error
	Delete(ctx context.Context, guid int64) error
	List(ctx context.Context, filter map[string]string) ([]*entity.Geolocation, error)
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

func (u geolocationService) Create(ctx context.Context, geolocation *entity.Geolocation) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return geolocation.Id, u.repo.Create(ctx, geolocation)
}

func (u geolocationService) Get(ctx context.Context, params map[string]int64) (*entity.Geolocation, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.Get(ctx, params)
}

func (u geolocationService) Update(ctx context.Context, geolocation *entity.Geolocation) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.Update(ctx, geolocation)
}

func (u geolocationService) Delete(ctx context.Context, guid int64) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.Delete(ctx, guid)
}

func (u geolocationService) List(ctx context.Context, filter map[string]string) ([]*entity.Geolocation, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	return u.repo.List(ctx, filter)
}