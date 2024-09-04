package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		log.Fatal("MONGOURI non défini dans le fichier .env")
	}
	if databaseName == "" {
		log.Fatal("MONGODBDB non défini dans le fichier .env")
	}

	if databaseCollection == "" {
		log.Fatal("MONGOCOLLECTION non défini dans le fichier .env")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}

	Client = client
	log.Println("Connected to MongoDB!")

	Collection = Client.Database(databaseName).Collection(databaseCollection)
}
