package app

import (
	"fmt"
	"time"
	userproto "user-service/genproto/user_service"
	grpc_server "user-service/internal/delivery/grpc/server"
	clean_grpc "user-service/internal/delivery/grpc/services"
	"user-service/internal/infrastructure/grpc_service_clients"
	"user-service/internal/infrastructure/kafka"
	repo "user-service/internal/infrastructure/repository/postgresql"
	"user-service/internal/pkg/config"
	"user-service/internal/pkg/logger"
	"user-service/internal/pkg/postgres"
	"user-service/internal/usecase"
	"user-service/internal/usecase/event"
	"user-service/internal/pkg/otlp"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	Config         *config.Config
	Logger         *zap.Logger
	DB             *postgres.PostgresDB
	GrpcServer     *grpc.Server
	ShutdownOTLP   func() error
	ServiceClients grpc_service_clients.ServiceClients
	BrokerProducer event.BrokerProducer
	BrokerConsumer event.BrokerConsumer
}

func NewApp(cfg *config.Config) (*App, error) {
	// init logger
	logger, err := logger.New(cfg.LogLevel, cfg.Environment, cfg.APP+".log")
	if err != nil {
		return nil, err
	}

	// kafaka init
	kafkaProducer := kafka.NewProducer(cfg, logger)
	kafkaConsumer := kafka.NewConsumer(logger)

	// otlp collector initialization
	shutdownOTLP, err := otlp.InitOTLPProvider(cfg)
	if err != nil {
		return nil, err
	}

	// init db
	db, err := postgres.New(cfg)
	if err != nil {
		return nil, err
	}

	consumerApp, err := NewUserCreateConsumerCLI(cfg, logger, db, kafkaConsumer)

	// grpc server init
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(logger),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_server.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_ctxtags.UnaryServerInterceptor(),
				grpc_zap.UnaryServerInterceptor(logger),
				grpc_recovery.UnaryServerInterceptor(),
			),
			grpc_server.UnaryInterceptorData(logger),
		)),
	)

	return &App{
		Config:         cfg,
		Logger:         logger,
		DB:             db,
		GrpcServer:     grpcServer,
		ShutdownOTLP:   shutdownOTLP,
		BrokerConsumer: consumerApp.BrokerConsumer,
		BrokerProducer: kafkaProducer,
	}, nil
}

func (a *App) Run() error {
	var (
		contextTimeout time.Duration
	)

	// context timeout initialization
	contextTimeout, err := time.ParseDuration(a.Config.Context.Timeout)
	if err != nil {
		return fmt.Errorf("error during parse duration for context timeout : %w", err)
	}
	// Initialize Service Clients
	serviceClients, err := grpc_service_clients.New(a.Config)
	if err != nil {
		return fmt.Errorf("error during initialize service clients: %w", err)
	}
	a.ServiceClients = serviceClients

	// repositories initialization
	ownerRepo := repo.NewOwnersRepo(a.DB)
	workerRepo := repo.NewWorkersRepo(a.DB)
	geolocationRepo := repo.NewGeolocationsRepo(a.DB)

	// usecase initialization
	ownerUsecase := usecase.NewOwnerService(contextTimeout, ownerRepo)
	workerUsecase := usecase.NewWorkerService(contextTimeout, workerRepo)
	geolocationUsecase := usecase.NewGeolocationService(contextTimeout, geolocationRepo)

	userproto.RegisterUserServiceServer(a.GrpcServer, clean_grpc.NewRPC(a.Logger, ownerUsecase, workerUsecase, geolocationUsecase, a.BrokerProducer))
	a.Logger.Info("gRPC Server Listening", zap.String("url", a.Config.RPCPort))
	if err := grpc_server.Run(a.Config, a.GrpcServer); err != nil {
		return fmt.Errorf("gRPC fatal to serve grpc server over %s %w", a.Config.RPCPort, err)
	}

	return nil
}

func (a *App) Stop() {
	// close broker producer
	a.BrokerProducer.Close()

	a.BrokerConsumer.Close()

	// closing client service connections
	a.ServiceClients.Close()
	// stop gRPC server
	a.GrpcServer.Stop()

	// database connection
	a.DB.Close()

	// shutdown otlp collector
	if err := a.ShutdownOTLP(); err != nil {
		a.Logger.Error("shutdown otlp collector", zap.Error(err))
	}

	// zap logger sync
	a.Logger.Sync()
}
