package usecase

import (
	"context"
	"time"
	"user-service/internal/entity"
	"user-service/internal/infrastructure/repository"
	"user-service/internal/pkg/otlp"
)

const (
	serviceNameGeolocation = "geolocationService"
	spanNameGeolocation    = "geolocationUsecase"
)


type Geolocation interface {
	Create(ctx context.Context, geolocation *entity.Geolocation) (*entity.Geolocation, error)
	Get(ctx context.Context, params map[string]int64) (*entity.Geolocation, error)
	Update(ctx context.Context, geolocation *entity.Geolocation) (*entity.Geolocation, error)
	Delete(ctx context.Context, guid int64) (*entity.CheckResponse, error)
	List(ctx context.Context, id string, limit, offset uint64, filter map[string]string) (*entity.AllGeolocation, error)
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

func (u geolocationService) Create(ctx context.Context, geolocation *entity.Geolocation) (*entity.Geolocation, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameGeolocation, spanNameGeolocation+"Create")
	defer span.End()

	return u.repo.Create(ctx, geolocation)
}

func (u geolocationService) Get(ctx context.Context, params map[string]int64) (*entity.Geolocation, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameGeolocation, spanNameGeolocation+"Create")
	defer span.End()

	return u.repo.Get(ctx, params)
}

func (u geolocationService) Update(ctx context.Context, geolocation *entity.Geolocation) (*entity.Geolocation, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameGeolocation, spanNameGeolocation+"Create")
	defer span.End()

	return u.repo.Update(ctx, geolocation)
}

func (u geolocationService) Delete(ctx context.Context, guid int64) (*entity.CheckResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, serviceNameGeolocation, spanNameGeolocation+"Create")
	defer span.End()

	return u.repo.Delete(ctx, guid)
}

func (u geolocationService) List(ctx context.Context, id string, limit, offset uint64, filter map[string]string) (*entity.AllGeolocation, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()
	
	ctx, span := otlp.Start(ctx, serviceNameGeolocation, spanNameGeolocation+"Create")
	defer span.End()

	return u.repo.List(ctx, id, limit, offset, filter)
}