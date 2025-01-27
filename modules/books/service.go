package books

import (
	"strconv"
)

type Service interface {
	CreateBookService(book Book) (Book, error)
	GetAllBookService() ([]Book, error)
	GetBookByIdService(bookId string) (Book, error)
	UpdateBookByIdService(bookId string, book Book) (Book, error)
	DeleteBookByIdService(bookId string) (Book, error)
}

type bookService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &bookService{
		repository,
	}
}

func (service *bookService) CreateBookService(book Book) (Book, error) {
	validationError := ValidateReleaseYear(book.Release_Year)

	if validationError != nil {
		return Book{}, validationError
	}

	book.Thickness = DefineBookThickness(book.Total_Page)

	createdBook, err := service.repository.CreateBookRepository(book)

	if err != nil {
		return Book{}, err
	}

	return createdBook, nil
}

func (service *bookService) GetAllBookService() ([]Book, error) {
	book, err := service.repository.GetAllBookRepository()

	if err != nil {
		return []Book{}, err
	}

	return book, nil
}

func (service *bookService) GetBookByIdService(bookId string) (Book, error) {
	id, err := strconv.Atoi(bookId)

	if err != nil {
		return Book{}, err
	}

	book, err := service.repository.GetBookByIdRepository(id)

	if err != nil {
		return Book{}, err
	}

	return book, nil
}

func (service *bookService) UpdateBookByIdService(bookId string, book Book) (Book, error) {
	id, err := strconv.Atoi(bookId)

	if err != nil {
		return Book{}, err
	}

	validationError := ValidateReleaseYear(book.Release_Year)

	if validationError != nil {
		return Book{}, validationError
	}

	book.Thickness = DefineBookThickness(book.Total_Page)

	updatedBook, err := service.repository.UpdateBookByIdRepository(id, book)

	if err != nil {
		return Book{}, err
	}

	return updatedBook, err
}

func (service *bookService) DeleteBookByIdService(bookId string) (Book, error) {
	id, err := strconv.Atoi(bookId)

	if err != nil {
		return Book{}, err
	}

	deletedBook, err := service.repository.DeleteBookByIdRepository(id)

	if err != nil {
		return Book{}, err
	}

	return deletedBook, err
}
