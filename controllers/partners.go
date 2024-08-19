package controllers

import (
	"fazztrack/demo/lib"
	"fazztrack/demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListAllPartner(ctx *gin.Context) {
	results := models.DetailPartner()
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Detail partners",
		Results: results,
	})
}