package repositories

import (
	"app/entity"
	"app/models"
)

type CustomerRepository interface {
	FindAll() (customers []entity.Customer)
	FindById(customerId int64) (customer entity.Customer)
	Create(request models.CreateCustomerRequest)
	Update(customerId int64, request models.CreateCustomerRequest)
	Delete(customerId int64)
}
