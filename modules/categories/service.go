package categories

import (
	"strconv"
)

type Service interface {
	CreateCategoryService(category Category) (Category, error)
	GetAllCategoryService() ([]Category, error)
	GetCategoryByIdService(categoryId string) (Category, error)
	UpdateCategoryByIdService(categoryId string, category Category) (Category, error)
	DeleteCategoryByIdService(categoryId string) (Category, error)
}

type categoryService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &categoryService{
		repository,
	}
}

func (service *categoryService) CreateCategoryService(category Category) (Category, error) {
	createdCategory, err := service.repository.CreateCategoryRepository(category)

	if err != nil {
		return Category{}, err
	}

	return createdCategory, nil
}

func (service *categoryService) GetAllCategoryService() ([]Category, error) {
	category, err := service.repository.GetAllCategoryRepository()

	if err != nil {
		return []Category{}, err
	}

	return category, nil
}

func (service *categoryService) GetCategoryByIdService(categoryId string) (Category, error) {
	id, err := strconv.Atoi(categoryId)

	if err != nil {
		return Category{}, err
	}

	category, err := service.repository.GetCategoryByIdRepository(id)

	if err != nil {
		return Category{}, err
	}

	return category, nil
}

func (service *categoryService) UpdateCategoryByIdService(categoryId string, category Category) (Category, error) {
	id, err := strconv.Atoi(categoryId)

	if err != nil {
		return Category{}, err
	}

	updatedCategory, err := service.repository.UpdateCategoryByIdRepository(id, category)

	if err != nil {
		return Category{}, err
	}

	return updatedCategory, err
}

func (service *categoryService) DeleteCategoryByIdService(categoryId string) (Category, error) {
	id, err := strconv.Atoi(categoryId)

	if err != nil {
		return Category{}, err
	}

	deletedCategory, err := service.repository.DeleteCategoryByIdRepository(id)

	if err != nil {
		return Category{}, err
	}

	return deletedCategory, err
}
