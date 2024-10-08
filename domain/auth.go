package domain

import "github.com/denaaay/task-management-api/model"

type AuthUsecase interface {
	Register(user *model.User) (model.CreateUserResp, error)
}
