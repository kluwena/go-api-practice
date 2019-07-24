package order

import (
	"context"
	"log"
	"time"
)

// Order represent the order entity
type Order struct {
	ID              int
	TransactionTime time.Time
}

// CreateOrderParamsRequest request for Order Request
type CreateOrderParamsRequest struct {
	Title           string    `json: "title"`
	TransactionTime time.Time `json: "transactionTime"`
}

// ListOrdersParams is params for list orders
type ListOrdersParams struct {
	TimeNow         *time.Time
	TransactionTime string
}

// ServiceInterface specifies the interface of order service.
type ServiceInterface interface {
	CreateOrder(ctx context.Context, params *CreateOrderParamsRequest) (*Order, error)
	ListOrders(ctx context.Context, params *ListOrdersParams) ([]*Order, error)
}

// Repository represents the interface for order entity
type Repository interface {
	Insert(ctx context.Context, order *Order) error
	FindAll(ctx context.Context, params *ListOrdersParams) ([]*Order, error)
	CountAll(ctx context.Context, params *ListOrdersParams) (int, error)
}

// Service represents the implementation details of order service interface
type Service struct {
	orderRepository Repository
}

// NewService creates a new service
func NewService(
	orderRepository Repository,
) *Service {
	return &Service{
		orderRepository: orderRepository,
	}
}

// CreateOrder creates an order
func (s *Service) CreateOrder(ctx context.Context, request *CreateOrderParamsRequest) (*Order, error) {
	order := &Order{
		TransactionTime: request.TransactionTime,
	}
	err := s.orderRepository.Insert(ctx, order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// ListOrders lists all orders
func (s *Service) ListOrders(ctx context.Context, params *ListOrdersParams) ([]*Order, int, error) {
	log.Println("list orders executed")

	// get orders in array
	orders := []*Order{}
	orders, err := s.orderRepository.FindAll(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	count, err := s.orderRepository.CountAll(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	return orders, count, nil

}
