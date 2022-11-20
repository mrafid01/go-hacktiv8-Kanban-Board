package controller

import (
	"project3/service"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService}
}

func (h *userController) RegisterUser(c *gin.Context) {}
func (h *userController) LoginUser(c *gin.Context)    {}
func (h *userController) UpdateUser(c *gin.Context)   {}
func (h *userController) DeleteUser(c *gin.Context)   {}
