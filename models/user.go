package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User : Es el modelo de usuario de la base de datos en MongoDB
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Names     string             `bson:"names" json:"names,omitempty"`
	Surnames  string             `bson:"surnames" json:"surnames,omitempty"`
	Birthdate time.Time          `bson:"birthdate" json:birthdate",omitempty`
	Email     string             `bson:"email" json:email"`
	Password  string             `bson:"password" json:password,omitempty"`
	Avatar    string             `bson:"avatar" json:avatar,omitempty"`
	Banner    string             `bson:"banner" json:banner,omitempty"`
	Biography string             `bson:"biography" json:biography,omitempty"`
	Location  string             `bson:"location" json:location,omitempty"`
	Website   string             `bson:"website" json:website,omitempty"`
}
