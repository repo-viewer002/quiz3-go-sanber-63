package categories

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
	CreateCategoryController(ctx *gin.Context)
	GetAllCategoryController(ctx *gin.Context)
	GetCategoryByIdController(ctx *gin.Context)
	UpdateCategoryByIdController(ctx *gin.Context)
	DeleteCategoryByIdController(ctx *gin.Context)
}

type categoryController struct {
	service Service
}

func NewController(service Service) Controller {
	return &categoryController{
		service,
	}
}

// @tags Categories
// @Summary Create category
// @Description Create new category
// @Accept  json
// @Produce  json
// @Param category body swagger.CategoryInput true "Category data"
// @Success 201 {object} responses.BaseResponse{data=Category} "create category success"
// @Failure 400 {object} responses.BaseResponse "bad request errors"
// @Failure 401 {object} responses.BaseResponse "invalid credentials"
// @Router /api/categories [post]
func (controller *categoryController) CreateCategoryController(ctx *gin.Context) {
	_, username, err := middlewares.GetClaims(ctx)

	if err != nil {
		responses.GenerateUnauthorizedResponse(ctx, err.Error())

		return
	}

	var category Category

	if err := ctx.ShouldBindJSON(&category); err != nil {
		responses.GenerateBadRequestResponse(ctx, err.Error())

		return
	}

	category.Created_By = username
	category.Modified_By = username

	createdCategory, err := controller.service.CreateCategoryService(category)

	if err != nil {
		responses.GenerateBadRequestResponse(ctx, err.Error())

		return
	}

	responses.GenerateSuccessResponseWithData(ctx, http.StatusCreated, "create category success", createdCategory)
}

// @tags Categories
// @Summary Get all category
// @Description get all category
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.BaseResponse{data=[]Category} "get all category success"
// @Failure 400 {object} responses.BaseResponse "bad request errors"
// @Failure 401 {object} responses.BaseResponse "invalid credentials"
// @Router /api/categories [get]
func (controller *categoryController) GetAllCategoryController(ctx *gin.Context) {
	category, err := controller.service.GetAllCategoryService()

	if err != nil {
		responses.GenerateBadRequestResponse(ctx, err.Error())

		return
	}

	responses.GenerateSuccessResponseWithData(ctx, http.StatusOK, "get all category success", category)
}

// @tags Categories
// @Summary Get category by id
// @Description get category by ud
// @Accept  json
// @Produce  json
// @Param id query string true "category id"
// @Param Authorization header string true "JWT access token (default: Bearer accessToken)"
// @Success 200 {object} responses.BaseResponse{data=Category} "get category by id success"
// @Failure 400 {object} responses.BaseResponse "bad request errors"
// @Failure 401 {object} responses.BaseResponse "invalid credentials"
// @Router /api/categories/:id [get]
func (controller *categoryController) GetCategoryByIdController(ctx *gin.Context) {
	getId := ctx.Param("id")

	category, err := controller.service.GetCategoryByIdService(getId)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			responses.GenerateNotFoundResponse(ctx, err.Error())
		} else {
			responses.GenerateBadRequestResponse(ctx, err.Error())
		}

		return
	}

	responses.GenerateSuccessResponseWithData(ctx, http.StatusOK, fmt.Sprintf("get category by id \"%s\" success", getId), category)
}

// @tags Categories
// @Summary Update category by id
// @Description Update category by id
// @Accept  json
// @Produce  json
// @Param id query string true "category id"
// @Param category body swagger.CategoryInput true "Category data"
// @Param Authorization header string true "JWT access token (default: Bearer accessToken)"
// @Success 200 {object} responses.BaseResponse{data=Category} "update category by id success"
// @Failure 400 {object} responses.BaseResponse "bad request errors"
// @Failure 401 {object} responses.BaseResponse "invalid credentials"
// @Router /api/categories/:id [put]
func (controller *categoryController) UpdateCategoryByIdController(ctx *gin.Context) {
	_, username, err := middlewares.GetClaims(ctx)

	if err != nil {
		responses.GenerateUnauthorizedResponse(ctx, err.Error())

		return
	}

	var category Category

	getId := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&category); err != nil {
		responses.GenerateNotFoundResponse(ctx, err.Error())

		return
	}

	category.Modified_By = username
	updatedCategory, err := controller.service.UpdateCategoryByIdService(getId, category)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			responses.GenerateNotFoundResponse(ctx, err.Error())
		} else {
			responses.GenerateBadRequestResponse(ctx, err.Error())
		}

		return
	}

	responses.GenerateSuccessResponseWithData(ctx, http.StatusOK, fmt.Sprintf("update category by id \"%s\" success", getId), updatedCategory)
}

// @tags Categories
// @Summary Delete category by id
// @Description Delete category by id
// @Accept  json
// @Produce  json
// @Param id query string true "category id"
// @Param Authorization header string true "JWT access token (default: Bearer accessToken)"
// @Success 200 {object} responses.BaseResponse{data=Category} "delete category by id success"
// @Failure 400 {object} responses.BaseResponse "bad request errors"
// @Failure 401 {object} responses.BaseResponse "invalid credentials"
// @Router /api/categories/:id [delete]
func (controller *categoryController) DeleteCategoryByIdController(ctx *gin.Context) {
	getId := ctx.Param("id")

	deletedCategory, err := controller.service.DeleteCategoryByIdService(getId)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			responses.GenerateNotFoundResponse(ctx, err.Error())
		} else {
			responses.GenerateBadRequestResponse(ctx, err.Error())
		}

		return
	}

	responses.GenerateSuccessResponseWithData(ctx, http.StatusOK, fmt.Sprintf("delete category by id \"%s\" success", getId), deletedCategory)
}
