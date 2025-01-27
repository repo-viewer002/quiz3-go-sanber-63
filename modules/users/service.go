package users

import (
	"errors"
	"quiz3/middlewares"
)

type Service interface {
	CreateUserService(user User) (User, error)
	LoginService(user User) (string, error)
}

type userService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &userService{
		repository,
	}
}

func (service *userService) CreateUserService(user User) (User, error) {
	user.Created_By = user.Username
	user.Modified_By = user.Username

	createdUser, err := service.repository.CreateUserRepository(user)

	if err != nil {
		return User{}, err
	}

	return createdUser, nil
}

func (service *userService) LoginService(user User) (string, error) {
	validUser, err := service.repository.ValidateUsername(user.Username)

	if err != nil {
		return "", err
	}

	if user.Password != validUser.Password {
		return "", errors.New("invalid credentials")
	}

	token, err := middlewares.CreateToken(validUser.Id, validUser.Username)

	if err != nil {
		return "", err
	}

	return token, nil
}
