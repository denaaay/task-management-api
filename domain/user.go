package domain

import "github.com/denaaay/task-management-api/model"

type UserRepository interface {
	CreateUser(user *model.User) (model.CreateUserResp, error)
}

type UserUsecase interface {
}
