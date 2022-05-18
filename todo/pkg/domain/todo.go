package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `json:"title"`
	Details   string             `json:"details"`
	Completed bool               `json:"completed"`
}
