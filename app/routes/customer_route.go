package routes

import (
	"app/controllers"
	"app/repositories"
	"app/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func InitCustomerRoute(db *gorm.DB, route *gin.Engine) {
	validate := validator.New()
	customerRepository := repositories.NewCustomerRepository(db)
	customerService := services.NewCustomerService(&customerRepository, validate)
	customerController := controllers.NewCustomerController(&customerService)
	groupRoute := route.Group("api/v1")
	groupRoute.GET("/customer", customerController.List)
	groupRoute.GET("/customer/:id", customerController.GetCustomer)
	groupRoute.POST("/customer", customerController.CreateCustomer)
	groupRoute.PUT("/customer/:id", customerController.EditCustomer)
	groupRoute.DELETE("/customer/:id", customerController.DeleteCustomer)
	groupRoute.GET("customer/log/test", customerController.TestLog)
}
