package v1

import (
	"time"

	"go.uber.org/zap"

	grpcClients "evrone_service/api_gateway/internal/infrastructure/grpc_service_client"
	"evrone_service/api_gateway/internal/pkg/config"
	// "evrone_service/api_gateway/internal/usecase/event"
)


type HandlerV1 struct {
	Config         *config.Config
	Logger         *zap.Logger
	ContextTimeout time.Duration
	Service        grpcClients.ServiceClient
	// BrokerProducer  event.BrokerProducer
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Config         *config.Config
	Logger         *zap.Logger
	ContextTimeout time.Duration
	Service        grpcClients.ServiceClient
	// BrokerProducer  event.BrokerProducer
}

// New ...
func New(c *HandlerV1Config) *HandlerV1 {
	return &HandlerV1{
	  Config: c.Config,
	  Logger: c.Logger,
	  Service: c.Service,
	  ContextTimeout: c.ContextTimeout,
	//   BrokerProducer: c.BrokerProducer,
	}
}
