package room

import (
	"context"
	"encoding/json"
	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/entity/room"
	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/logger"
	roomUseCase "github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/useCase/room"
	"go.uber.org/zap"
	"net/http"
)

type UseCaseRoom interface {
	CreateRoom(ctx context.Context, r roomUseCase.CreateRoomRequest) (*room.Room, error)
	GetRoomList(ctx context.Context) ([]*room.Room, error)
}

func CreateRoomHandler(uc UseCaseRoom) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Log.Info("POST /api/v1/room/create")
		req := roomUseCase.CreateRoomRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			logger.Log.Debug(
				"error func CreateRoomHandler, method BodyDecoder by path internal/handlers/room/room.go",
				zap.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response, err := uc.CreateRoom(r.Context(), req)
		if err != nil {
			logger.Log.Debug(
				"error func CreateRoomHandler, method uc.CreateRoom by path internal/handlers/room/room.go",
				zap.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	})
}

func GetRoomListHandler(uc UseCaseRoom) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Log.Info("GET /api/v1/room/list")
		response, err := uc.GetRoomList(r.Context())
		if err != nil {
			logger.Log.Debug("error func GetRoomListHandler,"+
				" method uc.GetRoomList by path internal/handlers/room/room.go",
				zap.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	})
}
