package app

import (
	"fmt"
	"user-service/internal/delivery/kafka/handlers"
	"user-service/internal/infrastructure/kafka"
	"user-service/internal/infrastructure/repository/postgresql"
	"user-service/internal/pkg/config"
	logPkg "user-service/internal/pkg/logger"
	"user-service/internal/pkg/postgres"
	"user-service/internal/usecase"
	"user-service/internal/usecase/event"

	"go.uber.org/zap"
)

type UserCreateConsumerCLI struct {
	Config         *config.Config
	Logger         *zap.Logger
	DB             *postgres.PostgresDB
	BrokerConsumer event.BrokerConsumer
}

func NewUserCreateConsumerCLI(config *config.Config, logger *zap.Logger, db *postgres.PostgresDB, brokerConsumer event.BrokerConsumer) (*UserCreateConsumerCLI, error) {
	logger, err := logPkg.New(config.LogLevel, config.Environment, config.APP+"_cli"+".log")
	if err != nil {
		return nil, err
	}

	consumer := kafka.NewConsumer(logger)

	db, err = postgres.New(config)
	if err != nil {
		return nil, err
	}

	return &UserCreateConsumerCLI{
		Config:         config,
		DB:             db,
		Logger:         logger,
		BrokerConsumer: consumer,
	}, nil
}

func (c *UserCreateConsumerCLI) Run() error {
	fmt.Print("consume is running ....")
	// repo init
	userRepo := postgresql.NewOwnersRepo(c.DB)

	// usecase init
	userUsecase := usecase.NewOwnerService(c.DB.Config().ConnConfig.ConnectTimeout, userRepo)

	eventHandler := handlers.NewUserCreateHandler(c.Config, c.BrokerConsumer, c.Logger, userUsecase)

	return eventHandler.HandlerEvents()
}

func (c *UserCreateConsumerCLI) Close() {
	c.BrokerConsumer.Close()

	c.Logger.Sync()
}
