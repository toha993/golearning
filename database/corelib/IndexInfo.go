package corelib

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Indexinfo struct {
	CompanyId    string
	ProjectId    string
	CompanyIndex string
	Index        string
	Topic        string
	Group        string
}

var ClientIndexFormat string = "[IndexPrefix]-Client"
var ClientIndexFormatWithCompanyId string = "[IndexPrefix]-[CompanyId]"
var ClientIndexFormatWithDemoId string = "[IndexPrefix]-DEMO-[DemoId]"
var _companyIndexFormat string

func CompanyIndexFormat() string {
	err := godotenv.Load("dev.env")
	if err != nil {
		log.Fatal("error in environment file")
	}
	_companyIndexFormat = os.Getenv("CompanyIndexFormat")
	if len(_companyIndexFormat) == 0 {
		_companyIndexFormat = ClientIndexFormat
	}
	return _companyIndexFormat
}

func getCompanyIndex(companyId string) string {
	var userInfo = GetUserDetails()
	var demoId = userInfo.DemoId
	if strings.Compare(companyId, "orbitaxsupport") == 0 && len(demoId) != 0 {
		GetDemoIndex(demoId)
	}
	index := strings.Replace(CompanyIndexFormat(), "[IndexPrefix]", GetIndexPrefix(), -1)
	index = strings.Replace(index, "[CompanyId]", companyId, -1)
	return index
}

func GetDemoIndex(demoId string) string {
	demoIndex := strings.Replace(ClientIndexFormatWithDemoId, "[IndexPrefix]", GetIndexPrefix(), -1)
	demoIndex = strings.Replace(demoIndex, "[DemoId]", demoId, -1)
	return demoIndex
}

//index
//companyId, projectId
//companyId, wildcard
//companyId, projectId, index
func IndexInfo(arg ...interface{}) Indexinfo {
	switch len(arg) {
	case 1:
		return Indexinfo{
			CompanyId:    "",
			ProjectId:    "",
			CompanyIndex: "",
			Index:        fmt.Sprintf("%v", arg[0]),
			Topic:        "",
			Group:        "",
		}
	case 2:
		{
			index := getCompanyIndex(fmt.Sprintf("%v", arg[0]))
			switch fmt.Sprintf("%T", arg[1]) {
			case "bool":
				{
					CompanyId := fmt.Sprintf("%v", arg[0])
					Topic := GetIndex(index, CompanyId)
					Group := GetIndex(index, CompanyId)
					return Indexinfo{
						CompanyId:    CompanyId,
						ProjectId:    "",
						CompanyIndex: index,
						Index:        index,
						Topic:        Topic,
						Group:        Group,
					}
				}
			case "string":
				CompanyId := fmt.Sprintf("%v", arg[0])
				ProjectId := fmt.Sprintf("%v", arg[1])
				Topic := GetIndex(index, CompanyId, ProjectId)
				Group := GetIndex(index, CompanyId, ProjectId)
				return Indexinfo{
					CompanyId:    CompanyId,
					ProjectId:    ProjectId,
					CompanyIndex: index,
					Index:        index,
					Topic:        Topic,
					Group:        Group,
				}

			}
		}
	case 3:
		CompanyId := fmt.Sprintf("%v", arg[0])
		ProjectId := fmt.Sprintf("%v", arg[1])
		CompanyIndex := fmt.Sprintf("%v", arg[2])
		Topic := GetIndex(CompanyIndex, CompanyId, ProjectId)
		Group := GetIndex(CompanyIndex, CompanyId, ProjectId)
		return Indexinfo{
			CompanyId:    CompanyId,
			ProjectId:    ProjectId,
			CompanyIndex: CompanyIndex,
			Index:        CompanyIndex,
			Topic:        Topic,
			Group:        Group,
		}

	}
	return Indexinfo{}
}

func GetIndex(arg ...string) string {
	if len(arg) >= 2 {
		if len(arg[1]) != 0 || len(arg[2]) != 0 {
			return getCompanyIndex(arg[1])
		}
	}
	return arg[0]
}
