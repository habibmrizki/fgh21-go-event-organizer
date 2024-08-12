// package controllers

// import (
// 	"fazztrack/demo/lib"
// 	"strconv"

// 	"fazztrack/demo/models"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func ListAllUsers(ctx *gin.Context){
// 	users := models.FindAllUsers()
// 	ctx.JSON(http.StatusOK, lib.Response{
// 		Success: true,
// 		Message: "Ok",
// 		Results:  users,
// 	})
// }

// func DetailUser(ctx *gin.Context){
// 	id, _ := strconv.Atoi(ctx.Param("id"))
// 	data := models.FindOneUser(id)
// 		if data.Id != 0 {
// 			ctx.JSON(http.StatusOK, lib.Response{
// 				Success: true,
// 				Message: "Invalid ID",
// 				Results: data,
// 			})
// 		} else {
// 			ctx.JSON(http.StatusNotFound, lib.Response{
// 				Success: false,
// 				Message: "User ID not found",
// 			})
// 		}
// }

// func CreateUser(c *gin.Context) {
// 	user := models.User{}
// 	c.Bind(&user)

// 	data := models.CreateUser(user)

// 	c.JSON(http.StatusOK, lib.Response{
// 		Success: true,
// 		Message: "Cearte user success",
// 		Results: data,
// 	})
// }

// func DeleteUser(ctx *gin.Context) {
// 	id, _ := strconv.Atoi(ctx.Param("id"))
// 	data := models.DeleteUser(id)
// 	if data.Id != 0 {
// 		ctx.JSON(http.StatusOK, lib.Response{
// 			Success: true,
// 			Message: "Delete use success",
// 			Results: data,
// 		})
// 	} else {
// 		ctx.JSON(http.StatusNotFound, lib.Response{
// 			Success: false,
// 			Message: "User ID not found",
// 		})
// 	}
// }

// func UpdateUser(ctx *gin.Context) {
// 		id, _ := strconv.Atoi(ctx.Param("id"))
// 		user := models.User{}
// 		ctx.Bind(&user)

// 		data := models.EditUser(user, id)

// 		if data.Id != 0 {
// 			ctx.JSON(http.StatusOK, lib.Response{
// 				Success: true,
// 				Message: "Update data is success",
// 				Results: data,
// 			})
// 		} else {
// 			ctx.JSON(http.StatusNotFound, lib.Response{
// 				Success: false,
// 				Message: "Id is not found",
// 			})
// 		}
// 	}

package controllers

import (
	"fazztrack/demo/lib"
	"fazztrack/demo/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListAllUsers(ctx *gin.Context) {
	results := models.FindAllUsers()
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List all users",
		Results: results,
	})
}

func DetailUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	dataUser := models.FindOneUserById(id)

	if dataUser.Id !=0 {
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "User Found",
			Results: dataUser,
		})
	} else {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "User Not Found",
			Results: dataUser,
		})
	}
}

func CreateUser(ctx *gin.Context) {
   newUser := models.User{}
   
    if err := ctx.ShouldBind(&newUser); 
		err != nil {
        ctx.JSON(http.StatusBadRequest, lib.Response{
            Success: false,
            Message: "Invalid input data",
        })
        return
    }

    data := models.CreateUser(newUser)

    ctx.JSON(http.StatusOK, lib.Response{
        Success: true,
        Message: "User created successfully",
        Results: data,
    })
}

func UpdateUser(c *gin.Context) {
    param := c.Param("id")
    id, _  := strconv.Atoi(param)
    dataUser := models.FindAllUsers()

    user := models.User{}
    err := c.Bind(&user)
    if err != nil {
        fmt.Println(err)
        return
    }

    result := models.User{}
    for _, v := range dataUser {
        if v.Id == id {
            result = v
        }
    }
	
    if result.Id == 0 {
        c.JSON(http.StatusNotFound, lib.Response{
            Success: false,
            Message: "user whit id " + param + " not found",
        })
        return
    }
	
    models.EditUser(user.Email, user.Username, user.Password, param)

    c.JSON(http.StatusOK, lib.Response{
        Success: true,
        Message: "user whit id " + param + " Edit Success",
        Results: user,
    })
}


func DeleteUser(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    dataUser := models.FindOneUserById(id)

    if err != nil {
        ctx.JSON(http.StatusBadRequest, lib.Response{
            Success: false,
            Message: "Invalid user ID",
        })
        return
    }

    err = models.DeleteUser(id)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, lib.Response{
            Success: false,
            Message: "Id Not Found",
        })
        return
    }

    ctx.JSON(http.StatusOK, lib.Response{
        Success: true,
        Message: "User deleted successfully",
        Results: dataUser,
    })
}