package books

import (
	"database/sql"
	"fmt"
	"quiz3/configs/database"
)

type Repository interface {
	CreateBookRepository(book Book) (Book, error)
	GetAllBookRepository() ([]Book, error)
	GetBookByIdRepository(id int) (Book, error)
	UpdateBookByIdRepository(id int, book Book) (Book, error)
	DeleteBookByIdRepository(id int) (Book, error)
}

type bookRepository struct{}

func NewRepository() Repository {
	return &bookRepository{}
}

func (repository *bookRepository) CreateBookRepository(book Book) (Book, error) {
	query := `
		INSERT INTO books
		(
			title, 
			description, 
			image_url, 
			release_year, 
			price, 
			total_page, 
			thickness,
			category_id,
			created_by, 
			modified_by
		)
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING *
	`

	err := database.DB.QueryRow(query, book.Title, book.Description, book.Image_Url, book.Release_Year, book.Price, book.Total_Page, book.Thickness, book.Category_Id, book.Created_By, book.Modified_By).
		Scan(&book.Id, &book.Title, &book.Description, &book.Image_Url, &book.Release_Year, &book.Price, &book.Total_Page, &book.Thickness, &book.Category_Id, &book.Created_At, &book.Created_By, &book.Modified_At, &book.Modified_By)

	if err != nil {
		return Book{}, err
	}

	return book, err
}

func (repository *bookRepository) GetAllBookRepository() ([]Book, error) {
	var books []Book

	query := "SELECT * FROM books"

	rows, err := database.DB.Query(query)

	if err != nil {
		return []Book{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var book Book

		err = rows.Scan(&book.Id, &book.Title, &book.Description, &book.Image_Url, &book.Release_Year, &book.Price, &book.Total_Page, &book.Thickness, &book.Category_Id, &book.Created_At, &book.Created_By, &book.Modified_At, &book.Modified_By)

		if err != nil {
			return []Book{}, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (repository *bookRepository) GetBookByIdRepository(id int) (Book, error) {
	var book Book

	query := "SELECT * FROM books WHERE id = $1"

	err := database.DB.QueryRow(query, id).
		Scan(&book.Id, &book.Title, &book.Description, &book.Image_Url, &book.Release_Year, &book.Price, &book.Total_Page, &book.Thickness, &book.Category_Id, &book.Created_At, &book.Created_By, &book.Modified_At, &book.Modified_By)

	if err != nil {
		if err == sql.ErrNoRows {
			return Book{}, fmt.Errorf("failed to get book data, book with id \"%d\" not found", id)
		}

		return Book{}, err
	}

	return book, nil
}

func (repository *bookRepository) UpdateBookByIdRepository(id int, book Book) (Book, error) {
	query := `
		UPDATE books 
		SET
			title = $2, 
			description = $3, 
			image_url = $4, 
			release_year = $5, 
			price = $6, 
			total_page = $7, 
			category_id = $8,
			modified_by = $9 
		WHERE id = $1 
		RETURNING *
	`

	err := database.DB.QueryRow(query, id, book.Title, book.Description, book.Image_Url, book.Release_Year, book.Price, book.Total_Page, book.Category_Id, book.Modified_By).
		Scan(&book.Id, &book.Title, &book.Description, &book.Image_Url, &book.Release_Year, &book.Price, &book.Total_Page, &book.Thickness, &book.Category_Id, &book.Created_At, &book.Created_By, &book.Modified_At, &book.Modified_By)

	if err != nil {
		if err == sql.ErrNoRows {
			return book, fmt.Errorf("failed updating book, book with id \"%d\" not found", id)
		}

		return Book{}, err
	}

	return book, nil
}

func (repository *bookRepository) DeleteBookByIdRepository(id int) (Book, error) {
	var deletedBook Book

	query := "DELETE FROM books WHERE id = $1 RETURNING *"

	err := database.DB.QueryRow(query, id).
		Scan(&deletedBook.Id, &deletedBook.Title, &deletedBook.Description, &deletedBook.Image_Url, &deletedBook.Release_Year, &deletedBook.Price, &deletedBook.Total_Page, &deletedBook.Thickness, &deletedBook.Category_Id, &deletedBook.Created_At, &deletedBook.Created_By, &deletedBook.Modified_At, &deletedBook.Modified_By)

	if err != nil {
		if err == sql.ErrNoRows {
			return deletedBook, fmt.Errorf("failed deleting book, book with id \"%d\" not found", id)
		}

		return Book{}, err
	}

	return deletedBook, nil
}
