package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ReturnTweetsFollowers : Es la estructura con la que devolvemos los tweets
type ReturnTweetsFollowers struct {
	ID                 primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID             string             `bson:"userID" json:"userID,omitempty"`
	UserRelationshipID string             `bson:"userRelationshipID" json:"userRelationshipID,omitempty"`
	Tweets             struct {
		Message   string    `bson:"message" json:"message,omitempty"`
		CreatedAt time.Time `bson:"createdAt" json:"createdAt,omitempty"`
		ID        string    `bson:"_id" json:"_id,omitempty"`
	}
}
