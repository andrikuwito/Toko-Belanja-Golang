package service

import (
	"errors"
	"tokobelanja-golang/model/entity"
	"tokobelanja-golang/model/input"
	"tokobelanja-golang/repository"
)

type CategoryService interface {
	CreateCategory(input input.CreateCategory) (entity.Category, error)
	GetCategoryByID(ID int) (entity.Category, error)
	UpdateCategory(ID int, input input.UpdateCategory) (entity.Category, error)
	DeleteCategory(ID int) (entity.Category, error)
	GetCategory() ([]entity.Category, error)
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) *categoryService {
	return &categoryService{categoryRepository}
}

func (s *categoryService) CreateCategory(input input.CreateCategory) (entity.Category, error) {
	newCategory := entity.Category{}
	newCategory.Type = input.Type

	created, err := s.categoryRepository.Save(newCategory)

	if err != nil {
		return entity.Category{}, err
	}

	return created, nil
}

func (s *categoryService) GetCategoryByID(ID int) (entity.Category, error) {
	category, err := s.categoryRepository.GetByID(ID)

	if err != nil {
		return category, err
	}

	return category, nil
}

func (s *categoryService) GetCategory() ([]entity.Category, error) {
	category, err := s.categoryRepository.GetCategory()

	if err != nil {
		return category, err
	}

	return category, nil
}

func (s *categoryService) UpdateCategory(ID int, input input.UpdateCategory) (entity.Category, error) {
	Result, err := s.categoryRepository.GetByID(ID)

	if err != nil {
		return entity.Category{}, err
	}

	if Result.ID == 0 {
		return entity.Category{}, errors.New("not found!")
	}

	updated := entity.Category{
		Type: input.Type,
	}

	categoryUpdate, err := s.categoryRepository.Update(ID, updated)

	if err != nil {
		return categoryUpdate, err
	}

	return categoryUpdate, nil
}

func (s *categoryService) DeleteCategory(ID int) (entity.Category, error) {
	category, err := s.categoryRepository.GetByID(ID)

	if err != nil {
		return entity.Category{}, err
	}

	if category.ID == 0 {
		return entity.Category{}, nil
	}

	Deleted, err := s.categoryRepository.Delete(ID)

	if err != nil {
		return entity.Category{}, err
	}

	return Deleted, nil
}
