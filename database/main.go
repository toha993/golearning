package main

import (
	"fmt"
	core "mongodbgolang/corejob"
	query "mongodbgolang/mongodb"
	"reflect"
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

type User struct {
	UserName  string `json:"UserName"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	CompanyId string `json:"CompanyId"`
}

func modeltype(c interface{}) string {
	// v := c
	fmt.Println(c)
	return reflect.TypeOf(c).Name()

}

func main() {

	v := Country{
		Id:            "r2r2r23",
		CompanyId:     "23rr",
		ProjectId:     "2324",
		Code:          "",
		Name:          "Botswana",
		ItrceCode:     "",
		ISO3:          "",
		ISO2:          "",
		IRS:           "",
		IsGroup:       false,
		Region:        "",
		Tags:          []string{},
		DisplayName:   "",
		EYDisplayName: "",
	}
	// fmt.Println(modeltype(v))

	// help := User{
	// 	UserName:  "toha",
	// 	FirstName: "",
	// 	LastName:  "",
	// 	Email:     "",
	// 	CompanyId: "",
	// }
	// fmt.Println(modeltype(help))
	// fmt.Println(len("") == 0)

	query.Save(core.GetCompanyIndex(), v, "r2r2r23")

	//success

}
