package routers

import (
	"fazztrack/demo/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(routerGroup *gin.RouterGroup){
	routerGroup.GET("", controllers.ListAllUsers)
	routerGroup.GET("/:id", controllers.DetailUser)
	// routerGroup.POST("", controllers.CreateUser)
	// routerGroup.PATCH("/:id", controllers.UpdateUser)
	// routerGroup.DELETE("/:id", controllers.DeleteUser)
}
