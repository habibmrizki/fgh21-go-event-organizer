package models

import (
	"context"
	"fazztrack/demo/lib"
)

type FormRegister struct {
	Id              int    `form:"id"`
	FullName        string `json:"full_name" form:"fullname" db:"full_name"`
	Email           string `json:"email" form:"form" db:"email"`
	Password        string `json:"-" form:"password" db:"password"`
	ConfirmPassword string `json:"confirmPassword" binding:"eqfield=Password"`
}

func RegisterUser(user FormRegister) FormRegister {
	db := lib.DB()
	defer db.Close(context.Background())

	user.Password = lib.Encrypt(user.Password)

	sql := `INSERT INTO "users" (email, password) VALUES ($1, $2) returning "id", "email", "password"`
	row := db.QueryRow(context.Background(), sql, user.FullName, user.Email, user.Password)

	var results FormRegister
	row.Scan(
		&results.Id,
		&results.Email,
		&results.Password,
		&results.Password,
	)
	return results
}