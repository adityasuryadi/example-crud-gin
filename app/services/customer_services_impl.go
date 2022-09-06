package services

import (
	"app/models"
	"app/repositories"

	"github.com/go-playground/validator/v10"
)

type customerServiceImpl struct {
	CustomerRepository repositories.CustomerRepository
	Validate           *validator.Validate
}

func NewCustomerService(customerRepository *repositories.CustomerRepository, validate *validator.Validate) CustomerService {
	return &customerServiceImpl{
		CustomerRepository: *customerRepository,
		Validate:           validate,
	}
}

func (service *customerServiceImpl) ListCustomers() (responses []models.GetCustomersResponse) {
	customers := service.CustomerRepository.FindAll()
	for _, customer := range customers {
		responses = append(responses, models.GetCustomersResponse{
			Id:        customer.Id,
			Name:      customer.Name,
			Email:     customer.Email,
			Balance:   customer.Balance,
			CreatedAt: customer.CreatedAt,
			UpdatedAt: customer.UpdatedAt,
		})
	}
	return responses
}

func (service *customerServiceImpl) ListCustomer(customerId int64) (response models.GetCustomersResponse) {
	customer := service.CustomerRepository.FindById(customerId)
	response.Id = customer.Id
	response.Name = customer.Name
	response.Email = customer.Email
	response.Balance = customer.Balance
	response.CreatedAt = customer.CreatedAt
	response.UpdatedAt = customer.UpdatedAt

	return response
}

func (service *customerServiceImpl) SaveCustomer(request models.CreateCustomerRequest) {
	service.CustomerRepository.Create(request)
}

func (service *customerServiceImpl) UpdateCustomer(customerId int64, request models.CreateCustomerRequest) {
	service.CustomerRepository.Update(customerId, request)
}

func (service *customerServiceImpl) DeleteCustomer(customerId int64) {
	service.CustomerRepository.Delete(customerId)
}
