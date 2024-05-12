package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"wcrm/product-service/internal/app"
	"wcrm/product-service/internal/pkg/config"

	"go.uber.org/zap"
)

func main() {
	// initialization config
	config := config.New()

	// initialization app
	app, err := app.NewApp(config)
	if err != nil {
		log.Fatal(err)
	}

	// runing
	go func() {
		if err := app.Run(); err != nil {
			app.Logger.Error("app run", zap.Error(err))
		}
	}()

	// graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	app.Logger.Info("Conten service stops !")

	// app stops
	app.Stop()

}
