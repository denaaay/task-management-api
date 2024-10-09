package domain

import "github.com/denaaay/task-management-api/model"

type AuthUsecase interface {
	Register(user *model.User) error
	Login(user model.LoginReq) (model.UserPayload, error)

	RegisInputCheck(email, password string) error
}
