package middlewares

import (
	"fazztrack/demo/lib"
	"net/http"

	"github.com/gin-gonic/gin"
)


func TokenFailed(ctx *gin.Context){
	if e := recover(); e != nil{
		ctx.JSON(http.StatusUnauthorized, lib.Response{
			Success: false,
			Message: "Not Found",
		})
		ctx.Abort()
	}
}





func AuthMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		defer TokenFailed(ctx)
		token := ctx.GetHeader("Authorization")[7:] // mengeluarkan string (berarer token) dan gagal melakukan pemotongan string
		isValidated, userId := lib.ValidateToken(token)
		if isValidated {
			ctx.Set("userId", userId)
			ctx.Next()
		} else {
			panic("Error: token invalid")
		}
	}
}