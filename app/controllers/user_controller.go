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
	"github.com/sirupsen/logrus"
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
	var validate *validator.Validate
	var request models.RequestLoginUser
	var ve validator.ValidationErrors
	validate = validator.New()
	ctx.ShouldBindJSON(&request)

	err := validate.Struct(request)
	var loginResponse models.ResponseLoginUser
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

	user, errorCode := controller.UserService.LoginUser(&request)
	logrus.Info(errorCode)
	switch errorCode {
	case "LOGIN_NOT_FOUND_404":
		// ctx.JSON(http.StatusNotFound, nil)
		helpers.APIResponse(ctx, "user acount is not registered", http.StatusNotFound, http.MethodPost, nil)
		return
	case "LOGIN_WRONG_PASSWORD_403":
		// ctx.JSON(http.StatusForbidden, nil)
		helpers.APIResponse(ctx, "user or password wrong", http.StatusForbidden, http.MethodPost, nil)
		return
	default:
		accessTokenData := map[string]interface{}{"user_name": user.UserName, "password": user.Password}
		accesToken, _ := util.Sign(accessTokenData, os.Getenv("JWT_SECRET"), 24*60*1)

		loginResponse.UserName = user.UserName
		loginResponse.AccessToken = accesToken
		helpers.APIResponse(ctx, "Login Success", http.StatusOK, http.MethodPost, loginResponse)
		return
	}

}
