package service

import (
	"project3/model/entity"
	"project3/model/input"
	"project3/repository"
)

type TaskService interface {
	CreateTask(input input.TaskCreateInput) (entity.Task, error)
	GetTaskByID(id_task int) (entity.Task, error)
	UpdateTask(id_task int, input input.TaskUpdateInput) (entity.Task, error)
	PatchStatusTask(id_task int, input input.TaskPatchStatusInput) (entity.Task, error)
	PatchCategoryTask(id_task int, input input.TaskPatchCategoryInput) (entity.Task, error)
	DeleteTask(id_task int) (entity.Task, error)
}

type taskService struct {
	taskRepository repository.TaskRepository
}

func NewTaskService(taskRepository repository.TaskRepository) *taskService {
	return &taskService{taskRepository}
}

func (s *taskService) CreateTask(input input.TaskCreateInput) (entity.Task, error) {
	return entity.Task{}, nil
}
func (s *taskService) GetTaskByID(id_task int) (entity.Task, error) { return entity.Task{}, nil }
func (s *taskService) UpdateTask(id_task int, input input.TaskUpdateInput) (entity.Task, error) {
	return entity.Task{}, nil
}
func (s *taskService) PatchStatusTask(id_task int, input input.TaskPatchStatusInput) (entity.Task, error) {
	return entity.Task{}, nil
}
func (s *taskService) PatchCategoryTask(id_task int, input input.TaskPatchCategoryInput) (entity.Task, error) {
	return entity.Task{}, nil
}
func (s *taskService) DeleteTask(id_task int) (entity.Task, error) { return entity.Task{}, nil }
