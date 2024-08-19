package models

import (
	"context"
	"fazztrack/demo/lib"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Event struct {
	Id          int    `json:"id" form:"id" db:"id"`
	Image       string `json:"image" form:"image" db:"image"`
	Title       string `json:"title" form:"title" db:"title"`
	Date        string `json:"date" form:"date" db:"date"`
	Description string `json:"description" form:"description" db:"description"`
	LocationId  *int    `json:"locationd_id"  db:"location_id"`
	CreatedBy   *int    `json:"created_by"  db:"created_by"`
}

func FindAllEvent() []Event {
	db := lib.DB()
	defer db.Close(context.Background())
	
	rows, _ := db.Query(
		context.Background(),
		`select * from "events" order by "id" asc`,
	)
	
	events, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Event])
	
	if err != nil {
		fmt.Println(err)
	}
	
	return events
}


func FindEventById(id int) Event {
	db := lib.DB() //melakukan koneksi ke database
	defer db.Close(context.Background())
	
	rows, _ := db.Query(
		context.Background(),
		`select * from "events" where "id"=$1`, 
		id,
	)
	
	events, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Event])
	
	if err != nil {
		fmt.Println(err)
		}	
		
		Event := Event{}
		for _, v := range events {
			if v.Id == id {
				Event = v
			}
		}
		fmt.Println(events)
		return Event
	}
	
	
func CreateEvent(data Event) Event{
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT INTO "events" ("image", "title", "date", "description")
	VALUES ($1, $2, $3, $4)`
	db.Exec(context.Background(), sql, data.Image, data.Title, data.Date, data.Description)
	id := 0
	for _, v := range FindAllEvent() {
		id = v.Id
	}
	data.Id = id

	return data
	}


func EditEvent(Image string, Title string, Date string, Description string, id int) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `UPDATE "events" SET ("image", "title", "date", "description") = ($1, $2, $3, $4) where id = $5`
	db.Exec(context.Background(), sql, Image, Title, Date, Description, id )
}

// func Updateevents(image string, tittle string, date int, description string, location int, created_by int, id string) {

//     db := lib.DB()
//     defer db.Close(context.Background())

//     dataSql := `update "events" set (image, tittle, date, description, location, created_by) = ($1, $2, $3, $4, $5, $6) where id=$7`

//     db.Exec(context.Background(), dataSql, image, tittle, date, description, location, created_by, id)
// }

func DeleteEvent(id int) error {
	db := lib.DB()
	defer db.Close(context.Background())
		
	commandTag, err := db.Exec(
		context.Background(),`DELETE FROM "events" WHERE id = $1`,id,
	)
			
	if err != nil {
		return fmt.Errorf("failed to execute delete")
	}

	if commandTag.RowsAffected() == 0 {
	return fmt.Errorf("no user found")
	}

	return nil
}