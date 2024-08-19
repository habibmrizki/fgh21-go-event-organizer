package models

import (
	"context"
	"fazztrack/demo/lib"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Locations struct {
    Id    int    `json:"id" db:"id"`
    Name  string `json:"name" form:"name" db:"name"`
    Lat   string `json:"lat" form:"lat" db:"lat"`
    Long  string `json:"long" form:"long" db:"long"`
    Image string `json:"image" form:"image" db:"image"`
}

func FindAllLocations(search string, limit int, page int) []Locations {
    db := lib.DB()

    defer db.Close(context.Background())

    offset := (page - 1) * limit

    rows, _ := db.Query(
        context.Background(),
        `select * from "locations" where "name" ilike '%' || $1 || '%' limit $2 offset $3`, search, limit, offset,
    )

    locations, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Locations])

    if err != nil {
        fmt.Println(err)
    }
    return locations
}