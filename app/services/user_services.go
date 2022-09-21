package services

import (
	"app/entity"
	"app/models"
)

type UserService interface {
	CreateUser(request *models.RequestCreateUser)
	LoginUser(request *models.RequestLoginUser) (*entity.User, string)
}
