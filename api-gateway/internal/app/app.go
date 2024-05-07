package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/casbin/casbin/v2"
	defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	"go.uber.org/zap"

	"api-gateway/api"
	grpcService "api-gateway/internal/infrastructure/grpc_service_client"
	"api-gateway/internal/pkg/config"
	"api-gateway/internal/pkg/logger"
	"api-gateway/internal/pkg/otlp"

	"github.com/casbin/casbin/v2/util"
)

type App struct {
	Config *config.Config
	Logger *zap.Logger

	server       *http.Server
	Enforcer     *casbin.Enforcer
	Clients      grpcService.ServiceClient
	ShutdownOTLP func() error
}

func NewApp(cfg config.Config) (*App, error) {
	// logger init
	logger, err := logger.New(cfg.LogLevel, cfg.Environment, cfg.APP+".log")
	if err != nil {
		return nil, err
	}

	// csv file bn accses berish
	casbinEnforcer, err := casbin.NewEnforcer(cfg.AuthConfigPath, cfg.CSVFilePath)
	if err != nil {
		log.Fatal("casbin enforcer error", err)
		return nil, err
	}

	// otlp collector init
	shutdownOTLP, err := otlp.InitOTLPProvider(&cfg)
	if err != nil {
		return nil, err
	}

	return &App{
		Config:       &cfg,
		Logger:       logger,
		Enforcer:     casbinEnforcer,
		ShutdownOTLP: shutdownOTLP,
	}, nil
}

func (a *App) Run() error {
	contextTimeout, err := time.ParseDuration(a.Config.Context.Timeout)
	if err != nil {
		return fmt.Errorf("error while parsing context timeout: %v", err)
	}

	clients, err := grpcService.New(a.Config)
	if err != nil {
		return err
	}
	a.Clients = clients

	// api init
	handler := api.NewRoute(api.RouteOption{
		Config:         a.Config,
		Logger:         a.Logger,
		ContextTimeout: contextTimeout,
		CasbinEnforcer: a.Enforcer,
		Service:        clients,
	})

	// casbin
	err = a.Enforcer.LoadPolicy()
	if err != nil {
		log.Fatal("casbin error load policy", err)
		return err
	}
	a.Enforcer.GetRoleManager().(*defaultrolemanager.RoleManagerImpl).AddMatchingFunc("keyMatch", util.KeyMatch)
	a.Enforcer.GetRoleManager().(*defaultrolemanager.RoleManagerImpl).AddMatchingFunc("keyMatch3", util.KeyMatch3)

	// server init
	a.server, err = api.NewServer(a.Config, handler)
	if err != nil {
		return fmt.Errorf("error while initializing server: %v", err)
	}

	return a.server.ListenAndServe()
}

func (a *App) Stop() {

	// close grpc connections
	a.Clients.Close()

	// shutdown server http
	if err := a.server.Shutdown(context.Background()); err != nil {
		a.Logger.Error("shutdown server http ", zap.Error(err))
	}

	// shutdown otlp collector
	if err := a.ShutdownOTLP(); err != nil {
		a.Logger.Error("shutdown otlp collector", zap.Error(err))
	}

	// zap logger sync
	a.Logger.Sync()
}
