package v1

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) GetVersion() string {
	return "v1"
}

func (h *Handler) GetContentType() string {
	return "application/json"
}

func (h *Handler) AddRoutes(r *mux.Router) {
	r.HandleFunc("/room/create", h.CreateRoom).Methods(http.MethodPost)
	r.HandleFunc("/room/list", h.ListRooms).Methods(http.MethodGet)
}
