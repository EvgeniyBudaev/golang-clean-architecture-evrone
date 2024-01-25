package http

import (
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	srv *http.Server
}

func NewServer(host string, port string) *Server {
	server := &Server{
		&http.Server{
			Addr: fmt.Sprintf("%s:%s", host, port),
		},
	}

	return server
}

func (s *Server) RegisterRoutes(r *Router) {
	s.srv.Handler = r.router
}

func (s *Server) Start() error {
	if s.srv.Handler == nil {
		return fmt.Errorf("no routes have registered")
	}

	err := s.srv.ListenAndServe()
	if err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	return s.srv.Shutdown(context.Background())
}
