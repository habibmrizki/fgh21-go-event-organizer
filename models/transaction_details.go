package models

import (
	"context"
	"fazztrack/demo/lib"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type TransactionDetail struct {
	Id            int `json:"id"`
	SectionId     int `json:"sectionId" form:"sectionId" db:"section_id" `
	TransactionId int `json:"transactionId" form:"transactionId" db:"transaction_id" `
	TickectQty    int `json:"ticketqty" form:"ticketqty" db:"ticket_qty" `
}

func CreateTransactionDetail(data TransactionDetail) TransactionDetail{
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT INTO "transaction_details" (section_id, transaction_id, ticket_qty) VALUES ($1, $2, $3) returning "id", "section_id", "transaction_id", "ticket_qty"`
	rows, err := db.Query(context.Background(), sql, data.SectionId, data.TransactionId, data.TickectQty)
	if err!=nil{
		fmt.Println(err)
	}

	results, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[TransactionDetail])
	if err!=nil{
		fmt.Println(err)
	}

	// var results TransactionDetail
	// row.Scan(
	// 	&results.Id,
	// 	&results.SectionId,
	// 	&results.TransactionId,
	// 	&results.TickectQty,
	// )
	return results
}