package api

import (
	"time"

	v1 "api-gateway/api/handlers/v1"
	"api-gateway/api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	grpcClients "api-gateway/internal/infrastructure/grpc_service_client"
	"api-gateway/internal/pkg/config"
	"api-gateway/internal/pkg/token"

)

type RouteOption struct {
	Config         *config.Config
	Logger         *zap.Logger
	ContextTimeout time.Duration
	Service        grpcClients.ServiceClient
	CasbinEnforcer *casbin.Enforcer
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

	jwtHandler := tokens.JWTHandler{
		SigninKey: option.Config.SigningKey,
	}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.NewAuthorizer(option.CasbinEnforcer, jwtHandler, *option.Config))

	HandlerV1 := v1.New(&v1.HandlerV1Config{
		Config:         option.Config,
		Logger:         option.Logger,
		ContextTimeout: option.ContextTimeout,
		Service:        option.Service,
	})

	api := router.Group("/v1")

	// products crud
	api.POST("/product/create", HandlerV1.CreateProduct)
	api.GET("/product/get/:id", HandlerV1.GetProduct)
	api.PUT("/product/update", HandlerV1.UpdateProduct)
	api.DELETE("/product/delete/:id", HandlerV1.DeleteProduct)
	api.GET("/products/get/:page/:limit", HandlerV1.ListProduct)

	// owner crud
	api.POST("/owner/create", HandlerV1.CreateOwner)
	api.GET("/owner/get/:id", HandlerV1.GetOwner)
	api.PUT("/owner/update", HandlerV1.UpdateOwner)
	api.DELETE("/owner/delete/:id", HandlerV1.DeleteOwner)
	api.GET("/owners/get/:page/:limit", HandlerV1.ListOwner)

	// geolocation crud
	api.POST("/geolocation/create", HandlerV1.CreateGeolocation)
	api.GET("/geolocation/get/:id", HandlerV1.GetGeolocation)
	api.PUT("/geolocation/update", HandlerV1.UpdateGeolocation)
	api.DELETE("/geolocation/delete/:id", HandlerV1.DeleteGeolocation)
	api.GET("/geolocations/get/:page/:limit", HandlerV1.ListGeolocation)

	// worker crud
	api.POST("/worker/create", HandlerV1.CreateWorker)
	api.GET("/worker/get/:id", HandlerV1.GetWorker)
	api.PUT("/worker/update", HandlerV1.UpdateWorker)
	api.DELETE("/worker/delete/:id", HandlerV1.DeleteWorker)
	api.GET("/workers/get/:page/:limit", HandlerV1.ListWorker)

	// order crud
	api.POST("/order/create", HandlerV1.CreateOrder)
	api.GET("/order/get/:id", HandlerV1.GetOrder)
	api.PUT("/order/update", HandlerV1.UpdateOrder)
	api.DELETE("/order/delete/:id", HandlerV1.DeleteOrder)
	api.GET("/orders/get/:page/:limit", HandlerV1.ListOrder)

	// registration
	api.POST("/register", HandlerV1.Register)
	api.GET("/verification", HandlerV1.Verify)
	api.GET("/login", HandlerV1.LogIn)

	// upload file
	api.POST("/file-upload", HandlerV1.UploadImage)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
