package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type 播客 = Podcast
type Podcast struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UID    int                `bson:"uid" json:"uid"`
	Title  string             `bson:"title,omitempty" json:"title"`
	Author string             `bson:"author,omitempty" json:"author"`
	Tags   []string           `bson:"tags,omitempty" json:"tags"`
}

type 章节 = Episode

type Episode struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Podcast     primitive.ObjectID `bson:"podcast,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	Duration    int32              `bson:"duration,omitempty"`
}

func (p *Podcast) All(ctx context.Context, c *mongo.Collection) {

}
