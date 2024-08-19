package controllers

import (
	"fazztrack/demo/lib"
	"fazztrack/demo/models"
	"fmt"
	"math"

	// "math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindAllCategories(ctx *gin.Context) {
	q := ctx.Query("search")
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ :=strconv.Atoi(ctx.Query("limit"))

	if limit < 1 {
        limit = 5
    }

    if page < 1 {
        page = 1
    }


	results, count:= models.FindAllCategories(q, page, limit)

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
		Results: results,
		PageInfo: pageInfo,
	})
}

func DetailCategories(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	dataCategories := models.FindCategoriesById(id)

	if dataCategories.Id != 0 {
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "category is found",
			Results: dataCategories,
		})
	} else {
		ctx.JSON(http.StatusNotFound, lib.Response{
			Success: false,
			Message: "category not found",
			Results: dataCategories,
		})
	}
}

func CreateCategories(ctx *gin.Context) {
	newCategories := models.Categories{}
	if err := ctx.ShouldBind(&newCategories);
	err != nil {
        ctx.JSON(http.StatusBadRequest, lib.Response{
            Success: false,
            Message: "Invalid input data",
        })
        return
    }

    data := models.CreateCategories(newCategories)

    ctx.JSON(http.StatusOK, lib.Response{
        Success: true,
        Message: "User created successfully",
        Results: data,
    })
}

func UpdateCategories(ctx *gin.Context) {
    param := ctx.Param("id")
    id, _  := strconv.Atoi(param)

    q := ctx.Query("search")
    limit, _ := strconv.Atoi(ctx.Query("limit"))
    page, _ := strconv.Atoi(ctx.Query("page"))
    dataUser, _ := models.FindAllCategories(q, limit, page)

    category := models.Categories{}
    err := ctx.Bind(&category)
    if err != nil {
        fmt.Println(err)
        return
    }

    result := models.Categories{}
    for _, v := range dataUser {
        if v.Id == id {
            result = v
        }
    }
	
    if result.Id == 0 {
        ctx.JSON(http.StatusNotFound, lib.Response{
            Success: false,
            Message: "user whit id " + param + " not found",
        })
        return
    }
	
    models.UpdateCategories(category.Name, id)

    ctx.JSON(http.StatusOK, lib.Response{
        Success: true,
        Message: "user whit id " + param + " Edit Success",
        Results: category,
    })
}

func DeleteCategories(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    dataUser := models.FindCategoriesById(id)

    if err != nil {
        ctx.JSON(http.StatusBadRequest, lib.Response{
            Success: false,
            Message: "Invalid user ID",
        })
        return
    }

    err = models.DeleteCategories(id)
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