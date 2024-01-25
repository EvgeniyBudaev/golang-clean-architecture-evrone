package http

import (
	"net/http"

	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/handler/middleware"
	"github.com/EvgeniyBudaev/golang-clean-architecture-evrone/internal/logger"
	"github.com/gorilla/mux"
)

type HandlerRouter interface {
	AddRoutes(r *mux.Router)
	GetVersion() string
	GetContentType() string
}

type Router struct {
	router *mux.Router
}

func NewRouter() *Router {
	return &Router{router: mux.NewRouter()}
}

func (r *Router) WithHandler(h HandlerRouter, logger logger.Logger) *Router {
	api := r.router.PathPrefix("/api/" + h.GetVersion()).Subrouter()
	r.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Kekpek")) })

	api.Use(middleware.AddContextMiddleware(logger))

	h.AddRoutes(api)

	return r
}
