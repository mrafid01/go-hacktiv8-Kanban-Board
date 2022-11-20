package repository

import (
	"project3/model/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(category entity.Category) (entity.Category, error)
	FindByID(id_category int) (entity.Category, error)
	Update(id_category int, category entity.Category) (entity.Category, error)
	Delete(id_category int) (entity.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) CreateCategory(category entity.Category) (entity.Category, error) {
	return entity.Category{}, nil
}

func (r *categoryRepository) FindByID(id_category int) (entity.Category, error) {
	return entity.Category{}, nil
}

func (r *categoryRepository) Update(id_category int, category entity.Category) (entity.Category, error) {
	return entity.Category{}, nil
}

func (r *categoryRepository) Delete(id_category int) (entity.Category, error) {
	return entity.Category{}, nil
}
