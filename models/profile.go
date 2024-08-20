package models

import (
	"context"
	"fazztrack/demo/lib"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Profile struct {
	Id       	  int    `json:"id" db:"id"`
	Picture       string `json:"picture" db:"picture"`
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
	FullName    string `json:"fullName" db:"full_name"`
	Username    string `json:"username,omitempty" db:"username"`
	Email       string `json:"email" db:"email"`
	PhoneNumber *string `json:"phoneNumber,omitempty" db:"phone_number"`
	Gender      *int `json:"gender,omitempty" db:"gender"`
	Profession  *string `json:"profession,omitempty" db:"profession"`
	Nationality *int    `json:"nationality,omitempty" db:"nationality_id"`
	BirthDate   *string `json:"birthDate,omitempty" db:"birth_date"`
}

type Nationalities struct {
	Id int `json:"id"`
	Name string `json:"nationality"`
}

// Error implements error.
func (p Profile) Error() string {
	panic("unimplemented")
}

func Createprofile(data Profile) JoinUserProfile {
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

func ListAllProfile(id int)[]JoinUserProfile {
	db := lib.DB()
	defer db.Close(context.Background()) 

	joinSql := `select "u"."id", "p"."full_name", "u"."username", "u"."email", "p"."phone_number", "p"."gender","p"."profession", "p"."nationality_id", "p"."birth_date"  
	from "users" "u" 
	join "profile" "p"
	on "u"."id" = "p"."user_id"
	`
		
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
	// fmt.Println(id)
	// var result JoinUserProfile
	// for _, v := range ListAllProfile() {
		// 	if v.Id == id {
			// 		result = v
			// 	}
			// }
	
	sql := `select "u"."id", "p"."full_name", "u"."username", "u"."email", "p"."phone_number", "p"."gender", "p"."phone_number","p"."profession", "p"."nationality_id", "p"."birth_date"  
	from "users" "u" 
	join "profile" "p"
	on "u"."id" = "p"."user_id" where "u"."id" = $1`
			

	row := db.QueryRow(
		context.Background(),
		sql, id,
	)
	// fmt.Println(row)
	
	var result JoinUserProfile
	row.Scan(
		&result.Id,
			&result.FullName,
			&result.Username,
			&result.Email,
			&result.PhoneNumber,
			&result.Gender,
			&result.Profession,
			&result.BirthDate,
			&result.Nationality,
		)
		
		return result
	}

func UpdateProfile(data Profile, id int) JoinUserProfile{
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `update "profile" "set" (full_name, birth_date, gender, phone_number, profession, nationality_id) =
			($1, $2,$3,$4,$5,$6) WHERE users_id = $7 `
	db.Exec(context.Background(), sql, data.FullName, data.BirthDate, data.Gender, data.PhoneNumber, data.Profession, data.NationalityId, id)

	result := FindProfileByUserId(id)
	return result
	}

func FindAllNationality() []*Nationalities {
		db := lib.DB()
		defer db.Close(context.Background())
	
		rows, _ := db.Query(
			context.Background(),
			`SELECT * FROM "nationalities" order by "id" asc`,
		)
	
		national, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[Nationalities])
	
		if err != nil {
			fmt.Println(err)
		}
	
		return national
	}

	func FindOneNational(id int) []Nationalities{
		db := lib.DB()
		defer db.Close(context.Background())
	
		rows, _ := db.Query(context.Background(),
			`select * from "nationalities" where "id" = $1`,id,
		)
		national, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Nationalities])
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(national)
		return national
	}