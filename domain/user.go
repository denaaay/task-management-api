package domain

import "github.com/denaaay/task-management-api/model"

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByEmail(email string) (model.User, error)
}

type UserUsecase interface {
	GetUserByEmail(email string) (model.User, error)
}
