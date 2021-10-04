package main

import (
	"Package/Orbitaxcrew/mongodb"
	"fmt"
)

func main() {
	col := mongodb.NewMongoCollection("_toha_Alert-DEV", "country")
	fmt.Println(col.GetAll())

}
