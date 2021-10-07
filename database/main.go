package main

import (
	"fmt"
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

func modeltype(c interface{}) string {
	fmt.Println(c)
	return reflect.TypeOf(c).Elem().Name()

}

func main() {

	v := Country{}
	fmt.Println(modeltype(&v))

}