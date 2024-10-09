package repository

import (
	"github.com/denaaay/task-management-api/domain"
	"github.com/denaaay/task-management-api/model"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) domain.UserRepository {
	return &userRepository{DB}
}

func (usr *userRepository) CreateUser(user *model.User) error {
	if err := usr.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (usr *userRepository) GetUserByEmail(email string) (model.User, error) {
	var result model.User
	if err := usr.DB.Table("users").Select("*").Where("email = ?", email).Find(&result).Error; err != nil {
		return model.User{}, err
	}

	return result, nil
}
