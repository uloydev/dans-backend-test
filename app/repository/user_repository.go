package repository

import (
	"dans-backend-test/app/entity"
	"dans-backend-test/exception"

	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB         *gorm.DB
	HTTPClient *resty.Client
}

func NewUserRepository(db *gorm.DB, HTTPClient *resty.Client) BaseRepository[entity.User] {
	return &UserRepository{
		DB:         db,
		HTTPClient: HTTPClient,
	}
}

func (repo *UserRepository) Insert(user entity.User) entity.User {
	result := repo.DB.Create(&user)
	exception.PanicWhenError(result.Error)
	return user
}

func (repo *UserRepository) FindAll() (users []entity.User) {
	result := repo.DB.Find(&users)
	exception.PanicWhenError(result.Error)
	return users
}

func (repo *UserRepository) FindById(ID uint) (user entity.User) {
	result := repo.DB.Where("id = ?", ID).First(&user)
	exception.PanicValidationWhenError(result.Error)
	return user
}

func (repo *UserRepository) FindByUsername(username string) (user entity.User) {
	result := repo.DB.Where("username = ?", username).First(&user)
	exception.PanicValidationWhenError(result.Error)
	return user
}
