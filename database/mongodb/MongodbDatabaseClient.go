package mongodb

import (
	"context"
	"mongodbgolang/config"
	"mongodbgolang/corelib"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Save(index corelib.Indexinfo, data interface{}, id string) error {

	filter := bson.M{"_id": id}
	opts := options.Update().SetUpsert(true)

	update := bson.D{primitive.E{Key: "$set", Value: data}}

	client, err := config.GetDBInstance()

	if err != nil {
		return err
	}
	database := index.Index
	collectionName := strings.ToLower(reflect.TypeOf(data).Name())

	collection := client.Database(database).Collection(collectionName)

	_, err = collection.UpdateOne(context.TODO(), filter, update, opts)

	return err
}

func Insert(index corelib.Indexinfo, data interface{}) error {

	client, err := config.GetDBInstance()
	if err != nil {
		return err
	}
	database := index.Index
	collectionName := strings.ToLower(reflect.TypeOf(data).Name())

	collection := client.Database(database).Collection(collectionName)
	_, err = collection.InsertOne(context.TODO(), data)

	if err != nil {
		return err
	}
	return nil
}

func GetAll(index corelib.Indexinfo, data interface{}) (interface{}, error) {

	var result []interface{}

	client, err := config.GetDBInstance()
	if err != nil {
		return result, err
	}
	database := index.Index
	collectionName := strings.ToLower(reflect.TypeOf(data).Name())

	collection := client.Database(database).Collection(collectionName)
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		return result, nil
	}

	for cur.Next(context.TODO()) {
		var elem interface{}

		cur.Decode(&elem)
		result = append(result, elem)
	}

	return result, nil
}

func GetId(index corelib.Indexinfo, data interface{}, id string) (interface{}, error) {
	var result interface{}
	filter := bson.M{"_id": id}

	client, err := config.GetDBInstance()
	if err != nil {
		return result, err
	}

	database := index.Index
	collectionName := strings.ToLower(reflect.TypeOf(data).Name())

	collection := client.Database(database).Collection(collectionName)
	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func DeleteById(index corelib.Indexinfo, data interface{}, id string) error {

	filter := bson.M{"_id": id}

	client, err := config.GetDBInstance()
	if err != nil {
		return err
	}
	database := index.Index
	collectionName := strings.ToLower(reflect.TypeOf(data).Name())

	collection := client.Database(database).Collection(collectionName)
	_, err = collection.DeleteOne(context.TODO(), filter)
	return err
}
