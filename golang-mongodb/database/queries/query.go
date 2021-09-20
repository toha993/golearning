package queries

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	model "golang-mongodb/database/models"
)

type MongoCollection struct {
	db         *mongo.Database
	collection string
}

func NewMongoCollection(db *mongo.Database, collection string) *MongoCollection {
	return &MongoCollection{
		db:         db,
		collection: collection,
	}
}

type Operation interface {
	ReadAll() (interface{}, error)
	ReadById(string) (interface{}, error)
	DeleteById(string) error
	Save(interface{}) error
}

func (r *MongoCollection) Save(data interface{}) {
	oldStruct, ok := data.(model.Country)
	if !ok {
		log.Fatal("Data not found")
	}
	newStruct := model.Country(oldStruct)

	filter := bson.M{"_id": newStruct.Id}

	fmt.Println(newStruct)
	opts := options.Update().SetUpsert(true)
	update := bson.D{primitive.E{Key: "$set", Value: data}}

	fmt.Println(update)
	result, err := r.db.Collection(r.collection).UpdateOne(context.TODO(), filter, update, opts)

	if err != nil {
		log.Fatal("Error in query")
	}
	fmt.Println(result)

	// opts := options.Update().SetUpsert(true)

	// update := bson.D{primitive.E{Key: "$set", Value: data}}

	// result, err := r.db.Collection(r.collection).UpdateOne(context.TODO(), filter, update, opts)
	// if err != nil {
	// 	log.Fatal("Error in query")
	// }
	// fmt.Println(result)

}

func (r *MongoCollection) ReadAll() (interface{}, error) {
	cur, err := r.db.Collection(r.collection).Find(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	var result []interface{}

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
	err := r.db.Collection(r.collection).FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}

func (r *MongoCollection) DeleteById(id string) error {

	filter := bson.M{"_id": id}
	_, err := r.db.Collection(r.collection).DeleteOne(context.TODO(), filter)
	return err
}
