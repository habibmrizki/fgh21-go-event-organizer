package controllers

import (
	"fazztrack/demo/lib"
	"fazztrack/demo/models"
	"fmt"

	// "fmt"

	// "fmt"
	"net/http"
	// "strconv"

	"github.com/gin-gonic/gin"
)

type Token struct {
	JWToken string `json:"token"`
}


type FormRegister struct {
	FullName string `form:"fullName"`
	Email string `form:"email"`
	Password string `form:"password"`
	ConfirmPassword string `form:"confirmPassword" binding:"eqfield=Password"`
}

func AuthLogin(ctx *gin.Context) {
	var user models.User
	ctx.Bind(&user)

	found := models.FindOneUserByEmail(user.Email)

	if found == (models.User{}){
		ctx.JSON(http.StatusUnauthorized, lib.Response{
			Success: false,
			Message: "Wrong Email or Password",
		})
		return
	}

	isVerified := lib.Verify(user.Password, found.Password)
	if isVerified {
		JWToken := lib.GenerateUserIdToken(found.Id)
		ctx.JSON(http.StatusOK, lib.Response{
			Success: true,
			Message: "login Success",
			Results: Token{
				JWToken,
			},
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, lib.Response{
			Success: false,
			Message: "Wrong Email or Password",
		})
	}
}

func AuthRegister(ctx *gin.Context) {
	form := FormRegister{}
	user := models.User{}
	profile := models.Profile{}

	err := ctx.Bind(&form)
	if err != nil {
		fmt.Println(nil)
	}

	user.Email = form.Email
	user.Password = form.Password
	profile.FullName = form.FullName
	createUser:= models.CreateUser(user)

	userId := createUser.Id
	profile.UserId = userId
	
	createProfile:= models.Createprofile(profile)
	createProfile.Email = form.Email
	createProfile.FullName = form.FullName

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Create success",
		Results: createProfile,
	})

}