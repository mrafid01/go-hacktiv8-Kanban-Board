package service

import (
	"project3/model/entity"
	"project3/model/input"
	"project3/repository"
)

type CategoryService interface {
	CreateCategory(input input.CategoryCreateInput) (entity.Category, error)
	GetCategoryByID(id_category int) (entity.Category, error)
	PatchCategory(id_category int, input input.CategoryPatchInput) (entity.Category, error)
	DeleteCategory(id_category int) (entity.Category, error)
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) *categoryService {
	return &categoryService{categoryRepository}
}

func (s *categoryService) CreateCategory(input input.CategoryCreateInput) (entity.Category, error) {}
func (s *categoryService) GetCategoryByID(id_category int) (entity.Category, error)                {}
func (s *categoryService) PatchCategory(id_category int, input input.CategoryPatchInput) (entity.Category, error) {
}
func (s *categoryService) DeleteCategory(id_category int) (entity.Category, error) {}
