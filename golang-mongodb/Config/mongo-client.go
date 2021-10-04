package config

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() (*mongo.Client, error){
	err := godotenv.Load(".env")
	if err != nil{
		return nil,err
	}

	MongoDb := os.Getenv("MONGODB_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	
	
	if err != nil{
		return client,err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	
	return client, err
}

var Client *mongo.Client = nil

func GetClientInstance() (*mongo.Client, error){
	var err error = nil
	if Client == nil{
		Client,err = DBinstance()
	}
	return Client,err
}