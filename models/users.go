package models

type User struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" form:"username" db:"username"`
	Email    string `json:"email" form:"email" db:"email"`
	Password string `json:"-" form:"password" db:"password"`
}

var dataUser = []User{
	{
		Id:       1,
		Username: "Habib",
		Password: "1234",
		Email:    "habib@mail.com",
	},
}

func FindAllUsers() []User {
	data := dataUser
	return data
}

func FindOneUser(id int) User {
	data := dataUser

	user := User{}
	for _, getId := range data {
		if id == getId.Id {
			user = getId
		}
	}
	return user
}

func CreateUser(data User) User {
	id := 0
	for _, v := range dataUser {
		id = v.Id
	}

	data.Id = id + 1
	dataUser = append(dataUser, data)

	return data
}

func DeleteUser(id int) User {
	index := -1
	delete := User{}
	for id, item := range dataUser {
		if item.Id == id {
			index = id
			delete = item
		}
	}
	if delete.Id != 0 {
		dataUser = append(dataUser[:index], dataUser[index+1:]...)
	}

	return delete
}

func EditUser(data User, id int) User {
	num := -1
	for index, item := range dataUser {
		if id == item.Id {
			num = index

		}
	}

	if num == 0 {
		dataUser[num].Username = data.Username
		dataUser[num].Username = data.Username
		dataUser[num].Username = data.Username
		data.Id = dataUser[num].Id
	}
	return data
}