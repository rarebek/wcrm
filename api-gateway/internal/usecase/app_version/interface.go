package app_version

import (
	"context"

	"evrone_service/api_gateway/internal/entity"
)

type AppVersion interface {
	Get(ctx context.Context) (*entity.AppVersion, error)
	Create(ctx context.Context, m *entity.AppVersion) error
	Update(ctx context.Context, m *entity.AppVersion) error
}
