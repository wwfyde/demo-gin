package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/wwfyde/demo-gin/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"strconv"
)

const (
	coll = "podcast"
)

type PodcastHandler struct { // https://github.com/PauloPortugal/gin-gonic-rest-mongodb/blob/main/handlers/books.go
	ctx context.Context
	cfg *viper.Viper
	MDB *mongo.Client
	RDB *redis.Client
}

func NewPodcastHandler(ctx context.Context, cfg *viper.Viper, mdb *mongo.Client, rdb *redis.Client) *PodcastHandler {
	return &PodcastHandler{
		ctx: ctx,
		cfg: cfg,
		MDB: mdb,
		RDB: rdb,
	}
}

func (handler *PodcastHandler) GetPodcast(c *gin.Context) {
	var podcast models.Podcast
	//var podcast bson.M
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	err := handler.MDB.Database(handler.cfg.GetString("database.mongo.db")).
		Collection(coll).FindOne(handler.ctx, bson.M{"uid": 1}).Decode(&podcast)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			fmt.Printf("No document found with the specified ID\n")
		} else {
			log.Fatalf("Error find by id %[2]s, err: %[1]v\n", err, id)
		}
	}
	c.JSON(http.StatusOK, podcast)
}
