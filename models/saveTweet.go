package models

import "time"

// SaveTweet : Es el formato o estructura que tendrá el tweet en la base de datos
type SaveTweet struct {
	UserID    string    `bson:"userID" json:"userID,omitempty"`
	Message   string    `bson:"message" json:"message,omitempty"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt,omitempty"`
}
