package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
) 

func DBInstance() *mongo.Client{
	MongoDBURL := os.Getenv("DB_CONNECTION")
	fmt.Print(MongoDBURL)

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDBURL))
	if err != nil{
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Connected to Mongo DB")
	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection{
	return client.Database("restaurant-crm").Collection(collectionName)
}