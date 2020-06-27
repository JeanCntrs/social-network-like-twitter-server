package db

import (
	"context"
	"fmt"
	"time"

	"github.com/JeanCntrs/social-network-like-twitter-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadAllUsers : Lee los usuarios registrados en el sistema, si se escribe "R" en quienes trae solo los que se relacionan conmigo
func ReadAllUsers(ID string, page int64, search string, searchType string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("SocialNetworkLikeTwitter")
	coll := db.Collection("users")

	var results []*models.User
	findOption := options.Find()
	findOption.SetSkip((page - 1) * 20)
	findOption.SetLimit(20)

	// Indiferente de mayúsculas o minúsculas
	query := bson.M{
		"names": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := coll.Find(ctx, query, findOption)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var find, include bool

	for cursor.Next(ctx) {
		var u models.User
		err := cursor.Decode(&u)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relation
		r.UserID = ID
		r.UserRelationshipID = u.ID.Hex()

		include = false

		find, err = ConsultRelation(r)
		if searchType == "new" && find == false {
			include = true
		}
		if searchType == "follow" && find == true {
			include = true
		}
		if r.UserRelationshipID == ID {
			include = false
		}

		if include == true {
			u.Password = ""
			u.Biography = ""
			u.Website = ""
			u.Location = ""
			u.Banner = ""
			u.Email = ""

			results = append(results, &u)
		}
	}

	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cursor.Close(ctx)

	return results, true
}
