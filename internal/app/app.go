package app

import (
	"context"
	"os"
	"os/signal"
	"shorty/internal/config"
	"shorty/internal/repository"
	"shorty/internal/service"
	"shorty/internal/transport/http"
	"shorty/pkg/db/redis"
	"shorty/pkg/httpserver"
	"shorty/pkg/logger"
	"syscall"
	"time"
)

const timeout = 5 * time.Second

// @title Shorty API
// @version 1.0
// @description API Server for url Shorty app

// @host localhost:8000
// @BasePath /

func Run() {
	cfg := config.GetConfig()

	db := redis.NewRedis(redis.Config{
		Host:     cfg.Redis.Host,
		Port:     cfg.Redis.Port,
		Password: cfg.Redis.Password,
	})
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := http.NewHandler(services)
	srv := httpserver.NewServer(cfg.Server.Port, handler.Init())

	go func() {
		if err := srv.Run(); err != nil {
			logger.Errorf("error during running http server: %s\n", err.Error())
			return
		}
	}()

	graceful()
	shutdown(srv, db)
}

func graceful() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func shutdown(srv *httpserver.Server, redis *redis.Client) {
	ctx, shutdownCancel := context.WithTimeout(context.Background(), timeout)
	defer shutdownCancel()

	if err := redis.Redis.Close(); err != nil {
		logger.Errorf("failed to disconnect redis client: %s\n", err.Error())
		return
	}

	if err := srv.Shutdown(ctx); err != nil {
		logger.Errorf("failed to stop server: %s\n", err.Error())
		return
	}
}
