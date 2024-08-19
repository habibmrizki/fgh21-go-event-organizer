package controllers

import (
	"fazztrack/demo/lib"
	"fazztrack/demo/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowSectionsByEventId(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
    dataEvent, err := models.FindSectionsByEventId(id)
    fmt.Println(dataEvent)

    if err != nil {
        ctx.JSON(http.StatusNotFound, lib.Response{
            Success: false,
            Message: "Failed all section by event",
        })
        return
    }

    ctx.JSON(http.StatusOK, lib.Response{
        Success: true,
        Message: "List all section by event",
        Results: dataEvent,
    })
}