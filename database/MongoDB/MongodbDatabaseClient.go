package mongodb

import (
	"context"
	"mongo-db/config"
	core "mongo-db/corejob"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type MongoCollection struct {
// 	database   string
// 	collection string
// }

// func NewMongoCollection(database, collection string) contracts.IBulkRepositoryContext {
// 	return &MongoCollection{
// 		database:   database,
// 		collection: collection,
// 	}
// }

func Save(db core.Indexinfo, data interface{}, Id string) error {

	filter := bson.M{"id": Id}
	opts := options.Update().SetUpsert(true)

	update := bson.D{primitive.E{Key: "$set", Value: data}}

	client, err := config.GetDBInstance()

	if err != nil {
		return err
	}

	collection := client.Database(db.Index).Collection(r.collection)

	_, err = collection.UpdateOne(context.TODO(), filter, update, opts)

	return err
}

// func Insert(data interface{}) error {

// 	client, err := config.GetDBInstance()
// 	if err != nil {
// 		return err
// 	}

// 	collection := client.Database(r.database).Collection(r.collection)
// 	_, err = collection.InsertOne(context.TODO(), data)

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func GetAll() (interface{}, error) {

// 	var result []interface{}

// 	client, err := config.GetDBInstance()
// 	if err != nil {
// 		return result, err
// 	}

// 	collection := client.Database(r.database).Collection(r.collection)
// 	cur, err := collection.Find(context.TODO(), bson.M{})

// 	if err != nil {
// 		return result, nil
// 	}

// 	for cur.Next(context.TODO()) {
// 		var elem interface{}

// 		cur.Decode(&elem)
// 		result = append(result, elem)
// 	}

// 	return result, nil
// }

// func GetId(id string) (interface{}, error) {
// 	var result interface{}
// 	filter := bson.M{"id": id}

// 	client, err := config.GetDBInstance()
// 	if err != nil {
// 		return result, err
// 	}

// 	collection := client.Database(r.database).Collection(r.collection)
// 	err = collection.FindOne(context.TODO(), filter).Decode(&result)

// 	if err != nil {
// 		return result, err
// 	}

// 	return result, nil
// }

// func DeleteById(id string) error {

// 	filter := bson.M{"id": id}

// 	client, err := config.GetDBInstance()
// 	if err != nil {
// 		return err
// 	}

// 	collection := client.Database(r.database).Collection(r.collection)
// 	_, err = collection.DeleteOne(context.TODO(), filter)
// 	return err
// }
