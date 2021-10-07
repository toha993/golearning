package corejob

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func GetClientIndex() Indexinfo {
	err := godotenv.Load("dev.env")
	if err != nil {
		log.Fatal("error in environment file")
	}
	index := strings.Replace(ClientIndexFormat, "[IndexPrefix]", GetIndexPrefix(), -1)
	return IndexInfo(index)
}

func GetGlobalIndex() Indexinfo {
	err := godotenv.Load("dev.env")
	if err != nil {
		log.Fatal("error in environment file")
	}
	index := os.Getenv("GlobalPrefix") + os.Getenv("IndexPrefix")
	return IndexInfo(index)
}

func GetIndexPrefix() string {
	err := godotenv.Load("dev.env")
	if err != nil {
		log.Fatal("error in environment file")
	}
	return os.Getenv("DevPrefix") + os.Getenv("IndexPrefix")
}

func GetDevIndex() Indexinfo {

	err := godotenv.Load("dev.env")
	if err != nil {
		log.Fatal("error in environment file")
	}

	version := strings.ToUpper(os.Getenv("Version"))
	index := os.Getenv("GlobalPrefix") + os.Getenv("IndexPrefix") + "-" + version

	return IndexInfo(index)
}

func GetCompanyIndex() Indexinfo {
	var userInfo = GetUserDetails()
	var companyId = userInfo.CompanyId
	return IndexInfo(strings.ToLower(companyId), false)
}
