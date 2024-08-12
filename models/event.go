package models

import (
	"context"
	"fazztrack/demo/lib"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Event struct {
	Id          int    `jsob:"id" form:"id" db:"id"`
	Image       string `json:"image" form:"image" db:"image"`
	Title       string `json:"title" form:"title" db:"title"`
	Date        string `json:"date" form:"date" db:"date"`
	Description string `json:"description" form:"description" db:"description"`
	LocationId  *int    `json:"locationd_id" form:"location_id" db:"location_id"`
	CreatedBy   *int    `json:"created_by" form:"created_by" db:"created_by"`
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
	fmt.Println(data)
	
	sql := `insert into "events" (image, title, date, description) values ($1, $2, $3, $4) returning "id", "image", "title", "date", "description"`
	row := db.QueryRow(context.Background(), sql, data.Image, data.Title, data.Date, data.Description,)
	
	var results Event
	fmt.Println(data.Description)
	row.Scan(
		&results.Id,
		&results.Image,
		&results.Title,
		&results.Date,
		&results.Description,
	)
	return results
	}


func EditEvent(Image string, Title string, Date string, Description string) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `UPDATE events SET image = $1, title = $2, date = $3, description = $4 WHERE id = $5`
	db.Exec(context.Background(), sql, Image, Title, Date, Description, )
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