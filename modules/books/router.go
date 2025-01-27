package books

import (
	"quiz3/middlewares"

	"github.com/gin-gonic/gin"
)

func BookRouter(router *gin.Engine) {
	repository := NewRepository()
	service := NewService(repository)
	controller := NewController(service)

	api := router.Group("/api/books")
	api.Use(middlewares.JwtMiddleware())
	{
		api.POST("", controller.CreateBookController)
		api.GET("", controller.GetAllBookController)
		api.GET("/:id", controller.GetBookByIdController)
		api.PUT("/:id", controller.UpdateBookByIdController)
		api.DELETE("/:id", controller.DeleteBookByIdController)
	}
}
