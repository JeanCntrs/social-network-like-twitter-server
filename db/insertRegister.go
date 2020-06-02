package db

import (
	"context"
	"time"

	"github.com/JeanCntrs/social-network-like-twitter-server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertRegister : Es el final del endpoint, donde inserto los datos
func InsertRegister(u models.User) (string, bool, error) {
	ctn, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("SocialNetworkLikeTwitter")
	coll := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := coll.InsertOne(ctn, u)
	if err != nil {
		return "", false, err
	}

	ObjectID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjectID.String(), true, nil
}
