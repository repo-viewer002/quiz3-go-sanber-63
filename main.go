package main

import (
	"fmt"
	"quiz3/commons"
	"quiz3/configs/database"
	"quiz3/middlewares"
	"quiz3/modules/books"
	"quiz3/modules/categories"
	"quiz3/modules/users"
	"quiz3/utils/swagger"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitializeDB()

	router := gin.Default()
	router.Use(middlewares.Log())

	categories.CategoryRouter(router)
	books.BookRouter(router)
	users.UserRouter(router)

	swagger.Initiator(router)

	router.Run(fmt.Sprintf(":%d", commons.PORT))
}
