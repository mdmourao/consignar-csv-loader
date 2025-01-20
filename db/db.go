package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Connect() {
	var err error
	mongoUri := os.Getenv("MONGO_URI")
	if mongoUri == "" {
		panic("MONGO_URI is not set")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoUri).SetServerAPIOptions(serverAPI)

	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}

}

func Disconnect() {
	if client != nil {
		client.Disconnect(context.TODO())
	}
}
