package books

import (
	"fmt"
	"net/http"
	"quiz3/commons/responses"
	"quiz3/middlewares"
	_ "quiz3/utils/swagger"
	"strings"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	CreateBookController(ctx *gin.Context)
	GetAllBookController(ctx *gin.Context)
	GetBookByIdController(ctx *gin.Context)
	UpdateBookByIdController(ctx *gin.Context)
	DeleteBookByIdController(ctx *gin.Context)
}

type bookController struct {
	service Service
}

func NewController(service Service) Controller {
	return &bookController{
		service,
	}
}

// @tags Books
// @Summary Create book
// @Description Create new book
// @Accept  json
// @Produce  json
// @Param book body swagger.BookInput true "Category data"
// @Success 201 {object} responses.BaseResponse{data=Book} "create book success"
// @Failure 400 {object} responses.BaseResponse "bad request errors"
// @Failure 401 {object} responses.BaseResponse "invalid credentials"
// @Router /api/books [post]
func (controller *bookController) CreateBookController(ctx *gin.Context) {
	_, username, err := middlewares.GetClaims(ctx)

	if err != nil {
		responses.GenerateUnauthorizedResponse(ctx, err.Error())

		return
	}

	var book Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		responses.GenerateBadRequestResponse(ctx, err.Error())

		return
	}

	book.Created_By = username
	book.Modified_By = username
	createdBook, err := controller.service.CreateBookService(book)

	if err != nil {
		if strings.Contains(err.Error(), "fk_category") {
			responses.GenerateBadRequestResponse(ctx, "invalid category")
		} else {
			responses.GenerateBadRequestResponse(ctx, err.Error())
		}

		return
	}

	responses.GenerateSuccessResponseWithData(ctx, http.StatusCreated, "create book success", createdBook)
}

// @tags Books
// @Summary Get all book
// @Description get all book
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.BaseResponse{data=Book} "get all book success"
// @Failure 400 {object} responses.BaseResponse "bad request errors"
// @Failure 401 {object} responses.BaseResponse "invalid credentials"
// @Router /api/books [get]
func (controller *bookController) GetAllBookController(ctx *gin.Context) {
	book, err := controller.service.GetAllBookService()

	if err != nil {
		responses.GenerateBadRequestResponse(ctx, err.Error())

		return
	}

	responses.GenerateSuccessResponseWithData(ctx, http.StatusOK, "get all book success", book)
}

// @tags Books
// @Summary Get book by id
// @Description get book by ud
// @Accept  json
// @Produce  json
// @Param id query string true "book id"
// @Param Authorization header string true "JWT access token (default: Bearer accessToken)"
// @Success 200 {object} responses.BaseResponse{data=[]Book} "get book by id success"
// @Failure 400 {object} responses.BaseResponse "bad request errors"
// @Failure 401 {object} responses.BaseResponse "invalid credentials"
// @Router /api/books/:id [get]
func (controller *bookController) GetBookByIdController(ctx *gin.Context) {
	getId := ctx.Param("id")

	book, err := controller.service.GetBookByIdService(getId)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			responses.GenerateNotFoundResponse(ctx, err.Error())
		} else {
			responses.GenerateBadRequestResponse(ctx, err.Error())
		}

		return
	}

	responses.GenerateSuccessResponseWithData(ctx, http.StatusOK, fmt.Sprintf("get book by id \"%s\" success", getId), book)
}

// @tags Books
// @Summary Update book by id
// @Description Update book by id
// @Accept  json
// @Produce  json
// @Param id query string true "book id"
// @Param book body swagger.CategoryInput true "Category data"
// @Param Authorization header string true "JWT access token (default: Bearer accessToken)"
// @Success 200 {object} responses.BaseResponse{data=Book} "update book by id success"
// @Failure 400 {object} responses.BaseResponse "bad request errors"
// @Failure 401 {object} responses.BaseResponse "invalid credentials"
// @Router /api/books/:id [put]
func (controller *bookController) UpdateBookByIdController(ctx *gin.Context) {
	_, username, err := middlewares.GetClaims(ctx)

	if err != nil {
		responses.GenerateUnauthorizedResponse(ctx, err.Error())

		return
	}

	var book Book

	getId := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&book); err != nil {
		responses.GenerateNotFoundResponse(ctx, err.Error())

		return
	}

	book.Modified_By = username
	updatedBook, err := controller.service.UpdateBookByIdService(getId, book)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			responses.GenerateNotFoundResponse(ctx, err.Error())
		} else if strings.Contains(err.Error(), "fk_category") {
			responses.GenerateBadRequestResponse(ctx, "invalid category")
		} else {
			responses.GenerateBadRequestResponse(ctx, err.Error())
		}

		return
	}

	responses.GenerateSuccessResponseWithData(ctx, http.StatusOK, fmt.Sprintf("update book by id \"%s\" success", getId), updatedBook)
}

// @tags Books
// @Summary Delete book by id
// @Description Delete book by id
// @Accept  json
// @Produce  json
// @Param id query string true "book id"
// @Param Authorization header string true "JWT access token (default: Bearer accessToken)"
// @Success 200 {object} responses.BaseResponse{data=Book} "delete book by id success"
// @Failure 400 {object} responses.BaseResponse "bad request errors"
// @Failure 401 {object} responses.BaseResponse "invalid credentials"
// @Router /api/books/:id [delete]
func (controller *bookController) DeleteBookByIdController(ctx *gin.Context) {
	getId := ctx.Param("id")

	deletedBook, err := controller.service.DeleteBookByIdService(getId)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			responses.GenerateNotFoundResponse(ctx, err.Error())
		} else {
			responses.GenerateBadRequestResponse(ctx, err.Error())
		}

		return
	}

	responses.GenerateSuccessResponseWithData(ctx, http.StatusOK, fmt.Sprintf("delete book by id \"%s\" success", getId), deletedBook)
}
