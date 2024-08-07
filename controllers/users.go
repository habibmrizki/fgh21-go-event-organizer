package controllers

import (
	"fazztrack/demo/lib"
	"strconv"

	"fazztrack/demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func ListAllUsers(ctx *gin.Context){
	users := models.FindAllUsers()
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Ok",
		Results:  users,
	})
}



func DetailUser(ctx *gin.Context){
	id, _ := strconv.Atoi(ctx.Param("id"))
	data := models.FindOneUser(id)
		if data.Id != 0 {
			ctx.JSON(http.StatusOK, lib.Response{
				Success: true,
				Message: "Invalid ID",
				Results: data,
			})
		} else {
			ctx.JSON(http.StatusNotFound, lib.Response{
				Success: false,
				Message: "User ID not found",
			})
		}
}


func CreateUser(c *gin.Context) {
	user := models.User{}
	c.Bind(&user)

	data := models.CreateUser(user)

	c.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Cearte user success",
		Results: data,
	})
}


func DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data := models.DeleteUser(id)
	if data.Id != 0 {
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "Delete use success",
			Results: data,
		})
	} else {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "User ID not found",
		})
	}
}

func UpdateUser(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		user := models.User{}
		ctx.Bind(&user)

		data := models.EditUser(user, id)


		if data.Id != 0 {
			ctx.JSON(http.StatusOK, lib.Response{
				Success: true,
				Message: "Update data is success",
				Results: data,
			})
		} else {
			ctx.JSON(http.StatusNotFound, lib.Response{
				Success: false,
				Message: "Id is not found",
			})
		}
	}

