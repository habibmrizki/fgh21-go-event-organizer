package routers

import (
	"fazztrack/demo/controllers"

	"github.com/gin-gonic/gin"
)

func LocationRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("", controllers.ShowAllLocation)
}