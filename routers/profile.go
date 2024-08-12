package routers

import (
	"fazztrack/demo/controllers"
	"fazztrack/demo/middlewares"

	"github.com/gin-gonic/gin"
)

func ProfileRouter(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/", controllers.CreateProfile)
	routerGroup.Use(middlewares.AuthMiddleware())
}
