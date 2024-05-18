// package config

// import (
// 	"os"
// 	// "time"

// 	"github.com/spf13/cast"
// )

// const (
// 	OtpSecret = "some_secret"
// )

// type webAddress struct {
// 	Host string
// 	Port string
// }

// type Config struct {
// 	APP         string
// 	Environment string
// 	LogLevel    string
// 	Server      struct {
// 		Host         string
// 		Port         string
// 		ReadTimeout  string
// 		WriteTimeout string
// 		IdleTimeout  string
// 	}

// 	Context struct {
// 		Timeout string
// 	}

// 	Redis struct {
// 		Host     string
// 		Port     string
// 		Password string
// 		Name     string
// 	}

// 	// context timeout in seconds
// 	CtxTimeout int

// 	// Jwt
// 	SigningKey        string
// 	AccessTokenTimout int

// 	// casbin
// 	AuthConfigPath string
// 	CSVFilePath    string

// 	// services
// 	UserService    webAddress
// 	ProductService webAddress
// 	OrderService   webAddress

// 	// otlp
// 	OTLPCollector webAddress
// }

// func NewConfig() Config {
// 	config := Config{}

// 	// general configuration
// 	config.APP = getEnv("APP", "app")
// 	config.Environment = getEnv("ENVIRONMENT", "develop")
// 	config.LogLevel = getEnv("LOG_LEVEL", "debug")
// 	config.Context.Timeout = getEnv("CONTEXT_TIMEOUT", "30s")

// 	// server configuration
// 	config.Server.Host = getEnv("SERVER_HOST", "127.0.0.1")
// 	config.Server.Port = getEnv("SERVER_PORT", ":8080")
// 	config.Server.ReadTimeout = getEnv("SERVER_READ_TIMEOUT", "10s")
// 	config.Server.WriteTimeout = getEnv("SERVER_WRITE_TIMEOUT", "10s")
// 	config.Server.IdleTimeout = getEnv("SERVER_IDLE_TIMEOUT", "120s")

// 	config.SigningKey = cast.ToString(getEnv("SIGNING_KEY", "test-key"))
// 	config.AccessTokenTimout = cast.ToInt(getEnv("ACCESS_TOKEN_TIMOUT", "30000"))

// 	// redis configuration
// 	config.Redis.Host = getEnv("REDIS_HOST", "redis")
// 	config.Redis.Port = getEnv("REDIS_PORT", "6379")
// 	config.Redis.Password = getEnv("REDIS_PASSWORD", "")
// 	config.Redis.Name = getEnv("REDIS_DATABASE", "0")

// 	// USER
// 	config.UserService.Host = getEnv("USER_SERVICE_GRPC_HOST", "127.0.0.1")
// 	config.UserService.Port = getEnv("USER_SERVICE_GRPC_PORT", ":2222")

// 	// PRODUCT
// 	config.ProductService.Host = getEnv("PRODUCT_SERVICE_GRPC_HOST", "127.0.0.1")
// 	config.ProductService.Port = getEnv("PRODUCT_SERVICE_GRPC_PORT", ":1111")

// 	// ORDER
// 	config.OrderService.Host = getEnv("ORDER_SERVICE_GRPC_HOST", "127.0.0.1")
// 	config.OrderService.Port = getEnv("ORDER_SERVICE_GRPC_PORT", ":3333")

// 	// casbin
// 	config.AuthConfigPath = cast.ToString(getEnv("AUTH_CONFIG_PATH", "auth.conf"))
// 	config.CSVFilePath = cast.ToString(getEnv("CSV_FILE_PATH", "auth.csv"))

// 	// otlp collector configuration
// 	config.OTLPCollector.Host = getEnv("OTLP_COLLECTOR_HOST", "otel-collector")
// 	config.OTLPCollector.Port = getEnv("OTLP_COLLECTOR_PORT", ":4317")

// 	config.CtxTimeout = cast.ToInt(getEnv("CTX_TIMEOUT", "7"))

// 	return config
// }

// func getEnv(key, defaultValue string) string {
// 	if value, exists := os.LookupEnv(key); exists {
// 		return value
// 	}

// 	return defaultValue
// }

package config

import (
	"os"
	// "time"

	"github.com/spf13/cast"
)

const (
	OtpSecret = "some_secret"
)

type webAddress struct {
	Host string
	Port string
}

type Config struct {
	APP         string
	Environment string
	LogLevel    string
	Server      struct {
		Host         string
		Port         string
		ReadTimeout  string
		WriteTimeout string
		IdleTimeout  string
	}

	Context struct {
		Timeout string
	}

	Redis struct {
		Host     string
		Port     string
		Password string
		Name     string
	}

	// context timeout in seconds
	CtxTimeout int

	// Jwt
	SigningKey        string
	AccessTokenTimout int

	// casbin
	AuthConfigPath string
	CSVFilePath    string

	// services
	UserService    webAddress
	ProductService webAddress
	OrderService   webAddress

	// otlp
	OTLPCollector webAddress
}

func NewConfig() Config {
	config := Config{}

	// general configuration
	config.APP = getEnv("APP", "app")
	config.Environment = getEnv("ENVIRONMENT", "develop")
	config.LogLevel = getEnv("LOG_LEVEL", "debug")
	config.Context.Timeout = getEnv("CONTEXT_TIMEOUT", "30s")

	// server configuration
	config.Server.Host = getEnv("SERVER_HOST", "")
	config.Server.Port = getEnv("SERVER_PORT", ":8080")
	config.Server.ReadTimeout = getEnv("SERVER_READ_TIMEOUT", "10s")
	config.Server.WriteTimeout = getEnv("SERVER_WRITE_TIMEOUT", "10s")
	config.Server.IdleTimeout = getEnv("SERVER_IDLE_TIMEOUT", "120s")

	config.SigningKey = cast.ToString(getEnv("SIGNING_KEY", "test-key"))
	config.AccessTokenTimout = cast.ToInt(getEnv("ACCESS_TOKEN_TIMOUT", "30000"))

	// redis configuration
	config.Redis.Host = getEnv("REDIS_HOST", "redis")
	config.Redis.Port = getEnv("REDIS_PORT", "6379")
	config.Redis.Password = getEnv("REDIS_PASSWORD", "")
	config.Redis.Name = getEnv("REDIS_DATABASE", "0")

	// USER
	config.UserService.Host = getEnv("USER_SERVICE_GRPC_HOST", "user-service")
	config.UserService.Port = getEnv("USER_SERVICE_GRPC_PORT", ":2222")

	// PRODUCT
	config.ProductService.Host = getEnv("PRODUCT_SERVICE_GRPC_HOST", "product-service")
	config.ProductService.Port = getEnv("PRODUCT_SERVICE_GRPC_PORT", ":1111")

	// ORDER
	config.OrderService.Host = getEnv("ORDER_SERVICE_GRPC_HOST", "order-service")
	config.OrderService.Port = getEnv("ORDER_SERVICE_GRPC_PORT", ":3333")

	// casbin
	config.AuthConfigPath = cast.ToString(getEnv("AUTH_CONFIG_PATH", "auth.conf"))
	config.CSVFilePath = cast.ToString(getEnv("CSV_FILE_PATH", "auth.csv"))

	// otlp collector configuration
	config.OTLPCollector.Host = getEnv("OTLP_COLLECTOR_HOST", "otel-collector")
	config.OTLPCollector.Port = getEnv("OTLP_COLLECTOR_PORT", ":4317")

	config.CtxTimeout = cast.ToInt(getEnv("CTX_TIMEOUT", "7"))

	return config
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
