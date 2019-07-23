package order

import (
	"time"
	"context"
)

// Order represent the order entity
type Order struct {
	ID              int
	TransactionTime time.Time
}

type (
	// CreateOrderParamsRequest request for Order Request
	CreateOrderParamsRequest struct {
		Title           string    `json: "title"`
		TransactionTime time.Time `json: "transactionTime"`
	}
)

type ServiceInterface interface {
	CreateOrder(ctx context.Context, params *CreateOrderParamsRequest) (*Order, error)
}

// Repository represents the interface for order entity
type	Repository interface {
	Insert(ctx context.Context , order *Order) (error)
}

// Service represents the implementation details of order service interface
type	Service struct {
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