package repositories

import (
	"app/entity"
	"app/models"
)

type UserRepository interface {
	Login(request models.RequestLoginUser) (*entity.User, error)
	Create(request models.RequestCreateUser)
}
