package db

import (
	"context"
	"log"
	"time"

	"github.com/JeanCntrs/social-network-like-twitter-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadTweets : Lee los tweets de un perfil
func ReadTweets(ID string, page int64) ([]*models.ReturnTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoConnection.Database("SocialNetworkLikeTwitter")
	coll := db.Collection(("tweets"))

	var results []*models.ReturnTweets

	condition := bson.M{
		"userID": ID,
	}

	tweetOptions := options.Find()
	tweetOptions.SetLimit(20)
	tweetOptions.SetSort(bson.D{{Key: "createdAt", Value: -1}})
	tweetOptions.SetSkip((page - 1) * 20)

	cursor, err := coll.Find(ctx, condition, tweetOptions)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var registry models.ReturnTweets
		err := cursor.Decode(&registry)
		if err != nil {
			return results, false
		}
		results = append(results, &registry)
	}

	return results, true
}
