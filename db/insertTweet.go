package db

import (
	"context"
	"time"

	"github.com/JeanCntrs/social-network-like-twitter-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertTweet : Intenta almacenar un tweet en la base de datos
func InsertTweet(t models.SaveTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("SocialNetworkLikeTwitter")
	coll := db.Collection("tweets")

	// Se crea un documento exactamente igual al que viene, pero en formato bson
	registry := bson.M{
		"userID":    t.UserID,
		"message":   t.Message,
		"createdAt": t.CreatedAt,
	}

	result, err := coll.InsertOne(ctx, registry)
	if err != nil {
		return "", false, err
	}

	objID := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
