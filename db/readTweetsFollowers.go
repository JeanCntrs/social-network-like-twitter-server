package db

import (
	"context"
	"time"

	"github.com/JeanCntrs/social-network-like-twitter-server/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ReadTweetsFollowers : Leo los tweets de los seguidores
func ReadTweetsFollowers(ID string, page int) ([]models.ReturnTweetsFollowers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("SocialNetworkLikeTwitter")
	coll := db.Collection("relations")

	skip := (page - 1) * 20

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userID": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweets",
			"localField":   "userRelationshipID",
			"foreignField": "userID",
			"as":           "tweets",
		}})
	conditions = append(conditions, bson.M{"$unwind": "$tweets"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweets.createdAt": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursor, err := coll.Aggregate(ctx, conditions)
	var result []models.ReturnTweetsFollowers
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}

	return result, true
}
