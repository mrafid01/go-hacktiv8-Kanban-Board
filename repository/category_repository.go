package repository

import (
	"project3/model/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(category entity.Category) (entity.Category, error)
	GetAll() ([]entity.Category, error)
	GetTasksByCategoryID(id_category int) ([]entity.Task, error)
	PatchType(id_category int, category entity.Category) (entity.Category, error)
	GetCategoryById(id_category int) (entity.Category, error)
	Delete(id_category int) (entity.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) CreateCategory(category entity.Category) (entity.Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *categoryRepository) GetAll() ([]entity.Category, error) {
	var categories []entity.Category

	err := r.db.Find(&categories).Error
	if err != nil {
		return []entity.Category{}, err
	}

	return categories, nil
}

func (r *categoryRepository) GetTasksByCategoryID(id_category int) ([]entity.Task, error) {
	var tasks []entity.Task

	err := r.db.Where("category_id = ?", id_category).Find(&tasks).Error
	if err != nil {
		return []entity.Task{}, err
	}

	return tasks, nil
}

func (r *categoryRepository) PatchType(id_category int, category entity.Category) (entity.Category, error) {
	err := r.db.Where("id = ?", id_category).Updates(&category).Error
	if err != nil {
		return entity.Category{}, err
	}

	return category, nil
}

func (r *categoryRepository) GetCategoryById(id_category int) (entity.Category, error) {
	var category entity.Category

	err := r.db.Where("id = ?", id_category).Find(&category).Error
	if err != nil {
		return entity.Category{}, err
	}

	return category, nil
}

func (r *categoryRepository) Delete(id_category int) (entity.Category, error) {
	var category entity.Category

	err := r.db.Where("id = ?", id_category).Delete(&category).Error
	if err != nil {
		return entity.Category{}, err
	}

	return category, nil
}
