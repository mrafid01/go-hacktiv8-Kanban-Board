package service

import (
	"errors"
	"project3/model/entity"
	"project3/model/input"
	"project3/repository"
)

type TaskService interface {
	CreateTask(id_user int, input input.TaskCreateInput) (entity.Task, error)
	GetAllTask() ([]entity.Task, error)
	UpdateTask(id_user int, id_task int, input input.TaskUpdateInput) (entity.Task, error)
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

func (s *taskService) GetAllTask() ([]entity.Task, error) {
	tasks, err := s.taskRepository.FindAll()
	if err != nil {
		return []entity.Task{}, err
	}

	return tasks, nil
}

func (s *taskService) UpdateTask(id_user int, id_task int, input input.TaskUpdateInput) (entity.Task, error) {
	taskData, err := s.taskRepository.FindByID(id_task)
	if err != nil {
		return entity.Task{}, err
	}
	if taskData.ID == 0 {
		return entity.Task{}, errors.New("data not found")
	}
	if taskData.UserID != id_user {
		return entity.Task{}, errors.New("can't update other people task")
	}

	task := entity.Task{
		Title:       input.Title,
		Description: input.Description,
	}

	_, err = s.taskRepository.Update(id_task, task)
	if err != nil {
		return entity.Task{}, err
	}

	return s.taskRepository.FindByID(id_task)
}

func (s *taskService) PatchStatusTask(id_task int, input input.TaskPatchStatusInput) (entity.Task, error) {
	return entity.Task{}, nil
}
func (s *taskService) PatchCategoryTask(id_task int, input input.TaskPatchCategoryInput) (entity.Task, error) {
	return entity.Task{}, nil
}
func (s *taskService) DeleteTask(id_task int) (entity.Task, error) { return entity.Task{}, nil }
