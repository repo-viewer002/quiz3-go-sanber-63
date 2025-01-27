package users

import (
	"quiz3/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	repository := NewRepository()
	service := NewService(repository)
	controller := NewController(service)

	api := router.Group("/api/users")
	api.Use(middlewares.Log())
	{
		api.POST("", controller.CreateUserController)
		api.POST("/login", controller.LoginController)
	}
}
