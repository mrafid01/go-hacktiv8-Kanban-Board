package service

import (
	"errors"
	"project3/model/entity"
	"project3/model/input"
	"project3/repository"
)

type TaskService interface {
	CreateTask(id_user int, input input.TaskCreateInput) (entity.Task, error)
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

func (s *taskService) CreateTask(id_user int, input input.TaskCreateInput) (entity.Task, error) {
	categoryData, err := s.taskRepository.FindCategoryByCategoryId(input.CategoryID)
	if err != nil {
		return entity.Task{}, err
	}
	if categoryData.ID == 0 {
		return entity.Task{}, errors.New("data not found")
	}

	task := entity.Task{
		Title:       input.Title,
		Description: input.Description,
		CategoryID:  input.CategoryID,
		UserID:      id_user,
		Status:      false,
	}

	TaskData, err := s.taskRepository.Create(task)
	if err != nil {
		return entity.Task{}, err
	}

	return TaskData, nil
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
