package router

import (
	"net/http"

	"github.com/denaaay/task-management-api/api/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(authController *controller.AuthController) *gin.Engine {
	service := gin.Default()

	service.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hallo")
	})

	route := service.Group("api")

	route.POST("/register", authController.Register)
	route.POST("/login", authController.Login)

	return service
}
