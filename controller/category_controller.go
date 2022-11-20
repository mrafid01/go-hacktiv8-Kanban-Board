package controller

import (
	"project3/service"

	"github.com/gin-gonic/gin"
)

type categoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *categoryController {
	return &categoryController{categoryService}
}

func (h *categoryController) CreateCategory(c *gin.Context) {}
func (h *categoryController) GetCategory(c *gin.Context)    {}
func (h *categoryController) PatchCategory(c *gin.Context)  {}
func (h *categoryController) DeleteCategory(c *gin.Context) {}
