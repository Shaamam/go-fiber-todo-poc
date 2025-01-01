package todo

import (
	"go-fiber-todo-poc/pkg/models"
	"gorm.io/gorm"
)

// TodoService provides methods to interact with todo-related data.
type TodoService struct {
	db *gorm.DB
}

// NewTodoService initializes a new TodoService with the given database connection.
func NewTodoService(db *gorm.DB) *TodoService {
	return &TodoService{db: db}
}

// GetTodos retrieves all todos from the database.
func (s *TodoService) GetTodos() ([]models.Todo, error) {
	var todos []models.Todo
	result := s.db.Find(&todos)
	return todos, result.Error
}

// CreateTodo adds a new todo to the database.
func (s *TodoService) CreateTodo(todo *models.Todo) error {
	result := s.db.Create(todo)
	return result.Error
}

// UpdateTodo updates an existing todo in the database.
func (s *TodoService) UpdateTodo(id string, updatedTodo *models.Todo) error {
	var todo models.Todo
	result := s.db.First(&todo, id)
	if result.Error != nil {
		return result.Error
	}
	result = s.db.Model(&todo).Updates(updatedTodo)
	return result.Error
}

// DeleteTodo removes a todo from the database.
func (s *TodoService) DeleteTodo(id string) error {
	result := s.db.Delete(&models.Todo{}, id)
	return result.Error
}
