package models

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"todo/pkg/domain"
)

type RegisterRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type CreateTodoRequest struct {
	Title   string `json:"title"`
	Details string `json:"details"`
}

type DeleteRequest struct {
	IDS []primitive.ObjectID `bson:"ids" json:"ids"`
}

func ValidateRegisterRequest(c *gin.Context) (*domain.User, *Error) {
	registerRequest := new(RegisterRequest)
	if err := c.Bind(registerRequest); err != nil {
		return nil, BindError()
	}
	var validationErrors []string

	if len(registerRequest.Password) < 8 {
		validationErrors = append(validationErrors, "Password must be 8 characters")

	}

	if len(registerRequest.UserName) < 3 {
		validationErrors = append(validationErrors, "Username must be longer than 2 characters")

	}

	if len(validationErrors) > 0 {
		return nil, ValidationError(validationErrors)

	}
	return &domain.User{
		UserName: registerRequest.UserName,
		Password: registerRequest.Password,
	}, nil

}

func ValidateLoginRequest(c *gin.Context) (*domain.User, *Error) {
	loginRequest := new(LoginRequest)
	if err := c.Bind(loginRequest); err != nil {
		return nil, BindError()
	}
	var validationErrors []string

	if len(loginRequest.Password) < 8 {
		validationErrors = append(validationErrors, "Password must be 8 characters")

	}

	if len(loginRequest.UserName) < 3 {
		validationErrors = append(validationErrors, "Username must be longer than 2 characters")

	}

	if len(validationErrors) > 0 {
		return nil, ValidationError(validationErrors)

	}
	return &domain.User{
		UserName: loginRequest.UserName,
		Password: loginRequest.Password,
	}, nil

}

func ValidateDeletionRequest(c *gin.Context) ([]primitive.ObjectID, *Error) {
	deleteRequest := new(DeleteRequest)

	if err := c.Bind(deleteRequest); err != nil {
		return nil, BindError()
	}

	if len(deleteRequest.IDS) == 0 {
		return nil, ValidationError([]string{"No ids provided"})
	}

	return deleteRequest.IDS, nil

}
func ValidateCreateTodoRequest(c *gin.Context) (*domain.Todo, *Error) {
	createTodoRequest := new(CreateTodoRequest)
	if err := c.Bind(createTodoRequest); err != nil {
		return nil, BindError()
	}
	//TODO add validation errors
	return &domain.Todo{
		Title:   createTodoRequest.Title,
		Details: createTodoRequest.Details,
	}, nil
}
