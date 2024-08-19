package controllers

import (
	"fazztrack/demo/lib"
	"fazztrack/demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListAllPaymentMethod(ctx *gin.Context) {
	results := models.FindAllPaymentMethod()
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "List all payment method",
		Results: results,
	})
}