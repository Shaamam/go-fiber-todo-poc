package todo

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-todo-poc/pkg/models"
)

type TodoHandler struct {
	service *TodoService
}

// NewTodoHandler creates a new handler with the given service.
func NewTodoHandler(service *TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

// GetTodos handles the GET request to retrieve all todos.
func (h *TodoHandler) GetTodos(c *fiber.Ctx) error {
	todos, err := h.service.GetTodos()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(todos)
}

// CreateTodo handles the POST request to create a new todo.
func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	if err := h.service.CreateTodo(todo); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(todo)
}

// UpdateTodo handles the PUT request to update an existing todo.
func (h *TodoHandler) UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	updatedTodo := new(models.Todo)
	if err := c.BodyParser(updatedTodo); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	if err := h.service.UpdateTodo(id, updatedTodo); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(updatedTodo)
}

// DeleteTodo handles the DELETE request to remove a todo.
func (h *TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.DeleteTodo(id); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendStatus(204)
}

func SetupTodoRoutes(app fiber.Router, todoService *TodoService) {
	handler := NewTodoHandler(todoService)

	app.Get("/todos", handler.GetTodos)
	app.Post("/todos", handler.CreateTodo)
	app.Put("/todos/:id", handler.UpdateTodo)
	app.Delete("/todos/:id", handler.DeleteTodo)
}
