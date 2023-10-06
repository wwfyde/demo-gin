package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// mongodb

type MongoSource interface {
	/* ... */
}

type MongoManager struct {
	Db *mongo.Database
	C  *mongo.Collection
}

// Quick Start: Golang & MongoDB - Modeling Documents with Go Data Structures
type Model struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}

var db *mongo.Database
