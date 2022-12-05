package main

import (
	"os"
	"project3/config"
	"project3/controller"
	"project3/middleware"
	"project3/repository"
	"project3/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	dbUsername := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")

	db := config.InitDB(dbUsername, dbPassword, dbHost, dbPort, dbName)

	userRepository := repository.NewUserRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)
	taskRepository := repository.NewTaskRepository(db)

	userService := service.NewUserService(userRepository)
	categoryService := service.NewCategoryService(categoryRepository, taskRepository)
	taskService := service.NewTaskService(taskRepository, categoryRepository)

	userController := controller.NewUserController(userService)
	categoryController := controller.NewCategoryController(categoryService, taskService)
	taskController := controller.NewTaskController(taskService)

	router := gin.Default()

	// User
	router.POST("/users/register", userController.RegisterUser)
	router.POST("/users/login", userController.LoginUser)
	router.PUT("/users/update-account", middleware.AuthMiddleware, userController.UpdateUser)
	router.DELETE("/users/delete-account", middleware.AuthMiddleware, userController.DeleteUser)
	// Create Admin
	router.POST("/users/admin", userController.RegisterAdmin)

	// Category
	router.POST("/categories", middleware.AuthMiddleware, categoryController.CreateCategory)
	router.GET("/categories", categoryController.GetCategory)
	router.PATCH("/categories/:categoryID", middleware.AuthMiddleware, categoryController.PatchCategory)
	router.DELETE("/categories/:categoryID", middleware.AuthMiddleware, categoryController.DeleteCategory)

	// Task
	router.POST("/tasks", middleware.AuthMiddleware, taskController.CreateTask)
	router.GET("/tasks", middleware.AuthMiddleware, taskController.GetTask)
	router.PUT("/tasks/:taskID", middleware.AuthMiddleware, taskController.UpdateTask)
	router.PATCH("/tasks/update-status/:taskID", middleware.AuthMiddleware, taskController.PatchStatusTask)
	router.PATCH("/tasks/update-category/:taskID", middleware.AuthMiddleware, taskController.PatchCategoryTask)
	router.DELETE("/tasks/:taskID", middleware.AuthMiddleware, taskController.DeleteTask)

	router.Run(":" + port)
}
