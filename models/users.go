// package models

// type User struct {
// 	Id       int    `json:"id" db:"id"`
// 	Username string `json:"username" form:"username" db:"username"`
// 	Email    string `json:"email" form:"email" db:"email"`
// 	Password string `json:"-" form:"password" db:"password"`
// }

// var dataUser = []User{
// 	{
// 		Id:       1,
// 		Username: "Habib",
// 		Password: "1234",
// 		Email:    "habib@mail.com",
// 	},
// }

// func FindAllUsers() []User {
// 	data := dataUser
// 	return data
// }

// func FindOneUser(id int) User {
// 	data := dataUser

// 	user := User{}
// 	for _, getId := range data {
// 		if id == getId.Id {
// 			user = getId
// 		}
// 	}
// 	return user
// }

// func CreateUser(data User) User {
// 	id := 0
// 	for _, v := range dataUser {
// 		id = v.Id
// 	}

// 	data.Id = id + 1
// 	dataUser = append(dataUser, data)

// 	return data
// }

// func DeleteUser(id int) User {
// 	index := -1
// 	delete := User{}
// 	for id, item := range dataUser {
// 		if item.Id == id {
// 			index = id
// 			delete = item
// 		}
// 	}
// 	if delete.Id != 0 {
// 		dataUser = append(dataUser[:index], dataUser[index+1:]...)
// 	}

// 	return delete
// }

// func EditUser(data User, id int) User {
// 	num := -1
// 	for index, item := range dataUser {
// 		if id == item.Id {
// 			num = index

// 		}
// 	}

// 	if num == 0 {
// 		dataUser[num].Username = data.Username
// 		dataUser[num].Username = data.Username
// 		dataUser[num].Username = data.Username
// 		data.Id = dataUser[num].Id
// 	}
// 	return data
// }

package models

import (
	"context"
	"fazztrack/demo/lib"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type User struct {
	Id       int    `json:"id" db:"id"`
	Email    string `json:"email" form:"email" db:"email"`
	Password string `json:"-" form:"password" db:"password"`
	Username string `json:"username" form:"username" db:"username"`
}

// Error implements error.
func (u User) Error() string {
	panic("unimplemented")
}

func FindAllUsers() []User {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(
		context.Background(),
		`select * from "users" order by "id" asc`,
	)

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])

	if err != nil {
		fmt.Println(err)
	}

	return users
}

func FindOneUserById(id int) User {
	db := lib.DB() //melakukan koneksi ke database
	defer db.Close(context.Background())

	rows, _ := db.Query(
		context.Background(),
		`select "id", "email", "password", "username" from "users" where "id"=$1`,
		id,
	)

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])

	if err != nil {
		fmt.Println(err)
	}	

	user := User{}
	for _, v := range users {
		if v.Id == id {
			user = v
		}
	}
	// fmt.Println(users)
	return user
}

func FindOneUserByEmail(email string) User{
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(
		context.Background(),
		`select *  from "users"`,
	)

	users , err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])

	if err != nil {
		fmt.Println(err)
	}

	user := User{}
	for _, v := range users {
		if v.Email == email {
			user = v
		}
	}

	// fmt.Println(user)
	return user


}

func CreateUser(user User) User {
	db := lib.DB()
	defer db.Close(context.Background())

	user.Password = lib.Encrypt(user.Password)

	sql := `INSERT INTO "users" (email, password, username) VALUES ($1, $2, $3) returning "id", "email", "password", "username"`
	row := db.QueryRow(context.Background(), sql, user.Email, user.Password, user.Username)

	var results User
	row.Scan(
		&results.Id,
		&results.Email,
		&results.Password,
		&results.Username,
	)
	return results
}

func DeleteUser(id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	commandTag, err := db.Exec(
		context.Background(),
		`DELETE FROM "users" WHERE id = $1`,
		id,
	)

	if err != nil {
		return fmt.Errorf("failed to execute delete")
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no user found")
	}

	return nil
}

func EditUser(email string, username string, password string, id string)  {
	db := lib.DB()
	defer db.Close(context.Background())

	dataSql := `update "users" set ("email", "username", "password") = ($1, $2, $3) where id=$4`

	db.Exec(context.Background(), dataSql, email, username, password, id)
}