package models

import (
	"context"
	"fazztrack/demo/lib"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Profile struct {
	Id       	  int    `json:"id" db:"id"`
	Picture       *string `json:"picture" db:"picture"`
	FullName      string `json:"full_name" form:"fullname" db:"full_name"`
	BirthDate     *string `json:"birth_date" form:"birth_date" db:"birth_date"`
	Gender        int    `json:"gender" form:"gender"`
	PhoneNumber   *string `json:"phone_number" form:"phone_number" db:"phone_number"`
	Profession    *string `json:"profession" form:"profession" `
	NationalityId *int    `json:"nationality_id" form:"nationality_id" db:"nationality_id"`
	UserId        int    `json:"user_id" form:"user_id" db:"user_id"`
}

type JoinUserProfile struct {
	Id          int    `json:"id"`
	FullName    string `json:"fullName"`
	Username    string `json:"username,omitempty"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone-number,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Profession  string `json:"profession,omitempty"`
	Nationality int    `json:"nationality,omitempty"`
	BirthDate   string `json:"birth-date,omitempty"`
}

// Error implements error.
func (p Profile) Error() string {
	panic("unimplemented")
}

func CreateProfile(data Profile) JoinUserProfile {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT INTO "profile" ("picture", "full_name", "birth_date", "gender", "phone_number", "profession", "nationality_id", "user_id") VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := db.Exec(context.Background(), sql, data.Picture, data.FullName, data.BirthDate, data.Gender, data.PhoneNumber, data.Profession, data.NationalityId, data.UserId)
	
	if err != nil {
		fmt.Println(err)
	}

	result := JoinUserProfile{}

	result.Id = data.UserId
	result.FullName = data.FullName

	return result
}

func ListAllProfile ()[]JoinUserProfile {
	db := lib.DB()
	defer db.Close(context.Background())

	joinSql := `select "u"."id", "u"."email", "p"."full_name", "u"."username", "p"."gender", "p"."phone_number","p"."profession", "p"."nationality_id", "p"."birth_date"  
	from "users" "u" 
	join "profile" "p"
	on "u"."id" = "p"."user_id"`
		
	rows, _:= db.Query(
		context.Background(),
		joinSql,
		)
	
	events, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[JoinUserProfile])
	return events
}


func FindProfileByUserId(id int) JoinUserProfile {
	db := lib.DB()
	defer db.Close(context.Background())
	
	var result JoinUserProfile
	for _, v := range ListAllProfile() {
		if v.Id == id {
			result = v
		}
	}
	
	return result
}

// func FindprofileById(id int) Profile {
// 	db := lib.DB() //melakukan koneksi ke database
// 	defer db.Close(context.Background())
// 	rows, _ := db.Query(
// 		context.Background(),
// 		`select "id", "picture", "full_name", "birth_date", "gender", "phone_number", "profession", "nationality_id", "user_id" from "profile" where "id"=$1`,
// 		id,
// 	)
	

// 	profiles, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Profile])

// 	if err != nil {
// 		fmt.Println(err)
// 	}	

// 	profile := Profile{}
// 	for _, v := range profiles {
// 		if v.Id == id {
// 			profile = v
// 		}
// 	}
// 	return profile
// }