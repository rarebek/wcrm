package v1

import (
	"time"

	tokens "api-gateway/internal/pkg/token"

	"go.uber.org/zap"

	grpcClients "api-gateway/internal/infrastructure/grpc_service_client"
	"api-gateway/internal/pkg/config"

	// "api-gateway/internal/usecase/event"
	"github.com/casbin/casbin/v2"
)

const (
	ErrorCodeInvalidURL          = "INVALID_URL"
	ErrorCodeInvalidJSON         = "INVALID_JSON"
	ErrorCodeInvalidParams       = "INVALID_PARAMS"
	ErrorCodeInternalServerError = "INTERNAL_SERVER_ERROR"
	ErrorCodeUnauthorized        = "UNAUTHORIZED"
	ErrorCodeAlreadyExists       = "ALREADY_EXISTS"
	ErrorCodeNotFound            = "NOT_FOUND"
	ErrorCodeInvalidCode         = "INVALID_CODE"
	ErrorBadRequest              = "BAD_REQUEST"
	ErrorInvalidCredentials      = "INVALID_CREDENTIALS"
)

type HandlerV1 struct {
	Config         *config.Config
	Logger         *zap.Logger
	ContextTimeout time.Duration
	Service        grpcClients.ServiceClient
	jwthandler     tokens.JWTHandler
	enforcer       *casbin.Enforcer
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Config         *config.Config
	Logger         *zap.Logger
	ContextTimeout time.Duration
	Service        grpcClients.ServiceClient
	JWTHandler     tokens.JWTHandler
	Enforcer       *casbin.Enforcer
}

// New ...
func New(c *HandlerV1Config) *HandlerV1 {
	return &HandlerV1{
		Config:         c.Config,
		Logger:         c.Logger,
		Service:        c.Service,
		ContextTimeout: c.ContextTimeout,
		jwthandler:     c.JWTHandler,
		enforcer:       c.Enforcer,
	}
}
