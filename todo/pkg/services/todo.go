package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"todo/pkg/config"
	"todo/pkg/data"
	"todo/pkg/domain"
	"todo/pkg/models"
)

type ITodoService interface {
	CreateTodo(todo *domain.Todo) *models.Error
	GetAllTodos() ([]*domain.Todo, *models.Error)
	DeleteTodos([]primitive.ObjectID) *models.Error
}

type TodoService struct {
	todoProvider data.ITodoProvider
	cfg          *config.Settings
}

func NewTodoService(cfg *config.Settings, todoProvider data.ITodoProvider) ITodoService {
	return &TodoService{
		todoProvider: todoProvider,
		cfg:          cfg,
	}
}

func (u TodoService) CreateTodo(todo *domain.Todo) *models.Error {

	todo.ID = primitive.NewObjectID()

	err := u.todoProvider.CreateTodo(todo)

	if err != nil {
		return &models.Error{
			Code:    500,
			Name:    "SERVER_ERROR",
			Message: "Someting went wrong",
			Error:   err,
		}
	}

	return nil
}
func (u TodoService) GetAllTodos() ([]*domain.Todo, *models.Error) {

	list, err := u.todoProvider.GetAllTodos()

	if err != nil {
		return nil, &models.Error{
			Code:    500,
			Name:    "SERVER_ERROR",
			Message: "Someting went wrong",
			Error:   err,
		}
	}

	return list, nil
}
func (u TodoService) DeleteTodos(ids []primitive.ObjectID) *models.Error {

	err := u.todoProvider.DeleteTodos(ids)

	if err != nil {
		return &models.Error{
			Code:    500,
			Name:    "SERVER_ERROR",
			Message: "Someting went wrong",
			Error:   err,
		}
	}

	return nil
}
