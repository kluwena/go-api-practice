package order

import (
	"context"
	"time"
)

// Order represent the order entity
type Order struct {
	ID              int       `db:"id"`
	TransactionTime time.Time `db:"transaction_time"`
}

// CreateOrderParams request for Order Request
type CreateOrderParams struct {
	TransactionTime time.Time `json:"transactionTime"`
}

// ListOrdersParams is params for list orders
type ListOrdersParams struct {
	TransactionTime string
}

// ServiceInterface specifies the interface of order service.
type ServiceInterface interface {
	CreateOrder(ctx context.Context, params *CreateOrderParams) (*Order, error)
	ListOrders(ctx context.Context, params *ListOrdersParams) ([]*Order, int, error)
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
func (s *Service) CreateOrder(ctx context.Context, params *CreateOrderParams) (*Order, error) {
	order := &Order{
		TransactionTime: params.TransactionTime,
	}
	err := s.orderRepository.Insert(ctx, order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// ListOrders lists all orders
func (s *Service) ListOrders(ctx context.Context, params *ListOrdersParams) ([]*Order, int, error) {
	// get orders in array
	// orders := []*Order{}
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
