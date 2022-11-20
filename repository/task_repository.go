package repository

import (
	"project3/model/entity"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task entity.Task) (entity.Task, error)
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

func (r *taskRepository) Create(task entity.Task) (entity.Task, error) { return entity.Task{}, nil }

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
