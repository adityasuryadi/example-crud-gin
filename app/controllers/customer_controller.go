package controllers

import (
	"app/exception"
	"app/helpers"
	"app/models"
	"app/services"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

type CustomerController struct {
	CustomerService services.CustomerService
}

func NewCustomerController(customerService *services.CustomerService) CustomerController {
	return CustomerController{CustomerService: *customerService}
}

func (controller *CustomerController) List(ctx *gin.Context) {
	responses := controller.CustomerService.ListCustomers()
	response := helpers.WebResponse{
		Code:    200,
		Status:  "success",
		Message: "Sccess",
		Data:    responses,
	}
	ctx.JSON(http.StatusOK, response)
}

func (controller *CustomerController) GetCustomer(ctx *gin.Context) {
	id := ctx.Param("id")
	customerId, err := strconv.Atoi(id)
	exception.PanicIfNeeded(err)
	customer := controller.CustomerService.ListCustomer(int64(customerId))
	response := helpers.WebResponse{
		Code:    200,
		Status:  "success",
		Message: "Success",
		Data:    customer,
	}
	ctx.JSON(http.StatusOK, response)
}

func (controller *CustomerController) CreateCustomer(ctx *gin.Context) {
	var request models.CreateCustomerRequest
	// ctx.ShouldBind(&request) //handle json
	// // ctx.PostForm("name") //handle form
	if err := ctx.ShouldBindJSON(&request); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]helpers.ErrorMessage, len(ve))
			for i, fe := range ve {
				out[i] = helpers.ErrorMessage{fe.Field(), helpers.GetErrorMsg(fe)}
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}
	controller.CustomerService.SaveCustomer(request)

}

func (controller *CustomerController) EditCustomer(ctx *gin.Context) {
	var request models.CreateCustomerRequest
	customerId, _ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.ShouldBindJSON(&request); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]helpers.ErrorMessage, len(ve))
			for i, fe := range ve {
				out[i] = helpers.ErrorMessage{fe.Field(), helpers.GetErrorMsg(fe)}
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}
	controller.CustomerService.UpdateCustomer(int64(customerId), request)
}

func (controller *CustomerController) DeleteCustomer(ctx *gin.Context) {
	customerId, _ := strconv.Atoi(ctx.Param("id"))
	controller.CustomerService.DeleteCustomer(int64(customerId))
}

// untuk belajar log
func (controller *CustomerController) TestLog(ctx *gin.Context) {
	responses := controller.CustomerService.ListCustomers()
	log.Error(responses)
}
