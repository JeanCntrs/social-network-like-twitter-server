package db

import (
	"context"
	"time"

	"github.com/JeanCntrs/social-network-like-twitter-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ModifyRegister : Permite modificar el perfil de usuario
func ModifyRegister(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("SocialNetworkLikeTwitter")
	coll := db.Collection("users")

	registry := make(map[string]interface{})
	if len(u.Names) > 0 {
		registry["names"] = u.Names
	}
	if len(u.Surnames) > 0 {
		registry["surnames"] = u.Surnames
	}
	registry["birthdate"] = u.Birthdate
	if len(u.Avatar) > 0 {
		registry["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		registry["banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		registry["biography"] = u.Biography
	}
	if len(u.Location) > 0 {
		registry["location"] = u.Location
	}
	if len(u.Website) > 0 {
		registry["website"] = u.Website
	}

	updateString := bson.M{
		"$set": registry,
	}

	// ObjectIDFromHex : Convierte mi ID que es un string en un formato ObjectID
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := coll.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil
}
