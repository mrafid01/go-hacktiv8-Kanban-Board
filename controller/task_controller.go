package controller

import (
	"net/http"
	"project3/helper"
	"project3/model/input"
	"project3/model/response"
	"project3/service"

	"github.com/gin-gonic/gin"
)

type taskController struct {
	taskService service.TaskService
}

func NewTaskController(taskService service.TaskService) *taskController {
	return &taskController{taskService}
}

func (h *taskController) CreateTask(c *gin.Context) {
	var (
		inputBody input.TaskCreateInput
	)

	id_user := c.MustGet("currentUser").(int)

	err := c.ShouldBindJSON(&inputBody)
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	taskData, err := h.taskService.CreateTask(id_user, inputBody)
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	taskResponse := response.TaskCreateResponse{
		ID:          taskData.ID,
		Title:       taskData.Title,
		Status:      taskData.Status,
		Description: taskData.Description,
		UserID:      taskData.UserID,
		CategoryID:  taskData.CategoryID,
		CreatedAt:   taskData.CreatedAt,
	}

	c.JSON(
		http.StatusCreated,
		helper.NewResponse(
			http.StatusCreated,
			"created",
			taskResponse,
		),
	)

}
func (h *taskController) GetTask(c *gin.Context)           {}
func (h *taskController) UpdateTask(c *gin.Context)        {}
func (h *taskController) PatchStatusTask(c *gin.Context)   {}
func (h *taskController) PatchCategoryTask(c *gin.Context) {}
func (h *taskController) DeleteTask(c *gin.Context)        {}
