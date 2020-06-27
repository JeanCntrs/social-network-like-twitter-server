package db

import (
	"context"
	"time"

	"github.com/JeanCntrs/social-network-like-twitter-server/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ConsultRelation : Consulta relación entre 2 usuarios
func ConsultRelation(rel models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("SocialNetworkLikeTwitter")
	coll := db.Collection("relations")

	condition := bson.M{
		"userID":             rel.UserID,
		"userRelationshipID": rel.UserRelationshipID,
	}

	var result models.Relation
	err := coll.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return false, err
	}

	return true, nil
}
