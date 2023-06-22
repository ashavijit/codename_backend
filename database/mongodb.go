package database

import (
	"codename_backend/configs"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ConnectMongoDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(configs.EnvMONGO_URI()))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB! 🎉")
	return client
}

var DB *mongo.Database = ConnectMongoDB().Database("codename")

func GetCollection(collectionName string) *mongo.Collection {
	CollectionName := collectionName
	return DB.Collection(CollectionName)
}
