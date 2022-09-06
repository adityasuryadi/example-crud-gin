package controllers

import (
	"app/helpers"
	"app/models"
	"app/services"
	util "app/utils"
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService *services.UserService) UserController {
	return UserController{UserService: *userService}
}

func (controller *UserController) CreateUser(ctx *gin.Context) {
	var request models.RequestCreateUser
	ctx.ShouldBindJSON(&request)
	controller.UserService.CreateUser(&request)
}

func (controller *UserController) Login(ctx *gin.Context) {
	var request models.RequestLoginUser

	ctx.ShouldBindJSON(&request)

	var ve validator.ValidationErrors
	var loginResponse models.ResponseLoginUser
	user, err := controller.UserService.LoginUser(&request)

	if err != nil {
		if errors.As(err, &ve) {
			out := make([]helpers.ErrorMessage, len(ve))
			for i, fe := range ve {
				out[i] = helpers.ErrorMessage{fe.Field(), helpers.GetErrorMsg(fe)}
			}
			response := helpers.WebResponse{
				Code:    500,
				Status:  "Login Failure",
				Message: err.Error(),
				Data:    out,
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
		return
	}

	accessTokenData := map[string]interface{}{"user_name": user.UserName, "password": user.Password}
	accesToken, _ := util.Sign(accessTokenData, os.Getenv("JWT_SECRET"), 24*60*1)

	loginResponse.UserName = user.UserName
	loginResponse.AccessToken = accesToken

	response := helpers.WebResponse{
		Code:    200,
		Status:  "success",
		Message: "Success",
		Data:    loginResponse,
	}
	ctx.JSON(http.StatusOK, response)

}
