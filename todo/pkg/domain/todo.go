package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Title     string             `json:"title"`
	Details   string             `json:"description" bson:"details"`
	Completed bool               `json:"completed"`
}
