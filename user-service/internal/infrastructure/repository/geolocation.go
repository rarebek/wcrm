package repository

import (
	"context"
	"user-service/internal/entity"
)

type Geolocations interface {
	Create(ctx context.Context, kyc *entity.Geolocation) (*entity.Geolocation, error)
	Get(ctx context.Context, params map[string]int64) (*entity.Geolocation, error)
	Update(ctx context.Context, kyc *entity.Geolocation) (*entity.Geolocation, error)
	Delete(ctx context.Context, guid int64) (*entity.CheckResponse, error)
	List(ctx context.Context, id string, limit, offset uint64, filter map[string]string) (*entity.AllGeolocation, error)
}
