package controllers

import (
	"fazztrack/demo/lib"
	"fazztrack/demo/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProfile(ctx *gin.Context) {
	newProfile := models.Profile{}
	fmt.Println(newProfile)
	 if err := ctx.ShouldBind(&newProfile); 
		 err != nil {
		 ctx.JSON(http.StatusBadRequest, lib.Response{
			 Success: false,
			 Message: "Invalid input data",
		 })
		 return
	 }
 
	 dataProfile := models.CreateProfile(newProfile)
 
	 ctx.JSON(http.StatusOK, lib.Response{
		 Success: true,
		 Message: "User created successfully",
		 Results: dataProfile,
	 })
 }