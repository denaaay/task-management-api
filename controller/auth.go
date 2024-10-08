package controller

import "github.com/denaaay/task-management-api/domain"

type AuthController struct {
	authUseCase domain.AuthUsecase
	userUsecase domain.UserUsecase
}

func NewAuthController(authUseCase domain.AuthUsecase, userUseCase domain.UserUsecase) *AuthController {
	return &AuthController{authUseCase, userUseCase}
}
