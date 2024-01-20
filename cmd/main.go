package main

import (
	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/app"
	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/logger"
	"go.uber.org/zap"
	"log"
)

func main() {
	if err := app.Start(); err != nil {
		logger.Log.Debug("error func main, method Start by path cmd/main.go", zap.Error(err))
		log.Fatal(err)
	}
}
