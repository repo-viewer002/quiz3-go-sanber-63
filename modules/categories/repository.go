package categories

import (
	"database/sql"
	"fmt"
	"quiz3/configs/database"
)

type Repository interface {
	CreateCategoryRepository(category Category) (Category, error)
	GetAllCategoryRepository() ([]Category, error)
	GetCategoryByIdRepository(id int) (Category, error)
	UpdateCategoryByIdRepository(id int, category Category) (Category, error)
	DeleteCategoryByIdRepository(id int) (Category, error)
}

type categoryRepository struct{}

func NewRepository() Repository {
	return &categoryRepository{}
}

func (repository *categoryRepository) CreateCategoryRepository(category Category) (Category, error) {
	query := `
		INSERT INTO categories
		(
			name, 
			created_by, 
			modified_by
		)
		VALUES
		($1, $2, $3)
		RETURNING *
	`

	err := database.DB.QueryRow(query, category.Name, category.Created_By, category.Modified_By).
		Scan(&category.Id, &category.Name, &category.Created_At, &category.Created_By, &category.Modified_At, &category.Modified_By)

	if err != nil {
		return Category{}, err
	}

	return category, err
}

func (repository *categoryRepository) GetAllCategoryRepository() ([]Category, error) {
	var categories []Category

	query := "SELECT * FROM categories"

	rows, err := database.DB.Query(query)

	if err != nil {
		return []Category{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var category Category

		err = rows.Scan(&category.Id, &category.Name, &category.Created_At, &category.Created_By, &category.Modified_At, &category.Modified_By)

		if err != nil {
			return []Category{}, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (repository *categoryRepository) GetCategoryByIdRepository(id int) (Category, error) {
	var category Category

	query := `
		SELECT * FROM categories 
		WHERE id = $1
	`

	err := database.DB.QueryRow(query, id).
		Scan(&category.Id, &category.Name, &category.Created_At, &category.Created_By, &category.Modified_At, &category.Modified_By)

	if err != nil {
		if err == sql.ErrNoRows {
			return Category{}, fmt.Errorf("failed to get category data, category with id \"%d\" not found", id)
		}

		return Category{}, err
	}

	return category, nil
}

func (repository *categoryRepository) UpdateCategoryByIdRepository(id int, category Category) (Category, error) {
	query := `
		UPDATE categories 
		SET 
			name = $2, 
			modified_by = $3 
		WHERE id = $1 
		RETURNING *
	`

	err := database.DB.QueryRow(query, id, category.Name, category.Modified_By).
		Scan(&category.Id, &category.Name, &category.Created_At, &category.Created_By, &category.Modified_At, &category.Modified_By)

	if err != nil {
		if err == sql.ErrNoRows {
			return category, fmt.Errorf("failed updating category, category with id \"%d\" not found", id)
		}

		return Category{}, err
	}

	return category, nil
}

func (repository *categoryRepository) DeleteCategoryByIdRepository(id int) (Category, error) {
	var deletedCategory Category

	query := `
		DELETE FROM categories 
		WHERE id = $1 
		RETURNING *
	`

	err := database.DB.QueryRow(query, id).
		Scan(&deletedCategory.Id, &deletedCategory.Name, &deletedCategory.Created_At, &deletedCategory.Created_By, &deletedCategory.Modified_At, &deletedCategory.Modified_By)

	if err != nil {
		if err == sql.ErrNoRows {
			return deletedCategory, fmt.Errorf("failed deleting category, category with id \"%d\" not found", id)
		}

		return Category{}, err
	}

	return deletedCategory, nil
}
