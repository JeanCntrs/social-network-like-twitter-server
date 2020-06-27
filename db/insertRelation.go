package db

import (
	"context"
	"time"

	"github.com/JeanCntrs/social-network-like-twitter-server/models"
)

// InsertRelation : Graba relación en la base de datos
func InsertRelation(rel models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("SocialNetworkLikeTwitter")
	coll := db.Collection("relations")

	_, err := coll.InsertOne(ctx, rel)
	if err != nil {
		return false, err
	}

	return true, nil
}
