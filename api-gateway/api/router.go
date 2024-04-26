package api

import (
	// "github.com/casbin/casbin/v2"
	"time"

	v1 "evrone_service/api_gateway/api/handlers/v1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	grpcClients "evrone_service/api_gateway/internal/infrastructure/grpc_service_client"
	"evrone_service/api_gateway/internal/pkg/config"
	// "evrone_service/api_gateway/internal/usecase/event"
)

type RouteOption struct {
	Config         *config.Config
	Logger         *zap.Logger
	ContextTimeout time.Duration
	Service        grpcClients.ServiceClient
	// BrokerProducer event.BrokerProducer
}

// NewRoute
// @Title WCRM
// @Version 1.0
// @Description wcrm-system
// @Host localhost:8080
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewRoute(option RouteOption) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	HandlerV1 := v1.New(&v1.HandlerV1Config{
		Config:         option.Config,
		Logger:         option.Logger,
		ContextTimeout: option.ContextTimeout,
		Service:        option.Service,
		// BrokerProducer: option.BrokerProducer,
	})

	api := router.Group("/v1")

	// products crud
	api.POST("/product/create", HandlerV1.CreateProduct)
	api.GET("/product/get/:id", HandlerV1.GetProduct)
	api.PUT("/product/update", HandlerV1.UpdateProduct)
	api.DELETE("/product/delete/:id", HandlerV1.DeleteProduct)
	api.GET("/products/get", HandlerV1.ListProduct)

	// owner crud
	api.POST("/owner/create", HandlerV1.CreateOwner)
	api.GET("/owner/get/:id", HandlerV1.GetOwner)
	api.PUT("/owner/update", HandlerV1.UpdateOwner)
	api.DELETE("/owner/delete/:id", HandlerV1.DeleteOwner)
	api.GET("/owners/get", HandlerV1.ListOwner)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
