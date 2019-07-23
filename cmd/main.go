package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	internalhttp "github.com/kluwena/go-api-practice/internal/http"
	"github.com/kluwena/go-api-practice/internal/order"
	orderpg "github.com/kluwena/go-api-practice/internal/order/postgres"
	_ "github.com/lib/pq"
)

func main() {
	flag.Parse()

	db, err := sqlx.Open("postgres", "postgres://postgres@localhost:5433/postgres?sslmode=disable")
	if err != nil {
		log.Fatalln("error connection db:", err)
	}
	orderService := order.NewService(orderpg.NewOrderRepository(db))
	s := &http.Server{
		Addr: ":8080",
		Handler: internalhttp.NewServer(
			orderService,
		),
	}
	s.ListenAndServe()

}
