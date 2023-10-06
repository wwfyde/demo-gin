package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type MongoDb struct {
	Db *mongo.Database
}
