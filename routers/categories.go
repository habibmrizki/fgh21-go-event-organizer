package routers

import (
	"fazztrack/demo/controllers"

	"github.com/gin-gonic/gin"
)

func CategoriesRouter(routerGroup *gin.RouterGroup){
	routerGroup.GET("", controllers.FindAllCategories)
	routerGroup.GET("/:id", controllers.DetailCategories)
	routerGroup.POST("", controllers.CreateCategories)
	routerGroup.PATCH("/:id", controllers.UpdateCategories)
	routerGroup.DELETE("/:id", controllers.DeleteCategories)
}