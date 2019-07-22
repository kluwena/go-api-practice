package order

import (
	"time"
	"context"
)

// Order represent the order entity
type Order struct {
	ID              int
	transactionTime time.Time
}

type (
	// CreateOrderRequest request for Order Request
	CreateOrderParamsRequest struct {
		Title           string    `json: "title"`
		TransactionTime time.Time `json: "transactionTime"`
	}
)

type (
	ServiceInterface interface {
		CreateOrder(ctx context.Context, *CreateOrderRequest) (*Order, error)
	}
	
	// Repository represents the interface for order entity
	Repository interface {
		Insert(ctx context.Context order *Order) error
	}

	// Service represents the implementation details of order service interface
	Service struct {
		orderRepository Repository
	}

)



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
		transactionTime: request.TransactionTime,
	}
	err := s.orderRepository.Insert(ctx, order)
	if err != nil {
		return nil, err
	}

	return order, nil
}