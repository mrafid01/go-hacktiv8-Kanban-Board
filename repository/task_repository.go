package repository

import (
	"project3/model/entity"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task entity.Task) (entity.Task, error)
	FindCategoryByCategoryId(id_category int) (entity.Category, error)
	FindAll() ([]entity.Task, error)
	FindByID(id_task int) (entity.Task, error)
	Update(id_task int, task entity.Task) (entity.Task, error)
	PatchStatus(id_task int, task entity.Task) (entity.Task, error)
	PatchCategory(id_task int, task entity.Task) (entity.Task, error)
	Delete(id_task int) (entity.Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) Create(task entity.Task) (entity.Task, error) {
	err := r.db.Preload("User").Preload("Category").Create(&task).Error
	if err != nil {
		return task, err
	}

	return task, nil
}
func (r *taskRepository) FindCategoryByCategoryId(id_category int) (entity.Category, error) {
	var category entity.Category

	err := r.db.Where("id = ?", id_category).Find(&category).Error
	if err != nil {
		return entity.Category{}, err
	}

	return category, nil
}

func (r *taskRepository) FindAll() ([]entity.Task, error) {
	var tasks []entity.Task

	err := r.db.Preload("User").Preload("Category").Find(&tasks).Error
	if err != nil {
		return []entity.Task{}, err
	}

	return tasks, nil
}

func (r *taskRepository) FindByID(id_task int) (entity.Task, error) { return entity.Task{}, nil }

func (r *taskRepository) Update(id_task int, task entity.Task) (entity.Task, error) {
	return entity.Task{}, nil
}

func (r *taskRepository) PatchStatus(id_task int, task entity.Task) (entity.Task, error) {
	return entity.Task{}, nil
}

func (r *taskRepository) PatchCategory(id_task int, task entity.Task) (entity.Task, error) {
	return entity.Task{}, nil
}

func (r *taskRepository) Delete(id_task int) (entity.Task, error) { return entity.Task{}, nil }
