package models

import (
	"context"
	"fazztrack/demo/lib"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Partners struct {
	Id    int    `json:"id"`
	Image string `json:"image"`
}

func DetailPartner() []Partners {
	db := lib.DB()
	defer db.Close(context.Background())
	
	rows, _ := db.Query(
	context.Background(),
	`SELECT * FROM "partners" order by "id" asc`,
)

	partner, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Partners])

	if err != nil {
		fmt.Println(err)
	}

	return partner

}