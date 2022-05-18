package data

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"todo/pkg/config"
	"todo/pkg/domain"
)

type IUserProvider interface {
	CreateAccount(user *domain.User) error
	UsernameExists(username string) (bool, error)
	FindUserByName(username string) (*domain.User, error)
}

type UserProvider struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserProvider(cfg *config.Settings, mongo *mongo.Client) IUserProvider {
	userCollection := mongo.Database(cfg.DbName).Collection("users")
	return &UserProvider{
		userCollection: userCollection,
		ctx:            context.TODO(),
	}
}

func (u UserProvider) CreateAccount(user *domain.User) error {
	_, err := u.userCollection.InsertOne(u.ctx, user)
	if err != nil {
		return errors.Wrap(err, "Error inserting user")
	}
	return nil
}

func (u UserProvider) FindUserByName(username string) (*domain.User, error) {
	var userFound domain.User
	filter := bson.D{primitive.E{Key: "username", Value: username}}
	fmt.Println(filter)
	if err := u.userCollection.FindOne(u.ctx, filter).Decode(&userFound); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(err, "User not found")
		}
		return nil, errors.Wrap(err, "Error finding by username")
	}
	return &userFound, nil

}

func (u UserProvider) UsernameExists(username string) (bool, error) {
	var userFound *domain.User
	filter := bson.D{primitive.E{Key: "username", Value: username}}

	if err := u.userCollection.FindOne(u.ctx, filter).Decode(&userFound); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, errors.Wrap(err, "Error finding by username")
	}
	return true, nil
}
