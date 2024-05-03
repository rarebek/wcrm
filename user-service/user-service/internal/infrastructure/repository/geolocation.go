package repository

import (
	"context"
	"user-service/internal/entity"
)

type Geolocations interface {
	Create(ctx context.Context, kyc *entity.Geolocation) error
	Get(ctx context.Context, params map[string]int64) (*entity.Geolocation, error)
	Update(ctx context.Context, kyc *entity.Geolocation) error
	Delete(ctx context.Context, guid int64) error
	List(ctx context.Context, filter map[string]string) ([]*entity.Geolocation, error)
}