package repositories

import (
	"app/entity"
	"app/exception"
	"app/models"
	util "app/utils"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// Create implements UserRepository
func (repository *UserRepositoryImpl) Create(request models.RequestCreateUser) {
	user := entity.User{
		UserName: request.UserName,
		Password: getHash([]byte(request.Password)),
	}

	db := repository.db.Model(&user)
	err := db.Create(&user).Error
	if err != nil {
		exception.PanicIfNeeded(err)
	}
}

// Login implements UserRepository
func (repository *UserRepositoryImpl) Login(request models.RequestLoginUser) (*entity.User, error) {
	var user entity.User
	db := repository.db.Model(&user)
	// checkUserAccount := db.Debug().Select("*").Where("user_name = ?", request.UserName).Find(&user)
	checkUserAccount := db.Where("user_name = ? ", request.UserName).First(&user)

	if checkUserAccount.RowsAffected < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	comparePassword := util.ComparePassword(user.Password, request.Password)
	if comparePassword != nil {
		return nil, comparePassword
	}

	return &user, nil

}
