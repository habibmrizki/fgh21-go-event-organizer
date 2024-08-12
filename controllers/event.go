package controllers

import (
	"fazztrack/demo/lib"
	"fazztrack/demo/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListAllEvent(ctx *gin.Context) {
	results := models.FindAllEvent()
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List all users",
		Results: results,
	})
}

func DetailEvent(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	dataUser := models.FindEventById(id)

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

func CreateEvent(ctx *gin.Context) {
	newEvent := models.Event{}
	 if err := ctx.ShouldBind(&newEvent); 
		 err != nil {
		 ctx.JSON(http.StatusBadRequest, lib.Response{
			 Success: false,
			 Message: "Invalid input data",
		 })
		 return
	 }
 
	 data := models.CreateEvent(newEvent)
	 ctx.JSON(http.StatusOK, lib.Response{
		 Success: true,
		 Message: "User created successfully",
		 Results: data,
	 })
 }

 func DeleteEvent(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    dataEvent := models.FindEventById(id)

    if err != nil {
        ctx.JSON(http.StatusBadRequest, lib.Response{
            Success: false,
            Message: "Invalid user ID",
        })
        return
    }

    err = models.DeleteEvent(id)
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
        Results: dataEvent,
    })
}


func UpdateEvent(c *gin.Context) {
    param := c.Param("id")
    id, _  := strconv.Atoi(param)
    dataEvent := models.FindAllEvent()


    event := models.Event{}
    err := c.Bind(&event)
    if err != nil {
        fmt.Println(err)
        return
    }

    result := models.Event{}
    for _, v := range dataEvent {
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
	
    models.EditUser(event.Image, event.Image, event.Title, param)

    c.JSON(http.StatusOK, lib.Response{
        Success: true,
        Message: "user whit id " + param + " Edit Success",
        Results: event,
    })
}

// func UpdateEvent(ctx *gin.Context) {
//     param := ctx.Param("id")
//     id, _  := strconv.Atoi(param)
//     data := models.FindAllEvent()

//     event := models.Event{}
//     err := ctx.Bind(&event)
//     if err != nil {
//         fmt.Println(err)
//         return
//     }

//     result := models.Event{}
//     for _, v := range data {
//         if v.Id == id {
//             result = v
//         }
//     }

//     if result.Id == 0 {
//         ctx.JSON(http.StatusNotFound, lib.Response{
//             Success: false,
//             Message: "Events with id " + param + " not found",
//         })
//         return
//     }
//     models.EditUser(event.Image, event.Title,event.Date, event.Description, event.Created_by, param)

//     ctx.JSON(http.StatusOK, lib.Response{
//         Success: true,
//         Message: "Events with id " + param + " Edit Success",
//         Results: event,
//     })
// }