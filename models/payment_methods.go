package models

import (
	"context"
	"fazztrack/demo/lib"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type PaymentMethod struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func FindAllPaymentMethod() []PaymentMethod {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(
		context.Background(),
		`select * from "payment_methods" order by "id" asc`,
	)
	paymentMethod, err := pgx.CollectRows(rows, pgx.RowToStructByPos[PaymentMethod])

	if err != nil {
		fmt.Println(err)
	}
	return paymentMethod
}