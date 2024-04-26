package refresh_token

// import (
// 	"context"
// 	"time"

// 	"evrone_service/api_gateway/internal/entity"
// )

// type RefreshToken interface {
// 	Get(ctx context.Context, refreshToken string) (*entity.RefreshToken, error)
// 	Create(ctx context.Context, m *entity.RefreshToken) error
// 	Delete(ctx context.Context, refreshToken string) error
// 	GenerateToken(ctx context.Context, sub, tokenType, jwtSecret string, accessTTL, refreshTTL time.Duration, optionalFields ...map[string]interface{}) (string, string, error)
// }

// type RefreshTokenRepo interface {
// 	Get(ctx context.Context, refreshToken string) (*entity.RefreshToken, error)
// 	Create(ctx context.Context, m *entity.RefreshToken) error
// 	Delete(ctx context.Context, refreshToken string) error
// }
