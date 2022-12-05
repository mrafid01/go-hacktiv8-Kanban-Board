package repository

import (
	"project3/model/entity"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task entity.Task) (entity.Task, error)
	FindAll() ([]entity.Task, error)
	FindByID(id_task int) (entity.Task, error)
	FindByCategoryID(id_category int) ([]entity.Task, error)
	Update(id_task int, task entity.Task) (entity.Task, error)
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

func (r *taskRepository) FindAll() ([]entity.Task, error) {
	var tasks []entity.Task

	err := r.db.Preload("User").Preload("Category").Find(&tasks).Error
	if err != nil {
		return []entity.Task{}, err
	}

	return tasks, nil
}

func (r *taskRepository) FindByID(id_task int) (entity.Task, error) {
	var task entity.Task

	err := r.db.Preload("User").Preload("Category").Where("id = ?", id_task).Find(&task).Error
	if err != nil {
		return entity.Task{}, err
	}

	return task, nil
}

func (r *taskRepository) FindByCategoryID(id_category int) ([]entity.Task, error) {
	var tasks []entity.Task

	err := r.db.Where("category_id = ?", id_category).Find(&tasks).Error
	if err != nil {
		return []entity.Task{}, err
	}

	return tasks, nil
}

func (r *taskRepository) Update(id_task int, task entity.Task) (entity.Task, error) {
	err := r.db.Preload("User").Preload("Category").Where("id = ?", id_task).Updates(&task).Error
	if err != nil {
		return entity.Task{}, err
	}

	return task, nil
}

func (r *taskRepository) Delete(id_task int) (entity.Task, error) {
	var task entity.Task

	err := r.db.Where("id = ?", id_task).Delete(&task).Error
	if err != nil {
		return entity.Task{}, err
	}

	return task, nil
}
