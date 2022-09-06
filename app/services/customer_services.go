package services

import (
	"app/models"
)

type CustomerService interface {
	ListCustomers() (responses []models.GetCustomersResponse)
	ListCustomer(customerId int64) (response models.GetCustomersResponse)
	SaveCustomer(models.CreateCustomerRequest)
	UpdateCustomer(customerId int64, request models.CreateCustomerRequest)
	DeleteCustomer(customerId int64)
}
