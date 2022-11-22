package controller

import (
	"net/http"
	"project3/helper"
	"project3/model/input"
	"project3/model/response"
	"project3/service"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService}
}

func (h *userController) RegisterUser(c *gin.Context) {
	var (
		input input.UserRegisterInput
	)

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

	userData, err := h.userService.CreateUser(input)
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

	userResponse := response.UserRegisterResponse{
		ID:        userData.ID,
		FullName:  userData.FullName,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt,
	}

	c.JSON(
		http.StatusCreated,
		helper.NewResponse(
			http.StatusCreated,
			"created",
			userResponse,
		),
	)
}

func (h *userController) LoginUser(c *gin.Context) {
	var (
		input input.UserLoginInput
	)

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

	token, err := h.userService.LoginUser(input)
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

	userResponse := response.UserLoginResponse{
		Token: token,
	}

	c.JSON(
		http.StatusOK,
		helper.NewResponse(
			http.StatusOK,
			"ok",
			userResponse,
		),
	)
}

func (h *userController) UpdateUser(c *gin.Context) {
	var (
		input input.UserUpdateInput
	)

	id_user := c.MustGet("currentUser").(int)

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

	userData, err := h.userService.UpdateUser(id_user, input)
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

	userResponse := response.UserUpdateResponse{
		ID:        userData.ID,
		FullName:  userData.FullName,
		Email:     userData.Email,
		UpdatedAt: userData.UpdatedAt,
	}

	c.JSON(
		http.StatusOK,
		helper.NewResponse(
			http.StatusOK,
			"ok",
			userResponse,
		),
	)
}

func (h *userController) DeleteUser(c *gin.Context) {
	id_user := c.MustGet("currentUser").(int)

	_, err := h.userService.DeleteUser(id_user)
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

	userResponse := response.UserDeleteResponse{
		Message: "Your account has been successfully deleted",
	}

	c.JSON(
		http.StatusOK,
		helper.NewResponse(
			http.StatusOK,
			"ok",
			userResponse,
		),
	)
}

func (h *userController) RegisterAdmin(c *gin.Context) {
	var (
		input input.UserRegisterInput
	)

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

	userData, err := h.userService.CreateAdmin(input)
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

	userResponse := response.UserRegisterResponse{
		ID:        userData.ID,
		FullName:  userData.FullName,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt,
	}

	c.JSON(
		http.StatusCreated,
		helper.NewResponse(
			http.StatusCreated,
			"created",
			userResponse,
		),
	)
}
