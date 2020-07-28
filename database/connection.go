package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnect exports database connection created
var MongoConnect = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://u_db_twinstgo:dykD8mJs8cfXsjx8@twinstgo.00mmr.mongodb.net/test")

// ConnectDB function that connect application to database
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
	log.Println("Success database connection!")
	return client
}

// CheckConnection ping to database
func CheckConnection() bool {
	err := MongoConnect.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
