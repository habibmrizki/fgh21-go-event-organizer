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
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListAllUsers(ctx *gin.Context) {

    q := ctx.Query("search")
    limit, _ := strconv.Atoi(ctx.Query("limit"))
    page, _ := strconv.Atoi(ctx.Query("page"))

    if limit == 0 {
        limit = 10
    }

    if page < 1 {
        page = 1
    }
  
	results, count := models.FindAllUsers(q, limit, page)
    totalPage := math.Ceil(float64(count) / float64(limit))
    

    prev := int(totalPage) - 1
    next := int(totalPage) - page
	
	pageInfo := lib.PageInfo{
		TotalData: count,
		TotalPage: int(totalPage),
		Page: page,
		Limit: limit,
		Next: next,
		Prev: prev,
	}
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Show All Categories",
		PageInfo: pageInfo,
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

func UpdateUser(ctx *gin.Context) {
    id, _  := strconv.Atoi(ctx.Param("id"))
    user := models.User{}
    err := ctx.Bind(&user)
    if err != nil {
        fmt.Println(err)
        return
    }

    q := ctx.Query("search")
    limit, _ := strconv.Atoi(ctx.Query("limit"))
    page, _ := strconv.Atoi(ctx.Query("page"))
    dataUser, _ := models.FindAllUsers(q, limit, page)


    result := models.User{}
    for _, v := range dataUser {
        if v.Id == id {
            result = v
        }
    }
	
    if result.Id == 0 {
        ctx.JSON(http.StatusNotFound, lib.Response{
            Success: false,
            Message: "user whit id  not found",
        })
        return
    }
	
    data := models.EditUser(user, id)

    ctx.JSON(http.StatusOK, lib.Response{
        Success: true,
        Message: "user whit id  Edit Success",
        Results: data,
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

// func UpdatePassword(ctx *gin.Context) {
//     id := ctx.GetInt("userId")
//     user := models.FindOneusers(id)

//     if user.Id == 0 {
//         ctx.JSON(http.StatusNotFound, lib.Response{
//             Success: false,
//             Message: "User not found",
//         })
//         return
//     }

//     Password := 
//     var req struct {
//         Password `string form:"password" binding:"required,min=8"`
//     }
//     if err := ctx.ShouldBind(&req); err != nil {
//         ctx.JSON(http.StatusBadRequest, lib.Response{
//             Success: false,
//             Message: "Invalid input data",
//         })
//         return
//     }

//     if err := models.Updatepassword(req.Password, id); err != nil {
//         ctx.JSON(http.StatusInternalServerError, lib.Response{
//             Success: false,
//             Message: "Failed to update password",
//         })
//         return
//     }

//     ctx.JSON(http.StatusOK, lib.Response{
//         Success: true,
//         Message: "Password successfully updated",
//     })
// } 
