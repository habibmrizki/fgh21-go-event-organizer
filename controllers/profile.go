package controllers

import (
	"fazztrack/demo/lib"
	"fazztrack/demo/models"
	"fmt"
	"net/http"
	"strconv"

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
 
	 dataProfile := models.Createprofile(newProfile)
 
	 ctx.JSON(http.StatusOK, lib.Response{
		 Success: true,
		 Message: "User created successfully",
		 Results: dataProfile,
	 })
 }
 
 func DetailUserProfile(ctx *gin.Context) {
	userId := ctx.GetInt("userId")
	detailProfile := models.FindProfileByUserId(userId)
	fmt.Println(detailProfile)
	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Found",
		Results: detailProfile,
	})
 }

 func ListOneNational(ctx *gin.Context) {
    id,_ := strconv.Atoi(ctx.Param("id"))
    results := models.FindOneNational(id)
    ctx.JSON(http.StatusOK, lib.Response{
        Success: true,
        Message: "Id National",
        Results: results,
    })
}
func ListAllNational(ctx *gin.Context) {
    results := models.FindAllNationality()
    ctx.JSON(http.StatusOK, lib.Response{
        Success: true,
        Message: "List All National",
        Results: results,
    })
}