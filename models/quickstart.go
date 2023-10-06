package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Quickstart struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}
