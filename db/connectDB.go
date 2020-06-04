package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnection : Se exporta esta función con el fin de ser utilizada en cualquier archivo dentro de la carpeta db, es el objeto de conexión
var MongoConnection = ConnectMongo()

// ConnectMongo : Función que retorna la conexión a la base de datos
func ConnectMongo() *mongo.Client {
	// Context : Permiten comunicar información entre ejecución y ejecución. También permite setear valores, por ejemplo TimeOut
	// Context.TODO : Conectar sin restricción, es por default
	setConnectionString()
	var clientOptions = options.Client().ApplyURI(os.Getenv("CONNECTION_STRING"))

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
