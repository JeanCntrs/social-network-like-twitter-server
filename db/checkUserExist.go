package db

import (
	"context"
	"time"

	"github.com/JeanCntrs/social-network-like-twitter-server/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckUserExist : Recibe un email de parámetro y comprueba si existe en la base de datos
func CheckUserExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("SocialNetworkLikeTwitter")
	coll := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err := coll.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
