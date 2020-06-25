package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ReturnTweet : Es la estructura con la que devolvemos los tweets
type ReturnTweets struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID    string             `bson:"userId" json:"userId,omiteempty"`
	Message   string             `bson:"message" json:"message,omitempty"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt,omitempty"`
}
