package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConnect = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb://127.0.0.1:27017/twitter-backend-clone")

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

	log.Println("********************** the connection was success **********************")
	return client
}

func CheckConnection() int {
	err := MongoConnect.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}