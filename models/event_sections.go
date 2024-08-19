package models

import (
	"context"
	"fazztrack/demo/lib"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type EventSection struct {
	Id       int    `json:"id"`
	Name     string `json:"name" form:"name"`
	Quantity int    `json:"quantity" form:"quantity"`
	Price    int    `json:"price" form:"price"`
	EventId  int    `json:"eventId" db:"event_id"`
}

func FindSectionsByEventId(id int) ([]EventSection, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, err := db.Query(
		context.Background(),
		`select * from "event_sections" where "event_id" = $1`, id, // mendapatkan data section
	)

	if err != nil {
		fmt.Println("Error")
	}

	events, err := pgx.CollectRows(rows, pgx.RowToStructByPos[EventSection])

	if err != nil {
		return nil, fmt.Errorf("Error")
	}

	return events, nil
}