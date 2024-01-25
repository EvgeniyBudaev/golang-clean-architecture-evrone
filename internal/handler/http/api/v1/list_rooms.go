package v1

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

func (h *Handler) ListRooms(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("GET /api/v1/room/list")

	w.Header().Set("Content-Type", h.GetContentType())

	rooms, err := h.useCase.ListRooms(r.Context())
	if err != nil {
		h.logger.Debug("error func ListRooms, method uc.ListRooms by path internal/handler/http/api/v1/list_rooms.go", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rooms)
}
