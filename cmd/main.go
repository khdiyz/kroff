package main

import (
	"context"
	"kroff"
	"kroff/config"
	"kroff/pkg/handler"
	"kroff/pkg/repository"
	"kroff/pkg/service"
	"kroff/pkg/storage"
	"kroff/utils/logger"
	"os"
	"os/signal"
	"syscall"
)

// @title Kroff API
// @version 1.0
// @description API Server for Kroff
// @host localhost:4040
// @BasePath
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	logger := logger.GetLogger()

	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	storage, err := storage.NewStorage(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewServices(repos, storage, cfg)
	handlers := handler.NewHandlers(services, logger)

	srv := new(kroff.Server)
	go func() {
		if err := srv.Run(cfg.HTTPHost, cfg.HTTPPort, handlers.InitRoutes()); err != nil {
			logger.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logger.Info("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Warn("App shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logger.Errorf("error occured on db connection close: %s", err.Error())
	}
}
