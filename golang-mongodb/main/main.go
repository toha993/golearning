package main

import (
	"fmt"
	db "golang-mongodb/database/connection"
	model "golang-mongodb/database/models"
	query "golang-mongodb/database/queries"
)

func main() {

	dbconfig := db.Config("_toha_Alert-DEV")

	//country
	country := query.NewMongoCollection(dbconfig, "country")

	//readbyid function
	// output, _ := country.ReadById("tt")
	// fmt.Println(output)

	//save function
	user := model.Country{
		Id:            "bw",
		CompanyId:     " ",
		ProjectId:     " ",
		Code:          "BW",
		Name:          "Botswana",
		ItrceCode:     "BW",
		ISO3:          "BWA",
		ISO2:          "BW",
		IRS:           "BC",
		IsGroup:       false,
		Region:        "Africa",
		Tags:          []string{"Rate", "IndividualRate", "VatRate"},
		DisplayName:   "Botswana",
		EYDisplayName: "Botswana",
	}
	country.Save("bw", user)

	//readall
	// result, _ := country.ReadAll()
	// fmt.Println(result)

	//delete
	//country.DeleteById("bw")

	//country

	//currency
	currency := query.NewMongoCollection(dbconfig, "currency")

	output, _ := currency.ReadById("bw")
	fmt.Println(output)

	//currency

}
