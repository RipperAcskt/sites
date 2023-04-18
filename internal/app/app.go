package app

import (
	"fmt"
	"net/http"

	"sites/config"
	"sites/internal/handler"
	"sites/internal/repo/postgres"
	"sites/internal/server"
	"sites/internal/service"

	"go.uber.org/zap"

	"github.com/golang-migrate/migrate/v4"
)

func Run() error {
	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("config new failed: %w", err)
	}

	postgres, err := postgres.New(cfg)
	if err != nil {
		return fmt.Errorf("postgres new failed: %w", err)
	}
	defer postgres.Close()

	err = postgres.Migrate.Up()
	if err != migrate.ErrNoChange && err != nil {
		return fmt.Errorf("migrate up failed: %w", err)
	}

	log, err := zap.NewProduction()
	if err != nil {
		return fmt.Errorf("zap new failed: %w", err)
	}

	defer func() {
		err := log.Sync()
		if err != nil {
			log.Error("log sync failed: %w", zap.Error(err))
		}
	}()

	service := service.New(postgres, cfg.SALT, cfg)
	handler := handler.New(service, cfg, log)
	server := &server.Server{
		Log: log,
	}

	go func() {
		if err := server.Run(handler.InitRouters(), cfg); err != nil && err != http.ErrServerClosed {
			log.Error(fmt.Sprintf("server run failed: %v", err))
			return
		}
	}()

	if err := server.ShutDown(); err != nil {
		return fmt.Errorf("server shut down failed: %w", err)
	}
	return nil
}
