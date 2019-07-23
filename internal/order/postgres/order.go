package posgres

import (
	"context"
	"log"

	"github.com/kluwena/go-api-practice/internal/data"
	"github.com/kluwena/go-api-practice/internal/order"
)



type OrderRepository struct {
	queryer data.Queryer
}


func (r *OrderRepository) Insert (ctx context.Context, order *order.Order) error {
	
	statement, err := db.PreparedNamed(`INSERT INTO "order"("transaction_time") VALUES (:transaction_time)`)
	if err != nil {
		return err
	}
	
	res, err := statement.Exec(map[string]interface{}){
		"transaction_time": order
	}
	
	if err != nil {
		return err
	}

	log.Println(res)

	return nil
}

func newOrderRepository(
	queryer data.Queryer,
) *OrderRepository {
	return &OrderRepository{
		queryer: queryer,
	}
}