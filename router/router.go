package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	service := gin.Default()

	service.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hallo")
	})

	return service
}
