package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/kluwena/go-api-practice/internal/order"
)

// Server represents the http server of order
type Server struct {
	router chi.Router

	orderService order.ServiceInterface
}

// NewServer create a new order http server
func NewServer(
	orderService order.ServiceInterface,
) *Server {
	s := &Server{
		router:       chi.NewRouter(),
		orderService: orderService,
	}
	s.buildRoutes()
	return s
}

func (s *Server) buildRoutes() {
	s.router.Post("/orders", func(w http.ResponseWriter, r *http.Request) {
		var createOrderParams order.CreateOrderParams
		if err := json.NewDecoder(r.Body).Decode(&createOrderParams); err != nil {
			w.Write([]byte("internal server error"))
			return
		}

		_, err := s.orderService.CreateOrder(r.Context(), &createOrderParams)
		if err != nil {
			log.Println(err)
		}
	})

	s.router.Get("/orders", func(w http.ResponseWriter, r *http.Request) {
		var listOrderParams order.ListOrdersParams
		orders, _, err := s.orderService.ListOrders(r.Context(), &listOrderParams)
		if err != nil {
			log.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(orders)

	})

}

// ServeHTTP serves the http requests
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
