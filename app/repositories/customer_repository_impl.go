package repositories

import (
	"app/entity"
	"app/exception"
	"app/models"

	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &CustomerRepositoryImpl{
		db: db,
	}
}

// FindAll implements CustomerRepository
func (repository *CustomerRepositoryImpl) FindAll() (customers []entity.Customer) {
	var customer []entity.Customer
	result := repository.db.Find(&customer)
	exception.PanicIfNeeded(result.Error)
	for _, value := range customer {
		customers = append(customers, entity.Customer{
			Id:        value.Id,
			Name:      value.Name,
			Email:     value.Email,
			Balance:   value.Balance,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		})
	}
	return customers
}

func (repository *CustomerRepositoryImpl) FindById(customerId int64) (customer entity.Customer) {
	db := repository.db.Model(&customer)
	db.Select("*").First(&customer, customerId)
	return customer
}

func (repository *CustomerRepositoryImpl) Create(request models.CreateCustomerRequest) {

	// var customer entity.Customer
	customer := entity.Customer{
		Name:    request.Name,
		Email:   request.Email,
		Balance: request.Balance,
	}
	db := repository.db.Model(&customer)

	err := db.Create(&customer).Error
	if err != nil {
		exception.PanicIfNeeded(err)
	}
}

func (repository *CustomerRepositoryImpl) Update(customerId int64, request models.CreateCustomerRequest) {
	var customer entity.Customer

	db := repository.db.Model(&customer)

	checkCustomer := db.First(&customer, customerId)
	if checkCustomer.RowsAffected < 1 {
		return
	}

	customer = entity.Customer{
		Name:    request.Name,
		Email:   request.Email,
		Balance: request.Balance,
	}
	db.Updates(customer)
}

func (repository *CustomerRepositoryImpl) Delete(customerId int64) {
	var customer entity.Customer
	db := repository.db.Model(&customer)
	// db.Select("*").Where("id = ?", customerId).Find(&customer).Delete(&customer)
	db.Delete(&customer, customerId)
}
