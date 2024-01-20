package app

import (
	"database/sql"
	"fmt"
	roomHandler "github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/app/handlers/room"
	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/config"
	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/db"
	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/db/room"
	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/logger"
	roomUseCase "github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/useCase/room"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func Start() error {
	// Config
	cfg, err := config.Load()
	if err != nil {
		logger.Log.Debug("error func Start, method Load by path internal/app/app.go", zap.Error(err))
		return err
	}
	// Logging
	if err := logger.Initialize(cfg.LoggerLevel); err != nil {
		logger.Log.Debug("error func Start, method Initialize by path internal/app/app.go", zap.Error(err))
		return err
	}
	// Database
	databaseURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
		cfg.DBSSlMode)
	conn, err := sql.Open("postgres", databaseURL)
	if err != nil {
		logger.Log.Debug("error func Start, method Open by path internal/app/app.go", zap.Error(err))
		return err
	}
	database := db.NewDatabase(conn)
	err = conn.Ping()
	if err != nil {
		logger.Log.Debug("error func Start, method Ping by path internal/app/app.go", zap.Error(err))
		return err
	}
	dbRoom := room.NewPGRoomDB(database.GetDB())
	useCaseRoom := roomUseCase.NewUseCaseRoom(dbRoom)

	r := mux.NewRouter()
	r.Handle("/api/v1/room/create", roomHandler.CreateRoomHandler(useCaseRoom)).Methods("POST")
	r.Handle("/api/v1/room/list", roomHandler.GetRoomListHandler(useCaseRoom)).Methods("Get")

	err = http.ListenAndServe(cfg.Port, r)
	return err
}
