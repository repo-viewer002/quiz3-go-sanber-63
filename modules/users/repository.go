package users

import (
	"database/sql"
	"errors"
	"quiz3/configs/database"
)

type Repository interface {
	CreateUserRepository(user User) (User, error)
	ValidateUsername(username string) (User, error)
}

type userRepository struct{}

func NewRepository() Repository {
	return &userRepository{}
}

func (repository *userRepository) CreateUserRepository(user User) (User, error) {
	query := `
		INSERT INTO users
		(
			username, 
			password, 
			created_by,
			modified_by
		)
		VALUES
		($1, $2, $3, $4)
		RETURNING *
	`

	err := database.DB.QueryRow(query, user.Username, user.Password, user.Created_By, user.Modified_By).
		Scan(&user.Id, &user.Username, &user.Password, &user.Created_At, &user.Created_By, &user.Modified_At, &user.Modified_By)

	if err != nil {
		return User{}, err
	}

	return user, err
}

func (repository *userRepository) ValidateUsername(username string) (User, error) {
	var user User

	query := `
		SELECT
			id,
			username, 
			password 
		FROM users
		WHERE username = $1
	`

	err := database.DB.QueryRow(query, username).
		Scan(&user.Id, &user.Username, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, errors.New("invalid credentials")
		}
		return User{}, err
	}

	return user, err
}
