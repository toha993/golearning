package main

import (
	db "main.go/database"
)

type dbdata struct {
	dbname         string
	collectionName string
	documentid     string
}

func main() {

	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// client, err := mongo.Connect(context.TODO(), clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = client.Ping(context.TODO(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Connected to MongoDB!")

	db.Config()

	r := dbdata{"_toha_Alert-DEV", "country", "tt"}
	db.ReadById(r)
	// col := client.Database("_toha_Alert-DEV").Collection("country")

	// var bw Country
	// bw.Id = "bw"

	// //finding one document

	// err = col.FindOne(context.TODO(), bson.M{"_id": bson.M{"$eq": bw.Id}}).Decode(&bw)

	// if err != nil {
	// 	log.Println("Error in query!")
	// }

	// log.Println(bw)

	//finding all documents

	// var countries []Country

	// cur, err := col.Find(context.TODO(), bson.D{{}})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for cur.Next(context.TODO()) {
	// 	var elem Country
	// 	err := cur.Decode(&elem)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	countries = append(countries, elem)
	// }

	// if err := cur.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// cur.Close(context.TODO())

	// fmt.Printf("Found multiple documents : %+v\n", countries)

	//updating one document

	// result, err := col.UpdateOne(context.TODO(), bson.M{"_id": bson.M{"$eq": bw.Id}}, bson.D{
	// 	{"$set", bson.D{{"Name", "Botswana"}}}})

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)

	//Deleting one document

	// deleted, err := col.DeleteOne(context.TODO(), bson.M{"_id": bson.M{"$eq": bw.Id}})

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("DeleteOne removed %v document(s)\n", deleted.DeletedCount)

	//Inserting one document

	// insertResult, err := col.InsertOne(context.TODO(), bson.D{
	// 	{Key: "_id", Value: "bw"},
	// 	{Key: "CompanyId", Value: " "},
	// 	{Key: "Code", Value: "BW"},
	// 	{Key: "Name", Value: "Botswana"},
	// 	{Key: "ItrceCode", Value: "BW"},
	// 	{Key: "ISO3", Value: "BWA"},
	// 	{Key: "ISO2", Value: "BW"},
	// 	{Key: "IRS", Value: "BC"},
	// 	{Key: "IsGroup", Value: "False"},
	// 	{Key: "Region", Value: "Africa"},
	// 	{Key: "Tags", Value: bson.A{"Rate", "IndividualRate", "VatRate"}},
	// 	{Key: "DisplayName", Value: "Botswana"},
	// 	{Key: "EYDisplayName", Value: "Botswana"},
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Inserted %v documents into the collection!\n", insertResult.InsertedID)

	//	[{Id:bw CompanyId: ProjectId: Code:BW Name:Botswana ItrceCode:BW ISO3:BWA ISO2:BW IRS:BC IsGroup:false Region:Africa Tags:[Rate IndividualRate VatRate] DisplayName:Botswana EYDisplayName:Botswana}]

}
