package app

import (
	"log"

	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/config"
	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/logger"
	"go.uber.org/zap"
)

type App struct {
	Logger    logger.Logger
	config    *config.Config
	container *Container
}

func NewApp() *App {
	default_logger, err := logger.NewLogger(logger.DEFAULT_LEVEL)
	if err != nil {
		log.Fatal("error func Start, method NewLogger by path internal/app/app.go", err)
		return nil
	}

	cfg, err := config.Load(default_logger)
	if err != nil {
		log.Fatal("error func Start, method Load by path internal/app/app.go", err)
		return nil
	}

	logger, err := logger.NewLogger(cfg.LoggerLevel)
	if err != nil {
		log.Fatal("error func Start, method NewLogger by path internal/app/app.go", err)
		return nil
	}

	postgresConnection, err := newPostgresConnection(cfg)
	if err != nil {
		logger.Fatal("error func Start, method newPostgresConnection by path internal/app/app.go", zap.Error(err))
	}

	container := NewContainer(logger, postgresConnection)

	return &App{
		config:    cfg,
		Logger:    logger,
		container: container,
	}
}
