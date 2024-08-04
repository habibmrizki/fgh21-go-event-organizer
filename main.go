package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Result  interface{} `json:"result"` //Menampung banyak data untuk inputanya
}


type Users struct {
	Id       int `json:"id"`
	Fullname string `json:"fullname" form:"fullname" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"-" form:"password" binding:"required,min=8"`
}

func main() {
	
	data := []Users{
		{
			Id:       1,
			Fullname: "Habib",
			Password: "1234",
			Email:    "habib@mail.com",
		},
		{
			Id:       2,
			Fullname:     "Joko",
			Password: "1234",
			Email:    "joko@mail.com",
		},
		{
			Id:       3,
			Fullname: "Budi",
			Password: "1234",
			Email:    "budi@mail.com",
		},
	}
	r := gin.Default()
	r.Use(corsMiddleware())

	r.POST("/users", func(c *gin.Context) { //Bikin kedua
		user := Users{}
		err := c.Bind(&user)

		result := 0
		for _, v := range data {
			result = v.Id
		}	
		user.Id = result + 1

		isExist := true
		for _, item := range data {
			if item.Email == user.Email {
				isExist = false
			}
		}
		if err == nil {
			if isExist {
				data = append(data, user)
				c.JSON(http.StatusOK, Response{
				Success: true,
				Message: "Create user success",
				Result: data,
				})
			} else {
				c.JSON(http.StatusUnauthorized, Response{
					Success: false,
					Message: "Email already exist",
					
				})
			}
		} else {
			c.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "data was prohibited",	
			})
		}
	})

	r.GET("/users", func(c *gin.Context) { //Bikin pertama. bagian c bebas untuk penamaan
		c.JSON(http.StatusOK, Response{
			Success: true,
			Message: "Ok",
			Result:  data,
		})
	})

	// GET USERS
	r.GET("/users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Invalid ID",
			})
			return
		}

		for _, user := range data {
			if user.Id == id {
				c.JSON(http.StatusOK, Response{
					Success: true,
					Message: "User found",
					Result:  []Users{user},
				})
				return
			}
		}
		c.JSON(http.StatusNotFound, Response{
			Success: false,
			Message: "User not found",
		})
	})



	// PACTH USERS
	r.PATCH("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		selected := -1 //ag

		for index, item := range data{ //mendapatkan sesuai id user
			if item.Id == id {
				selected = index
			}
		}	
		
		if selected != -1 {
			form := Users{}
			c.Bind(&form)
			data[selected].Fullname = form.Fullname
			data[selected].Email = form.Email
			data[selected].Password = form.Password
			c.JSON(http.StatusOK, Response{
				Success: true,
				Message: "Update susccess",
				Result: data[selected],
			}) 
		} else {
			c.JSON(http.StatusNotFound, Response{
				Success: false,
				Message: "User not found",
			}) 
		}
	})


	// POST AUTH LOGIN USERS
	r.POST("/auth/login", func(c *gin.Context) {
		var user Users
		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Invalid request",
			})
			return
		}

		email := user.Email
		password := user.Password

		for _, u := range data {
			if u.Email == email && u.Password == password {
				c.JSON(http.StatusOK, Response{
					Success: true,
					Message: "Login success",
				})
				return
			}
		}

		c.JSON(http.StatusUnauthorized, Response{
			Success: false,
			Message: "Invalid email or password",
		})
	})

	
	// DELETE USERS
	r.DELETE("/users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success:  false,
			Message: "ID Not Found",
		})
		return
	}

	for i, user := range data {
		if user.Id == id {
			data = append(data[:i], data[i+1:]...)
			c.JSON(http.StatusOK, Response{
				Success:  true,
				Message: "User deleted successfully",
			})
			return
		}
	}
	
	c.JSON(http.StatusNotFound, Response{
		Success:  false,
		Message: "User not found",
	})
})
r.Run("localhost:8080")
}
func corsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}


// r.DELETE("users/:id", func(c *gin.Context) {
	// 	id, _ := strconv.Atoi(c.Param("id"))
	
// 	selected := -1 //ag

// 	for index, item := range data{ //mendapatkan sesuai id user
// 		if item.Id == id {
	// 			selected = index
	// 		}
	// 	}	
	

	// 	if selected != -1 {
		// 		form := Users{}
		// 		c.Bind(&form)
		// 		data[selected].Fullname = form.Fullname
		// 		data[selected].Email = form.Email
		// 		data[selected].Password = form.Password
		// 		c.JSON(http.StatusOK, Response{
			// 			Success: true,
			// 			Message: "Update susccess",
			// 			Result: data[selected],
			// 		}) 
			// 	} else {
				// 		c.JSON(http.StatusOK, Response{
					// 			Success: false,
					// 			Message: "User not found",
					// 		}) 
					// 	}
					// })
					
					
					
					
				// 	 r.DELETE("/users/:id", func(c *gin.Context){
				//      index, _ := strconv.Atoi(c.Param("id"))
				//      for i := 0; i < len(data); i++ {
				//          if index == data[i].Id {
				//              data = append(data[:i], data[i+1:]...)
				//              c.JSON(http.StatusOK, Response{
				//                  Success: true,
				//                  Message: "login success",
				//              })
				//          } else {
				//              c.JSON(http.StatusUnauthorized, Response{
				//                  Success: true,
				//                  Message: "login success",
				//                  Result: data,
				//          })
				//          }
				//      }
				//  })