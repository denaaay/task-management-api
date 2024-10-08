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

func (usr *userRepository) CreateUser(user *model.User) (model.CreateUserResp, error) {
	return model.CreateUserResp{}, nil
}
