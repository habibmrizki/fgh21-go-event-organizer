package routers

import (
	"fazztrack/demo/controllers"

	"github.com/gin-gonic/gin"
)

func PartnerRouter(routerGroup *gin.RouterGroup) {
	// routerGroup.Use(middlewares.AuthMiddleware())
	routerGroup.GET("", controllers.ListAllPartner)
}