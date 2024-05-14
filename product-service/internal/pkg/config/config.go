package config

import (
	"os"

	"github.com/spf13/cast"
)

type webAddress struct {
	Host string
	Port int
}

type Config struct {
	APP         string
	Environment string
	LogLevel    string
	RPCPort     string

	Context struct {
		Timeout string
	}

	DB struct {
		Host     string
		Port     string
		Name     string
		User     string
		Password string
		SslMode  string
	}

	OTLPCollector struct {
		Host string
		Port string
	}

	// services
	OrderService webAddress
}

func New() *Config {
	var config Config

	// general configuration
	config.APP = getEnv("APP", "app")
	config.Environment = getEnv("ENVIRONMENT", "develop")
	config.LogLevel = getEnv("LOG_LEVEL", "debug")
	config.RPCPort = getEnv("RPC_PORT", ":1111")
	config.Context.Timeout = getEnv("CONTEXT_TIMEOUT", "30s")

	// db configuration
	config.DB.Host = getEnv("POSTGRES_HOST", "db")
	config.DB.Port = getEnv("POSTGRES_PORT", "5432")
	config.DB.User = getEnv("POSTGRES_USER", "postgres")
	config.DB.Password = getEnv("POSTGRES_PASSWORD", "nodirbek")
	config.DB.SslMode = getEnv("POSTGRES_SSLMODE", "disable")
	config.DB.Name = getEnv("POSTGRES_DATABASE", "productdb")

	// otlp collector configuration
	config.OTLPCollector.Host = getEnv("OTLP_COLLECTOR_HOST", "127.0.0.1")
	config.OTLPCollector.Port = getEnv("OTLP_COLLECTOR_PORT", ":4317")

	// connect with order service
	config.OrderService.Host = cast.ToString(getEnv("ORDER_SERVICE_HOST", "127.0.0.1"))
	config.OrderService.Port = cast.ToInt(getEnv("ORDER_SERVICE_PORT", "3333"))

	return &config
}

func getEnv(key string, defaultVaule string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultVaule
}
