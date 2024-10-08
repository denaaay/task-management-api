package domain

import "github.com/denaaay/task-management-api/model"

type UserRepository interface {
	CreateUser(user *model.User) error
}

type UserUsecase interface {
}
