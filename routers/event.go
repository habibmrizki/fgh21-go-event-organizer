package routers

import (
	"fazztrack/demo/controllers"
	"fazztrack/demo/middlewares"

	"github.com/gin-gonic/gin"
)

func EventRouter(routerGroup *gin.RouterGroup) {
	// routerGroup.Use(middlewares.AuthMiddleware())
	routerGroup.GET("", controllers.ListAllEvent)
	routerGroup.GET("/section/:id", controllers.ShowSectionsByEventId)
	routerGroup.GET("/payment_method", middlewares.AuthMiddleware(), controllers.ListAllPaymentMethod)
	routerGroup.GET("/:id", controllers.DetailEvent)
	routerGroup.POST("", middlewares.AuthMiddleware(), controllers.CreateEvent)
	routerGroup.PATCH("/:id", middlewares.AuthMiddleware(), controllers.UpdateEvent)
	routerGroup.DELETE("/:id", middlewares.AuthMiddleware(), controllers.DeleteEvent)
}
