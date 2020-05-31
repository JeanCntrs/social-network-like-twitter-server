package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnection : Se exporta esta función con el fin de ser utilizada en cualquier archivo dentro de la carpeta db, es el objeto de conexión
var MongoConnection = ConnectMongo()
var clientOptions = options.Client().ApplyURI("mongodb+srv://JeanCntrs:qCOTVU1QQpfmmcmY@cluster0-4ct1j.mongodb.net/SocialNetworkLikeTwitter")

// ConnectMongo : Función que retorna la conexión a la base de datos
func ConnectMongo() *mongo.Client {
	// Context : Permiten comunicar información entre ejecución y ejecución. También permite setear valores, por ejemplo TimeOut
	// Context.TODO : Conectar sin restricción, es por default
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión base de datos exitosa")
	return client
}

// CheckConnection : Retorna el ping realizado a la base de datos
func CheckConnection() int {
	err := MongoConnection.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
