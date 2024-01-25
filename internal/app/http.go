package app

import (
	"fmt"

	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/handler/http"
	v1 "github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/handler/http/api/v1"
	"github.com/hanagantig/gracy"
	"go.uber.org/zap"
)

func (app *App) StartHTTPServer() error {
	go func() {
		app.Logger.Info(fmt.Sprintf("Starting http server on port %s", app.config.Port))

		server := http.NewServer("0.0.0.0", app.config.Port)
		router := http.NewRouter().WithHandler(v1.NewHandler(app.Logger, app.container.GetUseCase()), app.Logger)
		server.RegisterRoutes(router)

		gracy.AddCallback(func() error {
			return server.Stop()
		})

		err := server.Start()
		if err != nil {
			app.Logger.Fatal("Failed to start http server", zap.Error(err))
		}
	}()

	err := gracy.Wait()
	if err != nil {
		app.Logger.Error("failed to gracefully shutdown server", zap.Error(err))
		return err
	}

	return nil
}
