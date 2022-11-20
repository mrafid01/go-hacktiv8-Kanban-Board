package controller

import (
	"project3/service"

	"github.com/gin-gonic/gin"
)

type taskController struct {
	taskService service.TaskService
}

func NewTaskController(taskService service.TaskService) *taskController {
	return &taskController{taskService}
}

func (h *taskController) CreateTask(c *gin.Context)        {}
func (h *taskController) GetTask(c *gin.Context)           {}
func (h *taskController) UpdateTask(c *gin.Context)        {}
func (h *taskController) PatchStatusTask(c *gin.Context)   {}
func (h *taskController) PatchCategoryTask(c *gin.Context) {}
func (h *taskController) DeleteTask(c *gin.Context)        {}
