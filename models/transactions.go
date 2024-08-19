package models

import (
	"context"
	"fazztrack/demo/lib"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Transaction struct {
	Id int	`json:"id"`
	EventId int `json:"eventId"  db:"event_id" `
	PaymentMethodId int	`json:"paymentMethodId" db:"payment_method_id"`
	UserId int `json:"userId" db:"user_id"`
}

type DetailTransactionId struct {
	TransactionId 	int `json:"transaction_id"`
	FullName 		string `json:"fullName" db:"full_name"`
	EventTitle 		string `json:"eventTitle"`
	Location 		*int `json:"location"`
	Date 			string `json:"date"`
	PaymentMethod 	string `json:"paymentMethod" db:"payment_method"`
	SectionName 	[]string `json:"sectionName" db:"section_name"`
	TickectQty 		[]int `json:"ticketQty" db:"ticket_qty"`
}

func CreateTransaction(data Transaction) Transaction {

	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT INTO "transactions" ("event_id", "payment_method_id", "user_id") VALUES ($1, $2, $3) returning "id", "event_id", "payment_method_id", "user_id"`
	rows,err := db.Query(context.Background(), sql, data.EventId, data.PaymentMethodId, data.UserId)

	if err!=nil{
		fmt.Println(err)
	}

	results, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[Transaction])
	if err!=nil{
		fmt.Println(err)
	}

	// var results Transaction
	// row.Scan(
	// 	&results.Id,
	// 	&results.EventId,
	// 	&results.PaymentMethodId,
	// 	&results.UserId,
	// )
	// fmt.Println(results)
	return results
}

// func CreateTransactionDetail(data Transaction) Transaction {
// 	db := lib.DB()
// 	defer db.Close(context.Background())

// 	sql := `INSERT INTO "transactions" (eventId, paymentMethodId, userId) VALUES ($1, $2, $3) returning "id", "eventId", "paymentMethodId", "userId"`
// 	row := db.QueryRow(context.Background(), sql, data.EventId, data.PaymentMethodId, data.UserId)

// 	var results Transaction
// 	row.Scan(
// 		&results.Id,
// 		&results.EventId,
// 		&results.PaymentMethodId,
// 		&results.UserId,
// 	)
// 	return results
// }

func DetailTransaction(id int) DetailTransactionId {
	db := lib.DB()
	defer db.Close(context.Background())
	sql := `select "t"."id", "p"."full_name", "e"."title" as "event_title",
			"e"."location_id", "e"."date", "pm"."name" as "payment_method",
			array_agg("es"."name") as "section_name", array_agg("td"."ticket_qty") as "ticket_qty"
			from "transactions" "t"
			join "users" "u" on "u"."id" = "t"."user_id"
			join "profile" "p" on "p"."user_id" = "u"."id"
			join "events" "e" on "e"."id" = "t"."event_id"
			join "payment_methods" "pm" on "pm"."id" = "t"."payment_method_id"
			join "transaction_details" "td" on "td"."transaction_id" = "t"."id"
			join "event_sections" "es" on "es"."id" = "td"."section_id"
			WHERE "t"."id" = $1
			group by "t"."id","p"."full_name", "e"."title", "e"."location_id", "e"."date", "pm"."name" `
	rows,err := db.Query(context.Background(),sql, id)
	if err!=nil{
		fmt.Println(err)
	}

	results, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[DetailTransactionId])
	if err!=nil{
		fmt.Println(err)
	}

	// var results DetailTransactionId
	// row.Scan(
	// 	&results.TransactionId,
	// 	&results.FullName,
	// 	&results.EventTitle,
	// 	&results.Location,
	// 	&results.Date,
	// 	&results.PaymentMethod,
	// 	&results.SectionName,
	// 	&results.TickectQty,
	// )
	return results

}

func FindAllTransactionOnByUserId(id int) ([]Transaction, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, err := db.Query(
		context.Background(),
		`select * from "transactions" where "user" = $1`, id,
	)

	if err != nil {
		fmt.Println("Error")
	}

	transactions, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Transaction])

	if err != nil {
		return nil, fmt.Errorf("Error")
	}

	return transactions, nil
}