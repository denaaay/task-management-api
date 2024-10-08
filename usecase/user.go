package usecase

import "github.com/denaaay/task-management-api/domain"

type userUseCase struct {
	userRepository domain.UserRepository
}

func NewUserUseCase(userRepository domain.UserRepository) domain.UserUsecase {
	return &userUseCase{userRepository}
}
