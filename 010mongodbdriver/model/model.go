package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Netflix struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Movie   string             `bson:"movie,omitempty"`
	Watched bool               `bson:"watched,omitempty"`
}