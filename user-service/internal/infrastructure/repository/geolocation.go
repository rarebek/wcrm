package repository

import (
	"context"
	"user-service/internal/entity"
)

type Geolocations interface {
	CreateGeolocation(ctx context.Context, kyc *entity.Geolocation) error
	GetGeolocation(ctx context.Context, params map[string]int64) (*entity.Geolocation, error)
	UpdateGeolocation(ctx context.Context, kyc *entity.Geolocation) error
	DeleteGeolocation(ctx context.Context, guid int64) error
	ListGeolocation(ctx context.Context, filter map[string]string) ([]*entity.Geolocation, error)
}