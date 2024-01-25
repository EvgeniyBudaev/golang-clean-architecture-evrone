package v1

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type createRoomRequest interface {
	GetRoomName() string
	GetTitle() string
}

type createRoomRequestImpl struct {
	RoomName string `json:"room_name"`
	Title    string `json:"title"`
}

func (r *createRoomRequestImpl) GetRoomName() string {
	return r.RoomName
}

func (r *createRoomRequestImpl) GetTitle() string {
	return r.Title
}

func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("POST /api/v1/room/create")

	w.Header().Set("Content-Type", h.GetContentType())

	req := &createRoomRequestImpl{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		h.logger.Debug("error func CreateRoom, method BodyDecoder by path internal/handler/http/api/v1/create_room.go", zap.Error(err))

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	room, err := h.useCase.CreateRoom(r.Context(), req)
	if err != nil {
		h.logger.Debug("error func CreateRoom, method uc.CreateRoom by path internal/handler/http/api/v1/create_room.go", zap.Error(err))

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(room)
}
