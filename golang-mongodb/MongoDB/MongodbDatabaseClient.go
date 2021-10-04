package mongodb

import (
	"context"

	"Package/Orbitaxcrew/config"
	"Package/Orbitaxcrew/contracts"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCollection struct {
	database   string
	collection string
}

func NewMongoCollection(database, collection string) contracts.IBulkRepositoryContext {
	return &MongoCollection{
		database:   database,
		collection: collection,
	}
}

func (r *MongoCollection) Save(data interface{}, Id string) error {

	filter := bson.M{"_id": Id}
	opts := options.Update().SetUpsert(true)

	update := bson.D{primitive.E{Key: "$set", Value: data}}

	client, err := config.GetClientInstance()
	if err != nil {
		return err
	}
	collection := client.Database(r.database).Collection(r.collection)

	_, err = collection.UpdateOne(context.TODO(), filter, update, opts)

	return err
}

func (r *MongoCollection) Insert(data interface{}) error {

	client, err := config.GetClientInstance()
	if err != nil {
		return err
	}

	collection := client.Database(r.database).Collection(r.collection)
	_, err = collection.InsertOne(context.TODO(), data)

	if err != nil {
		return err
	}
	return nil
}

func (r *MongoCollection) GetAll() (interface{}, error) {

	var result []interface{}

	client, err := config.GetClientInstance()
	if err != nil {
		return result, err
	}

	collection := client.Database(r.database).Collection(r.collection)
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

func (r *MongoCollection) ReadById(id string) (interface{}, error) {
	var result interface{}
	filter := bson.M{"_id": id}

	client, err := config.GetClientInstance()
	if err != nil {
		return result, err
	}

	collection := client.Database(r.database).Collection(r.collection)
	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *MongoCollection) DeleteById(id string) error {

	filter := bson.M{"_id": id}

	client, err := config.GetClientInstance()
	if err != nil {
		return err
	}

	collection := client.Database(r.database).Collection(r.collection)
	_, err = collection.DeleteOne(context.TODO(), filter)
	return err
}
