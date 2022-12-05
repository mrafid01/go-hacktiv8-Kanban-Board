package controller

import (
	"net/http"
	"project3/helper"
	"project3/model/input"
	"project3/model/response"
	"project3/service"

	"github.com/gin-gonic/gin"
)

type categoryController struct {
	categoryService service.CategoryService
	taskService     service.TaskService
}

func NewCategoryController(categoryService service.CategoryService, taskService service.TaskService) *categoryController {
	return &categoryController{categoryService, taskService}
}

func (h *categoryController) CreateCategory(c *gin.Context) {
	var (
		input input.CategoryCreateInput
	)

	role_user := c.MustGet("roleUser").(string)

	err := c.ShouldBindJSON(&input)
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

	categoryData, err := h.categoryService.CreateCategory(role_user, input)
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

	categoryResponse := response.CategoryCreateResponse{
		ID:        categoryData.ID,
		Type:      categoryData.Type,
		CreatedAt: categoryData.CreatedAt,
	}

	c.JSON(
		http.StatusCreated,
		helper.NewResponse(
			http.StatusCreated,
			"created",
			categoryResponse,
		),
	)
}

func (h *categoryController) GetCategory(c *gin.Context) {
	var (
		allTasks      []response.CategoryTask
		allCategories []response.CategoryGetResponse
	)

	categoryData, err := h.categoryService.GetAllCategory()
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

	for _, dataCategory := range categoryData {
		tasksData, err := h.taskService.GetTasksByCategoryID(dataCategory.ID)
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
		for _, dataTask := range tasksData {
			allTasksTmp := response.CategoryTask{
				ID:          dataTask.ID,
				Title:       dataTask.Title,
				Description: dataTask.Description,
				UserID:      dataTask.UserID,
				CategoryID:  dataTask.CategoryID,
				CreatedAt:   dataTask.CreatedAt,
				UpdatedAt:   dataTask.UpdatedAt,
			}
			allTasks = append(allTasks, allTasksTmp)
		}
		allCategoriesTmp := response.CategoryGetResponse{
			ID:        dataCategory.ID,
			Type:      dataCategory.Type,
			UpdatedAt: dataCategory.UpdatedAt,
			CreatedAt: dataCategory.CreatedAt,
			Tasks:     allTasks,
		}

		allCategories = append(allCategories, allCategoriesTmp)
		allTasks = []response.CategoryTask{}
	}

	c.JSON(
		http.StatusOK,
		helper.NewResponse(
			http.StatusOK,
			"ok",
			allCategories,
		),
	)
}

func (h *categoryController) PatchCategory(c *gin.Context) {
	var (
		inputBody input.CategoryPatchInput
		uri       input.CategoryIdUri
	)

	role_user := c.MustGet("roleUser").(string)

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

	err = c.ShouldBindUri(&uri)
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

	categoryData, err := h.categoryService.PatchCategory(role_user, uri.ID, inputBody)
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

	categoryResponse := response.CategoryPatchResponse{
		ID:        categoryData.ID,
		Type:      categoryData.Type,
		UpdatedAt: categoryData.UpdatedAt,
	}

	c.JSON(
		http.StatusOK,
		helper.NewResponse(
			http.StatusOK,
			"ok",
			categoryResponse,
		),
	)
}

func (h *categoryController) DeleteCategory(c *gin.Context) {
	var (
		uri input.CategoryIdUri
	)

	role_user := c.MustGet("roleUser").(string)

	err := c.ShouldBindUri(&uri)
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

	_, err = h.categoryService.DeleteCategory(role_user, uri.ID)
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

	categoryResponse := response.CategoryDeleteResponse{
		Message: "Category has been successfully deleted",
	}

	c.JSON(
		http.StatusOK,
		helper.NewResponse(
			http.StatusOK,
			"ok",
			categoryResponse,
		),
	)
}
