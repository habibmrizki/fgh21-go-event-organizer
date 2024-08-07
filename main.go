package main

import (
	"fazztrack/demo/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()	
	r.Use(cors.Default())

	routers.RouterCombine(r)

	r.Run("localhost:8080")
}