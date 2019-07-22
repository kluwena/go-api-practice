package http

import (
	"github.com/go-chi/chi"
	"encoding/json"
	// "log"
	// "net/http"
	// "context"
	"github.com/kluwena/go-api-practice/internal/order"
)

// Server represents the http server of order
type Server struct {
	router chi.Router

	orderService order.ServiceInterface
}


// NewServer create a new order http server
func NewServer(
	orderService order.serviceInterface,
) *Server {
	s:= &Server{
		router: chi.NewRouter(),
		orderService: orderService,
	}
	s.buildRoutes()
	return s
}

// ServeHTTP serves the http requests
func (s *Server) serveHTTP (w http.ResponseWriter r *http.Request) {
	s.router.serveHTTP(w, r)
}