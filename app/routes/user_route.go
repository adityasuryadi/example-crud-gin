package routes

import (
	"app/controllers"
	"app/repositories"
	"app/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func InitUserRoute(db *gorm.DB, route *gin.Engine) {
	validate := validator.New()
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(&userRepository, *validate)
	userController := controllers.NewUserController(&userService)
	groupRoute := route.Group("api/v1")
	groupRoute.POST("/user", userController.CreateUser)
	groupRoute.POST("/login", userController.Login)
}
