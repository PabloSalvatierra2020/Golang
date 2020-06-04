package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN es el objeto de conexion a la base de datos
var MongoCN = ConnectDB()

var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:admingo@cluster0-ggxk1.mongodb.net/test")

//ConnectDB es la funcion que me permite conectar a la DB
func ConnectDB() *mongo.Client {

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

	log.Println("conexion existosa con la BD")
	return client

}

//CheckedConnection verifica y chequea la conexion
func CheckedConnection() int {

	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
