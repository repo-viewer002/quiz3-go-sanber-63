package main

import (
	"fmt"
	"net/http"
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

	router.GET("/", indexController)

	categories.CategoryRouter(router)
	books.BookRouter(router)
	users.UserRouter(router)

	swagger.Initiator(router)

	router.Run(fmt.Sprintf(":%d", commons.PORT))
}

func indexController(ctx *gin.Context) {
	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}

	host := ctx.Request.Host

	apiDocumentation := scheme + "://" + host + "/swagger/index.html"

	ctx.Data(http.StatusOK, "text/html", []byte(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>API Documentation</title>
			</head>
			<body>
				<h1>Sanbercode Golang Backend Development Batch 63 | Quiz 3</h1>
				<a href="`+apiDocumentation+`" target="_blank">API Documentation</a></br>
				<a href="`+commons.REPOSITORY+`" target="_blank">Github Repository</a></br>
				<a href="`+commons.ENDPOINT+`" target="_blank">Railway Deployment URL</a>
			</body>
			</html>
		`))
}
