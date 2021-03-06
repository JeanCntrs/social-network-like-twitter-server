package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ReturnTweets : Es la estructura con la que devolvemos los tweets
type ReturnTweets struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID    string             `bson:"userID" json:"userID,omiteempty"`
	Message   string             `bson:"message" json:"message,omitempty"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt,omitempty"`
}
