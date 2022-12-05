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
	GetTasksByCategoryID(id_category int) ([]entity.Task, error)
	UpdateTask(id_user int, id_task int, input input.TaskUpdateInput) (entity.Task, error)
	PatchStatusTask(id_user int, id_task int, input input.TaskPatchStatusInput) (entity.Task, error)
	PatchCategoryTask(id_user int, id_task int, input input.TaskPatchCategoryInput) (entity.Task, error)
	DeleteTask(id_user int, id_task int) (entity.Task, error)
}

type taskService struct {
	taskRepository     repository.TaskRepository
	categoryRepository repository.CategoryRepository
}

func NewTaskService(taskRepository repository.TaskRepository, categoryRepository repository.CategoryRepository) *taskService {
	return &taskService{taskRepository, categoryRepository}
}

func (s *taskService) CreateTask(id_user int, input input.TaskCreateInput) (entity.Task, error) {
	categoryData, err := s.categoryRepository.GetCategoryById(input.CategoryID)
	if err != nil {
		return entity.Task{}, err
	}
	if categoryData.ID == 0 {
		return entity.Task{}, errors.New("category not found")
	}

	task := entity.Task{
		Title:       input.Title,
		Description: input.Description,
		CategoryID:  input.CategoryID,
		UserID:      id_user,
		Status:      newBool(false),
	}

	taskData, err := s.taskRepository.Create(task)
	if err != nil {
		return entity.Task{}, err
	}

	return taskData, nil
}

func (s *taskService) GetAllTask() ([]entity.Task, error) {
	tasksData, err := s.taskRepository.FindAll()
	if err != nil {
		return []entity.Task{}, err
	}

	return tasksData, nil
}

func (s *taskService) GetTasksByCategoryID(id_category int) ([]entity.Task, error) {
	return s.taskRepository.FindByCategoryID(id_category)
}

func (s *taskService) UpdateTask(id_user int, id_task int, input input.TaskUpdateInput) (entity.Task, error) {
	taskData, err := s.taskRepository.FindByID(id_task)
	if err != nil {
		return entity.Task{}, err
	}
	if taskData.ID == 0 {
		return entity.Task{}, errors.New("task not found")
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

func (s *taskService) PatchStatusTask(id_user int, id_task int, input input.TaskPatchStatusInput) (entity.Task, error) {
	taskData, err := s.taskRepository.FindByID(id_task)
	if err != nil {
		return entity.Task{}, err
	}
	if taskData.ID == 0 {
		return entity.Task{}, errors.New("task not found")
	}
	if taskData.UserID != id_user {
		return entity.Task{}, errors.New("can't update other people task")
	}

	task := entity.Task{
		Status: input.Status,
	}

	_, err = s.taskRepository.Update(id_task, task)
	if err != nil {
		return entity.Task{}, err
	}

	return s.taskRepository.FindByID(id_task)
}

func (s *taskService) PatchCategoryTask(id_user int, id_task int, input input.TaskPatchCategoryInput) (entity.Task, error) {
	taskData, err := s.taskRepository.FindByID(id_task)
	if err != nil {
		return entity.Task{}, err
	}
	if taskData.ID == 0 {
		return entity.Task{}, errors.New("task not found")
	}
	if taskData.UserID != id_user {
		return entity.Task{}, errors.New("can't update other people task")
	}

	categoryData, err := s.categoryRepository.GetCategoryById(input.CategoryID)
	if err != nil {
		return entity.Task{}, err
	}
	if categoryData.ID == 0 {
		return entity.Task{}, errors.New("category not found")
	}

	task := entity.Task{
		CategoryID: input.CategoryID,
	}

	_, err = s.taskRepository.Update(id_task, task)
	if err != nil {
		return entity.Task{}, err
	}

	return s.taskRepository.FindByID(id_task)
}

func (s *taskService) DeleteTask(id_user int, id_task int) (entity.Task, error) {
	taskData, err := s.taskRepository.FindByID(id_task)
	if err != nil {
		return entity.Task{}, err
	}
	if taskData.ID == 0 {
		return entity.Task{}, errors.New("task not found")
	}
	if taskData.UserID != id_user {
		return entity.Task{}, errors.New("can't delete other people task")
	}

	return s.taskRepository.Delete(id_task)
}

func newBool(b bool) *bool {
	return &b
}
