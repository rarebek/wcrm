package api

import (
	"time"

	v1 "api-gateway/api/handlers/v1"
	"api-gateway/api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	grpcClients "api-gateway/internal/infrastructure/grpc_service_client"
	"api-gateway/internal/pkg/config"
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
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewRoute(option RouteOption) *gin.Engine {
	router := gin.New()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowBrowserExtensions = true
	corsConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corsConfig))

	// jwtHandler := tokens.JWTHandler{
	// SigninKey: option.Config.SigningKey,
	// }

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// router.Use(middleware.NewAuthorizer(option.CasbinEnforcer, jwtHandler, *option.Config))

	HandlerV1 := v1.New(&v1.HandlerV1Config{
		Config:         option.Config,
		Logger:         option.Logger,
		ContextTimeout: option.ContextTimeout,
		Service:        option.Service,
	})

	router.Use(middleware.Tracing)
	api := router.Group("/v1")

	// // category
	api.POST("/category/create", HandlerV1.CreateCategory)
	api.GET("/category/get/:id", HandlerV1.GetCategory)
	api.PUT("/category/update", HandlerV1.UpdateCategory)
	api.DELETE("/category/delete/:id", HandlerV1.DeleteCategory)
	api.POST("/category/getall", HandlerV1.ListCategory)

	// products
	api.POST("/product/create", HandlerV1.CreateProduct)
	api.GET("/product/get/:id", HandlerV1.GetProduct)
	api.PUT("/product/update", HandlerV1.UpdateProduct)
	api.DELETE("/product/delete/:id", HandlerV1.DeleteProduct)
	api.GET("/products/get/:page/:limit/:owner-id", HandlerV1.ListProduct)
	api.POST("/product/search", HandlerV1.SearchProduct)
	api.POST("/products/get_category_id", HandlerV1.GetAllProductByCategoryId)

	// owner
	api.POST("/owner/create", HandlerV1.CreateOwner)
	api.GET("/owner/get/:id", HandlerV1.GetOwner)
	api.PUT("/owner/update", HandlerV1.UpdateOwner)
	api.DELETE("/owner/delete/:id", HandlerV1.DeleteOwner)
	api.GET("/owners/get/:page/:limit", HandlerV1.ListOwner)

	// geolocation
	api.POST("/geolocation/create", HandlerV1.CreateGeolocation)
	api.GET("/geolocation/get/:id", HandlerV1.GetGeolocation)
	api.PUT("/geolocation/update", HandlerV1.UpdateGeolocation)
	api.DELETE("/geolocation/delete/:id", HandlerV1.DeleteGeolocation)
	api.GET("/geolocations/get/:page/:limit", HandlerV1.ListGeolocation)

	// worker
	api.POST("/worker/create", HandlerV1.CreateWorker)
	api.GET("/worker/get/:id", HandlerV1.GetWorker)
	api.PUT("/worker/update", HandlerV1.UpdateWorker)
	api.DELETE("/worker/delete/:id", HandlerV1.DeleteWorker)
	api.GET("/workers/get/:page/:limit/:owner-id", HandlerV1.ListWorker)

	// order
	api.POST("/order/create", HandlerV1.CreateOrder)
	api.GET("/order/get/:id", HandlerV1.GetOrder)
	api.PUT("/order/update", HandlerV1.UpdateOrder)
	api.DELETE("/order/delete/:id", HandlerV1.DeleteOrder)
	api.GET("/orders/get/:page/:limit/:worker-id", HandlerV1.ListOrder)
	api.GET("/products/bot", HandlerV1.GetProducts)
	api.DELETE("/products/bot", HandlerV1.DeleteProductBot)

	//
	api.POST("/register", HandlerV1.Register)
	api.GET("/verification", HandlerV1.Verify)
	api.GET("/login", HandlerV1.LogIn)
	api.POST("/worker/login", HandlerV1.LogInWorker)

	// upload file
	api.POST("/file-upload", HandlerV1.UploadImage)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
