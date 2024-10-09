package usecase

import (
	"fmt"

	"github.com/denaaay/task-management-api/domain"
	"github.com/denaaay/task-management-api/model"
)

type userUseCase struct {
	userRepository domain.UserRepository
}

func NewUserUseCase(userRepository domain.UserRepository) domain.UserUsecase {
	return &userUseCase{userRepository}
}

func (usr *userUseCase) GetUserByEmail(email string) (model.User, error) {
	result, err := usr.userRepository.GetUserByEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("user not found")
	}
	return result, nil
}
