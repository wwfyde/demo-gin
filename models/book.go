package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Book is used for storage upcoming book info
type Book struct {
	ID  primitive.ObjectID
	col *mongo.Collection
}
