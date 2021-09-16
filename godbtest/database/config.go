package database

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Country struct {
	Id            string   `bson:"_id"`
	CompanyId     string   `json:"companyid"`
	ProjectId     string   `json:"projectid"`
	Code          string   `json:"code"`
	Name          string   `json:"name"`
	ItrceCode     string   `json:"itrcecode"`
	ISO3          string   `json:"iso3"`
	ISO2          string   `json:"iso2"`
	IRS           string   `json:"irs"`
	IsGroup       bool     `json:"isgroup"`
	Region        string   `json:"region"`
	Tags          []string `json:"tags"`
	DisplayName   string   `json:"displayname"`
	EYDisplayName string   `json:"eydisplayname"`
}

var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
var client, err = mongo.Connect(context.TODO(), clientOptions)

func Config() {

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func ReadById(a interface{}) {

	v := reflect.ValueOf(a)

	//fmt.Println("Number of fields", v.NumField())

	col := client.Database(v.Field(0).String()).Collection(v.Field(1).String())

	var bw Country
	bw.Id = v.Field(2).String()

	err = col.FindOne(context.TODO(), bson.M{"_id": bson.M{"$eq": bw.Id}}).Decode(&bw)

	if err != nil {
		log.Println("Error in query!")
	}

	log.Println(bw)

}
