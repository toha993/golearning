package config

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client

var clientInstanceError error

var mongoOnce sync.Once

func GetDBInstance() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		err := godotenv.Load("dev.env")
		if err != nil {
			log.Fatal("error in environment file")
		}
		clientOptions := options.Client().ApplyURI(os.Getenv("MongodbConnection"))
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}