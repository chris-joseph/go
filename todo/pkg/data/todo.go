package data

import (
	"context"
	"fmt"
	"log"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"todo/pkg/config"
	"todo/pkg/domain"
)

type ITodoProvider interface {
	CreateTodo(todo *domain.Todo) error
	TodoExists(id primitive.ObjectID) (bool, error)
	FindTodoById(id string) (*domain.Todo, error)
	GetAllTodos() ([]*domain.Todo, error)
	DeleteTodos([]primitive.ObjectID) error
	//UpdateTodo(todo *domain.Todo) error
}

type TodoProvider struct {
	todoCollection *mongo.Collection
	ctx            context.Context
}

func NewTodoProvider(cfg *config.Settings, mongo *mongo.Client) ITodoProvider {
	todoCollection := mongo.Database(cfg.DbName).Collection("todos")
	return &TodoProvider{
		todoCollection: todoCollection,
		ctx:            context.TODO(),
	}
}

func (u TodoProvider) CreateTodo(user *domain.Todo) error {
	_, err := u.todoCollection.InsertOne(u.ctx, user)
	if err != nil {
		return errors.Wrap(err, "Error inserting todo")
	}
	return nil
}

func (u TodoProvider) FindTodoById(id string) (*domain.Todo, error) {
	var todoFound *domain.Todo
	filter := bson.D{primitive.E{Key: "id", Value: id}}

	if err := u.todoCollection.FindOne(u.ctx, filter).Decode(&todoFound); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(err, "Todo not found")
		}
		return nil, errors.Wrap(err, "Error finding by id")
	}
	return todoFound, nil

}

func (u TodoProvider) TodoExists(id primitive.ObjectID) (bool, error) {
	var todoFound *domain.Todo

	filter := bson.D{primitive.E{Key: "id", Value: id}}

	if err := u.todoCollection.FindOne(u.ctx, filter).Decode(&todoFound); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, errors.Wrap(err, "Error finding by id")
	}
	return true, nil
}

func (u TodoProvider) GetAllTodos() ([]*domain.Todo, error) {
	var todos []*domain.Todo
	filter := bson.D{}
	cursor, err := u.todoCollection.Find(u.ctx, filter)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Error finding todos")
	}
	cursor.All(u.ctx, &todos)

	return todos, nil
}

func (u TodoProvider) DeleteTodos(ids []primitive.ObjectID) error {
	fmt.Println(ids)
	filter := bson.M{"_id": bson.M{"$in": ids}}
	result, err := u.todoCollection.DeleteMany(u.ctx, filter)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return errors.Wrap(err, "Error finding todos")
	}
	if result.DeletedCount == 0 {
		return errors.New("No todos deleted")
	}
	log.Printf("Deleted %d todos", result.DeletedCount)

	return nil
}
