package routers

import (
	"fazztrack/demo/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup){
	r.POST("/login", controllers.AuthLogin)
	r.POST("/register", controllers.AuthRegister)
}