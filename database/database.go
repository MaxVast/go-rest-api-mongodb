package database

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"os"
	"time"
)

var Client *mongo.Client
var Collection *mongo.Collection

func Connect() {
	mongoURI := os.Getenv("MONGOURI")
	databaseName := os.Getenv("MONGODBDB")
	databaseCollection := os.Getenv("MONGOCOLLECTION")

	if mongoURI == "" {
		log.Fatal("MONGOURI not defined in .env file")
	}
	if databaseName == "" {
		log.Fatal("MONGODBDB not defined in .env file")
	}
	if databaseCollection == "" {
		log.Fatal("MONGOCOLLECTION not defined in .env file")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}

	Client = client
	log.Println("Connected to MongoDB!")

	Collection = Client.Database(databaseName).Collection(databaseCollection)
}
