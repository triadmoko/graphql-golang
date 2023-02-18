package config

import (
	"context"
	"log"

	"github.com/triadmoko/grahpql-golang/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

func InitMongoDB() *mongo.Database {
	clientOptions := options.Client().ApplyURI(helpers.LoadEnv("MONGO_URL"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	database := client.Database("graphql")
	return database
}
