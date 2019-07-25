package posgres

import (
	"context"
	"log"

	"github.com/kluwena/go-api-practice/internal/data"
	"github.com/kluwena/go-api-practice/internal/order"
)

// OrderRepository is the implementation for order repository in postgres.
type OrderRepository struct {
	tableName string
	queryer   data.Queryer
}

// NewOrderRepository creates a new order repository
func NewOrderRepository(
	queryer data.Queryer,
) *OrderRepository {
	tableName := "order"
	return &OrderRepository{
		queryer:   queryer,
		tableName: tableName,
	}
}

// Insert push a new order into the database
func (r *OrderRepository) Insert(ctx context.Context, order *order.Order) error {

	db := data.TransactionFromContext(ctx, r.queryer)

	statement, err := db.PrepareNamed(`INSERT INTO "order"("transaction_time") VALUES (:transaction_time)`)
	if err != nil {
		return err
	}

	res, err := statement.Exec(map[string]interface{}{
		"transaction_time": order.TransactionTime,
	})

	if err != nil {
		return err
	}

	log.Println(res)

	return nil
}

// CountAll retrieves the order count
func (r *OrderRepository) CountAll(ctx context.Context, order *order.ListOrdersParams) (int, error) {
	return 0, nil
}

// FindAll retrieves the order history
func (r *OrderRepository) FindAll(ctx context.Context, params *order.ListOrdersParams) ([]*order.Order, error) {

	query := `select * from "order" limit 10 offset 0`
	db := data.TransactionFromContext(ctx, r.queryer)

	statement, err := db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}

	res := []*order.Order{}

	if err := statement.Select(&res, map[string]interface{}{}); err != nil {
		return nil, err
	}

	return res, nil
}
