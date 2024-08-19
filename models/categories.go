package models

import (
	"context"
	"fazztrack/demo/lib"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Categories struct {
	Id int `json:"id" db:"id"`
	Name string `json:"name" form:"name"`
}

func CountCategories(search string) (int) {
	db := lib.DB()
	
	defer db.Close(context.Background())

	sql := `select count(id) as "Total" from "categories" where "name"=$1`
	row := db.QueryRow(context.Background(),sql, search)

	var total int

	row.Scan(
		&total,
	)		
	
	return total
}

func FindAllCategories(search string, limit int, page int) ([]Categories, int) {
	db := lib.DB()
	offset := (page - 1) * limit

	defer db.Close(context.Background())
	sql := 	`
		select * from "categories" 
		WHERE "name" ilike '%' || $1 || '%'
		limit $2 offset $3
	`
	rows, _ := db.Query(
		context.Background(),
		sql, search, limit, offset,
	)
	category, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Categories])
	if err != nil {
		fmt.Println(err)
	}

	result := CountCategories(search)

	return category, result
}

func FindCategoriesById(id int) Categories {
	db := lib.DB() //melakukan koneksi ke database
	defer db.Close(context.Background())
	
	rows, _ := db.Query(
		context.Background(),
		`select "id", "name" from "categories" where "id"=$1`, 
		id,
	)
	
	categories, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Categories])
	
	if err != nil {
		fmt.Println(err)
		}	
		
		category := Categories{}
		for _, v := range categories {
			if v.Id == id {
				category = v
			}
		}
		fmt.Println(categories)
		return category
}

func CreateCategories(data Categories) Categories {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT INTO "categories" ("name") VALUES ($1) returning "id", "name"`
	row := db.QueryRow(context.Background(), sql, data.Name)

	var Results Categories
	row.Scan(
		&Results.Id,
		&Results.Name,
	)
	return Results
}

func UpdateCategories(name string, id int) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `UPDATE "categories" SET ("name") = ($1) where id = $2`
	db.Exec(context.Background(), sql, name, id )
}

func DeleteCategories(id int) error {
	db := lib.DB()
	defer db.Close(context.Background())
		
	commandTag, err := db.Exec(
		context.Background(),`DELETE FROM "categories" WHERE id = $1`,id,
	)
			
	if err != nil {
		return fmt.Errorf("failed to execute delete")
	}

	if commandTag.RowsAffected() == 0 {
	return fmt.Errorf("no user found")
	}

	return nil
}