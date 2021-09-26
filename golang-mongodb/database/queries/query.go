package queries

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	Save(string, interface{}) error
}

func (r *MongoCollection) Save(id string, data interface{}) {

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel() // should abandon its work after the timeout elapses.

	// v := reflect.ValueOf(data)
	// id := v.Field(0).String()

	filter := bson.M{"_id": id}

	//upsert does insert if no document found else do update
	opts := options.Update().SetUpsert(true)
	update := bson.D{primitive.E{Key: "$set", Value: data}}

	result, err := r.db.Collection(r.collection).UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Fatal("Error in query")
	}
	fmt.Println(result)

}

func (r *MongoCollection) ReadAll() (interface{}, error) {

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel() // should abandon its work after the timeout elapses.

	cur, err := r.db.Collection(r.collection).Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	var result []interface{}
	//getting all documents in cursor and iterating it to append in result array
	for cur.Next(context.TODO()) {
		var elem interface{}

		cur.Decode(&elem)
		result = append(result, elem)
	}

	return result, nil
}

func (r *MongoCollection) ReadById(id string) (interface{}, error) {
	var result interface{}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel() // should abandon its work after the timeout elapses.

	filter := bson.M{"_id": id}
	err := r.db.Collection(r.collection).FindOne(ctx, filter).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}

func (r *MongoCollection) DeleteById(id string) error {

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel() // should abandon its work after the timeout elapses.

	filter := bson.M{"_id": id}
	_, err := r.db.Collection(r.collection).DeleteOne(ctx, filter)
	return err
}
