package usecase

import (
	"fmt"

	"github.com/denaaay/task-management-api/domain"
	"github.com/denaaay/task-management-api/model"
	"github.com/denaaay/task-management-api/utils"
)

type authUseCase struct {
	userRepository domain.UserRepository
}

func NewAuthUseCase(userRepository domain.UserRepository) domain.AuthUsecase {
	return &authUseCase{userRepository}
}

func (auth *authUseCase) Register(user *model.User) error {
	getUser, err := auth.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		return fmt.Errorf("error getting user by email")
	}

	if user.Email == getUser.Email {
		return fmt.Errorf("error email already registered")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("error hashing password")
	}

	user.Password = hashedPassword

	err = auth.userRepository.CreateUser(user)
	return err
}

func (auth *authUseCase) RegisInputCheck(email, password string) error {
	if email == "" {
		return fmt.Errorf("email cannot null")
	}

	if password == "" {
		return fmt.Errorf("password cannot null")
	}

	if len(password) < 6 {
		return fmt.Errorf("password cannot less than 6 character")
	}

	return nil
}

func (auth *authUseCase) Login(user model.LoginReq) (model.UserPayload, error) {
	userResp, err := auth.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		return model.UserPayload{}, err
	}

	if userResp.ID == 0 {
		return model.UserPayload{}, fmt.Errorf("user not found")
	}

	err = utils.VerifyPassword(userResp.Password, user.Password)
	if err != nil {
		return model.UserPayload{}, fmt.Errorf("wrong password")
	}

	result := model.UserPayload{
		Id:   int(userResp.ID),
		Name: userResp.Name,
	}

	return result, nil
}
