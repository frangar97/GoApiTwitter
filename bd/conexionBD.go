package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoC = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://frangar97:96794518fran@cluster0.drj9m.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

//ConectarBD permite realizar la conexion con mongo
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		return client
	}

	log.Println("Conexion con BD")
	return client
}

//ChequeoConection verfica si existe conexion con base de datos
func ChequeoConection() int {
	err := MongoC.Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}

	return 1
}
