package usecase

import (
	"github.com/denaaay/task-management-api/domain"
	"github.com/denaaay/task-management-api/model"
)

type authUseCase struct {
	userRepository domain.UserRepository
}

func NewAuthUseCase(userRepository domain.UserRepository) domain.AuthUsecase {
	return &authUseCase{userRepository}
}

func (auth *authUseCase) Register(user *model.User) error {
	err := auth.userRepository.CreateUser(user)
	return err
}
