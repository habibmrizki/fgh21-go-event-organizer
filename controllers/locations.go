package controllers

import (
	"fazztrack/demo/lib"
	"fazztrack/demo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowAllLocation(ctx *gin.Context) {
    search := ctx.Query("search")
    limit, _  := strconv.Atoi(ctx.Query("limit"))
    page, _  := strconv.Atoi(ctx.Query("page"))

    if limit < 1 {
        limit = 7
    }

    if page < 1 {
        page = 1
    }

    result := models.FindAllLocations(search, limit, page)

    ctx.JSON(http.StatusOK, lib.Response{
        Success: true,
        Message: "List All Locations",
        Results: result,
    })
}