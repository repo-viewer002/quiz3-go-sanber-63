package users

import (
	"net/http"
	"quiz3/commons/responses"
	_ "quiz3/utils/swagger"
	"strings"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	CreateUserController(ctx *gin.Context)
	LoginController(ctx *gin.Context)
}

type userController struct {
	service Service
}

func NewController(service Service) Controller {
	return &userController{
		service,
	}
}

// @tags Users & Authentications
// @Summary Create user
// @Description Creating a new user
// @Accept  json
// @Produce  json
// @Param user body swagger.UserInput true "User data"
// @Success 201 {object} responses.BaseResponse{data=User} "create user success"
// @Failure 400 {object} responses.BaseResponse "bad request errors"
// @Failure 400 {object} responses.BaseResponse "duplicated username"
// @Router /api/users [post]
func (controller *userController) CreateUserController(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		responses.GenerateBadRequestResponse(ctx, err.Error())

		return
	}

	createdUser, err := controller.service.CreateUserService(user)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			responses.GenerateBadRequestResponse(ctx, "duplicated username")
		} else {
			responses.GenerateBadRequestResponse(ctx, err.Error())
		}

		return
	}

	responses.GenerateSuccessResponseWithData(ctx, http.StatusCreated, "create user success", createdUser)
}

// @tags Users & Authentications
// @Summary Login
// @Description User login
// @Accept  json
// @Produce  json
// @Param user body swagger.UserInput true "User data"
// @Success 200 {object} responses.BaseResponse{data=string} "login success"
// @Failure 400 {object} responses.BaseResponse "bad request errors"
// @Failure 401 {object} responses.BaseResponse "invalid credentials"
// @Router /api/users/login [post]
func (controller *userController) LoginController(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		responses.GenerateBadRequestResponse(ctx, err.Error())

		return
	}

	token, err := controller.service.LoginService(user)

	if err != nil {
		if strings.Contains(err.Error(), "invalid credentials") {
			responses.GenerateUnauthorizedResponse(ctx, err.Error())
		} else {
			responses.GenerateBadRequestResponse(ctx, err.Error())
		}
		return
	}

	responses.GenerateSuccessResponseWithData(ctx, http.StatusOK, "login success", token)
}
