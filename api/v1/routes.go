package v1

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-todo-poc/internal/todo"
	"go-fiber-todo-poc/pkg/database"
)

func SetupRoutes(fiberApp *fiber.App) {
	api := fiberApp.Group("/api/v1")

	// Use the initialized DB instance from the database package
	todoService := todo.NewTodoService(database.DB)
	todo.SetupTodoRoutes(api, todoService)
}
