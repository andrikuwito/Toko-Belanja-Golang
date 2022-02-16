package repository

import (
	"tokobelanja-golang/model/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Save(input entity.Category) (entity.Category, error)
	CheckType(tipe string) (entity.Category, error)
	GetByID(ID int) (entity.Category, error)
	Update(ID int, category entity.Category) (entity.Category, error)
	Delete(ID int) (entity.Category, error)
	GetCategory() ([]entity.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) Save(category entity.Category) (entity.Category, error) {
	err := r.db.Create(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *categoryRepository) CheckType(tipe string) (entity.Category, error) {
	typeSame := entity.Category{}

	err := r.db.Where("type = ?", tipe).Find(&typeSame).Error

	if err != nil {
		return entity.Category{}, err
	}

	return typeSame, nil
}

func (r *categoryRepository) GetByID(ID int) (entity.Category, error) {
	categoryResult := entity.Category{}

	err := r.db.Where("id = ?", ID).Find(&categoryResult).Error

	if err != nil {
		return entity.Category{}, err
	}

	return categoryResult, nil
}

func (r *categoryRepository) GetCategory() ([]entity.Category, error) {
	categoryResult := []entity.Category{}

	err := r.db.Find(&categoryResult).Error

	if err != nil {
		return categoryResult, err
	}

	return categoryResult, nil
}

func (r *categoryRepository) Update(ID int, category entity.Category) (entity.Category, error) {
	err := r.db.Where("id = ?", ID).Updates(&category).Error

	if err != nil {
		return entity.Category{}, err
	}

	return category, nil
}

func (r *categoryRepository) Delete(ID int) (entity.Category, error) {

	Deleted := entity.Category{
		ID: ID,
	}

	err := r.db.Where("id = ?", ID).Delete(&Deleted).Error

	if err != nil {
		return entity.Category{}, err
	}

	return Deleted, err
}
