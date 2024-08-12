package routers

import (
	"fazztrack/demo/controllers"
	// "fazztrack/demo/middlewares"

	"github.com/gin-gonic/gin"
)

func EventRouter(routerGroup *gin.RouterGroup) {
	// routerGroup.Use(middlewares.AuthMiddleware())
	routerGroup.GET("", controllers.ListAllEvent)
	routerGroup.GET("/:id", controllers.DetailEvent)
	routerGroup.POST("", controllers.CreateEvent)
	routerGroup.PATCH("/:id", controllers.UpdateEvent)
	routerGroup.DELETE("/:id", controllers.DeleteEvent)
}
