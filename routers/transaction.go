package routers

import (

	// "fazztrack/demo/middlewares"

	"fazztrack/demo/controllers"
	"fazztrack/demo/middlewares"

	"github.com/gin-gonic/gin"
)

func TransactionRouter(routerGroup *gin.RouterGroup) {	
 	routerGroup.Use(middlewares.AuthMiddleware())
	routerGroup.POST("", controllers.CreateTransaction)
	routerGroup.GET("/:id", controllers.FindAllTransactionByUserId)
}