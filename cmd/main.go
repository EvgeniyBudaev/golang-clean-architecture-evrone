package main

import (
	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/app"
	"go.uber.org/zap"
)

func main() {
	app := app.NewApp()
	if err := app.StartHTTPServer(); err != nil {
		app.Logger.Fatal("error func main, method StartHTTPServer by path cmd/main.go", zap.Error(err))
	}
}
