package services

import (
	"app/entity"
	"app/exception"
	"app/models"
	"app/repositories"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	Validate       *validator.Validate
}

func NewUserService(userRepository *repositories.UserRepository, validate validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: *userRepository,
		Validate:       &validate,
	}
}

// CreateUser implements UserService
func (service *UserServiceImpl) CreateUser(request *models.RequestCreateUser) {
	// service.Validate.RegisterValidation("UserName", ValidateUniqueUser)
	err := service.Validate.Struct(request)
	exception.PanicIfNeeded(err)
	service.UserRepository.Create(*request)
}

func (service *UserServiceImpl) LoginUser(request *models.RequestLoginUser) (*entity.User, string) {
	result, errorCode := service.UserRepository.Login(*request)
	return result, errorCode
}
