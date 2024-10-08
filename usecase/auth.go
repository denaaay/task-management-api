package usecase

import (
	"github.com/denaaay/task-management-api/domain"
)

type authUseCase struct {
	userRepository domain.UserRepository
}

func NewAuthUseCase(userRepository domain.UserRepository) domain.AuthUsecase {
	return &authUseCase{userRepository}
}
