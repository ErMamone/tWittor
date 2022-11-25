package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

/*MongoCn variable objeto de conexion para la base de datos*/
var MongoCn=dbConnect()
var clientOptions= options.Client().ApplyURI("mongodb+srv://tuvieja123:tuvieja123@twitternigo.lyze2.mongodb.net/?retryWrites=true&w=majority")


/*dbConnect es la funcion para conectarse a la base de datos*/
func dbConnect() *mongo.Client {
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

	log.Println("Success connection with DB")
	return client
}
/*ConnectionCheck ping a la db*/
func ConnectionCheck() int {
	err := MongoCn.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}