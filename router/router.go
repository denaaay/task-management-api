package router

import (
	"net/http"

	"github.com/denaaay/task-management-api/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(authController *controller.AuthController) *gin.Engine {
	service := gin.Default()

	service.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hallo")
	})

	return service
}
