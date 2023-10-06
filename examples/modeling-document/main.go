package main

import (
	"context"
	"errors"
	"github.com/wwfyde/demo-gin/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {

	podcast := models.Podcast{
		Title:  "A Go/Python Developer",
		Author: "Viakayn Meta",
		Tags:   []string{"Go", "Python", "MongoDB"},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:admin@localhost:27017"))
	if err != nil {
		panic(err)
	}
	db := client.Database("demo-gin")
	coll := db.Collection("podcast")
	//var result models.Podcast
	var result bson.M
	err = coll.FindOne(ctx, bson.M{"title": "A Go/Python Developer"}).Decode(&result)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Printf("Document not exist in MDB, insert\n")
			res, err := coll.InsertOne(ctx, podcast)
			if err != nil {
				//
				log.Fatalf("Error insert: %v\n", err)
			}
			log.Printf("Inserted Document with ID: %v\n", res.InsertedID)
		}

	} else {
		log.Printf("Find One :%v, \n%[2]v\n", result,
			primitive.ObjectID.Hex(
				result["_id"].(primitive.ObjectID),
			))
	}

}
