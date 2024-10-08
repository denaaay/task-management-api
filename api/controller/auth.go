package controller

import (
	"net/http"
	"strings"

	"github.com/denaaay/task-management-api/domain"
	"github.com/denaaay/task-management-api/model"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUseCase domain.AuthUsecase
	userUsecase domain.UserUsecase
}

func NewAuthController(authUseCase domain.AuthUsecase, userUseCase domain.UserUsecase) *AuthController {
	return &AuthController{authUseCase, userUseCase}
}

func (auth *AuthController) Register(c *gin.Context) {
	var newUser model.User
	err := c.ShouldBind(&newUser)
	if err != nil {
		resp := model.Response{
			StatusCode: 400,
			Message:    err.Error(),
			Data:       nil,
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	newUser.Name = strings.TrimSpace(newUser.Name)

	err = auth.authUseCase.Register(&newUser)
	if err != nil {
		resp := model.Response{
			StatusCode: 500,
			Message:    err.Error(),
			Data:       nil,
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}

	createResponse := model.CreateUserResp{
		Id:    int(newUser.ID),
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	resp := model.Response{
		StatusCode: 201,
		Message:    "success creating user",
		Data:       createResponse,
	}

	c.JSON(http.StatusCreated, resp)
}
