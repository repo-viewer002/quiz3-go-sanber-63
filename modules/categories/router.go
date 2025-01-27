package categories

import (
	"quiz3/middlewares"

	"github.com/gin-gonic/gin"
)

func CategoryRouter(router *gin.Engine) {
	repository := NewRepository()
	service := NewService(repository)
	controller := NewController(service)

	api := router.Group("/api/categories")
	api.Use(middlewares.JwtMiddleware())
	{
		api.POST("", controller.CreateCategoryController)
		api.GET("", controller.GetAllCategoryController)
		api.GET("/:id", controller.GetCategoryByIdController)
		api.PUT("/:id", controller.UpdateCategoryByIdController)
		api.DELETE("/:id", controller.DeleteCategoryByIdController)
	}
}
