package refresh_token

// import (
// 	"context"
// 	"time"

// 	"github.com/google/uuid"

// 	"api-gateway/internal/entity"
// 	"api-gateway/internal/pkg/otlp"
// 	"api-gateway/internal/pkg/token"
// )

// type refreshTokenService struct {
// 	ctxTimeout time.Duration
// 	repo       RefreshTokenRepo
// }

// func NewRefreshTokenService(ctxTimeout time.Duration, repo RefreshTokenRepo) RefreshToken {
// 	return &refreshTokenService{
// 		ctxTimeout: ctxTimeout,
// 		repo:       repo,
// 	}
// }

// func (r *refreshTokenService) beforeCreate(m *entity.RefreshToken) error {
// 	m.GUID = uuid.New().String()
// 	m.CreatedAt = time.Now().UTC()
// 	return nil
// }

// func (r *refreshTokenService) Get(ctx context.Context, refreshToken string) (*entity.RefreshToken, error) {
// 	ctx, cancel := context.WithTimeout(ctx, r.ctxTimeout)
// 	defer cancel()

// 	// tracing
// 	ctx, span := otlp.Start(ctx, "refreshTokenService", "refreshTokenUsecaseGet")
// 	defer span.End()

// 	return r.repo.Get(ctx, refreshToken)
// }

// func (r *refreshTokenService) Create(ctx context.Context, m *entity.RefreshToken) error {
// 	ctx, cancel := context.WithTimeout(ctx, r.ctxTimeout)
// 	defer cancel()

// 	// tracing
// 	ctx, span := otlp.Start(ctx, "refreshTokenService", "refreshTokenUsecaseCreate")
// 	defer span.End()

// 	r.beforeCreate(m)
// 	return r.repo.Create(ctx, m)
// }

// func (r *refreshTokenService) Delete(ctx context.Context, refreshToken string) error {
// 	ctx, cancel := context.WithTimeout(ctx, r.ctxTimeout)
// 	defer cancel()

// 	// tracing
// 	ctx, span := otlp.Start(ctx, "refreshTokenService", "refreshTokenUsecaseDelete")
// 	defer span.End()

// 	return r.repo.Delete(ctx, refreshToken)
// }

// func (r *refreshTokenService) GenerateToken(ctx context.Context, sub, tokenType, jwtSecret string, accessTTL, refreshTTL time.Duration, optionalFields ...map[string]interface{}) (string, string, error) {
// 	// tracing
// 	ctx, span := otlp.Start(ctx, "refreshTokenService", "refreshTokenUsecaseGenerateToken")
// 	defer span.End()

// 	accessToken, refreshToken, err := token.GenerateToken(sub, tokenType, jwtSecret, accessTTL, refreshTTL, optionalFields...)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	m := entity.RefreshToken{
// 		RefreshToken: refreshToken,
// 		ExpiryDate:   time.Now().Add(refreshTTL),
// 	}

// 	err = r.Create(ctx, &m)
// 	if err != nil {
// 		return "", "", err
// 	}
// 	return accessToken, refreshToken, nil
// }
