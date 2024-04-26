package app_version

import (
	"context"
	"time"

	"evrone_service/api_gateway/internal/entity"
	"evrone_service/api_gateway/internal/infrastructure/repository/postgresql/repo"
	"evrone_service/api_gateway/internal/pkg/otlp"
)

type appVersionService struct {
	ctxTimeout time.Duration
	repo       repo.AppVersionRepo
}

func NewAppVersionService(ctxTimeout time.Duration, repo repo.AppVersionRepo) AppVersion {
	return &appVersionService{
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

func (r *appVersionService) beforeCreate(m *entity.AppVersion) {
	m.CreatedAt = time.Now().UTC()
	m.UpdatedAt = time.Now().UTC()
}

func (r *appVersionService) beforeUpdate(m *entity.AppVersion) {
	m.UpdatedAt = time.Now().UTC()
}

func (r *appVersionService) Get(ctx context.Context) (*entity.AppVersion, error) {
	ctx, cancel := context.WithTimeout(ctx, r.ctxTimeout)
	defer cancel()

	// tracing
	ctx, span := otlp.Start(ctx, "refreshTokenService", "refreshTokenUsecaseGet")
	defer span.End()

	return r.repo.Get(ctx)
}

func (r *appVersionService) Create(ctx context.Context, m *entity.AppVersion) error {
	ctx, cancel := context.WithTimeout(ctx, r.ctxTimeout)
	defer cancel()

	// tracing
	ctx, span := otlp.Start(ctx, "refreshTokenService", "refreshTokenUsecaseCreate")
	defer span.End()

	r.beforeCreate(m)
	return r.repo.Create(ctx, m)
}

func (r *appVersionService) Update(ctx context.Context, m *entity.AppVersion) error {
	ctx, cancel := context.WithTimeout(ctx, r.ctxTimeout)
	defer cancel()

	// tracing
	ctx, span := otlp.Start(ctx, "refreshTokenService", "refreshTokenUsecaseUpdate")
	defer span.End()

	r.beforeUpdate(m)
	return r.repo.Update(ctx, m)
}
