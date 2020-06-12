package models

// Tweet : Captura del body, es el mensaje que nos llega
type Tweet struct {
	Message string `bson:"message" json:"message"`
}
