package main

import (
	"fmt"
	"mongodbgolang/corelib"
	"mongodbgolang/mongodb"
)

type ProjectModel struct {
	Id         string `bson:"_id" json:"id"`
	CompanyId  string
	ProjectId  string
	IsReadOnly bool
	IsDeleted  bool
}

func main() {
	//fmt.Println("Bismillah")

	data := ProjectModel{
		Id:         "a5d721c65fcb4000818c0438b1f22f48",
		CompanyId:  "orbitax",
		ProjectId:  "a5d721c65fcb4000818c0438b1f22f48",
		IsReadOnly: false,
		IsDeleted:  true,
	}

	//insert
	//mongodb.Insert(corelib.GetCompanyIndex(), data)

	//Getid
	//fmt.Println(mongodb.GetId(corelib.GetCompanyIndex(), data, data.Id))

	//delete
	//mongodb.DeleteById(corelib.GetCompanyIndex(), data, "a5d721c65fcb4000818c0438b1f22f48")

	//save
	//mongodb.Save(corelib.GetCompanyIndex(), data, data.Id)

	//Getall
	fmt.Println(mongodb.GetAll(corelib.GetCompanyIndex(), data))

	//success

}
