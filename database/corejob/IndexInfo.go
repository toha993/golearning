package corejob

import (
	"fmt"
	"strings"
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
	if len(arg) == 1 {
		return Indexinfo{
			CompanyId:    "",
			ProjectId:    "",
			CompanyIndex: "",
			Index:        fmt.Sprintf("%v", arg[0]),
			Topic:        "",
			Group:        "",
		}
	} else if len(arg) == 2 {
		fmt.Println("jerewfwegfw")
		index := getCompanyIndex(fmt.Sprintf("%v", arg[0]))
		if fmt.Sprintf("%T", arg[1]) == "bool" {
			CompanyId := fmt.Sprintf("%v", arg[0])
			Topic := GetIndex(CompanyId)
			Group := GetIndex(CompanyId)
			return Indexinfo{
				CompanyId:    CompanyId,
				ProjectId:    "",
				CompanyIndex: index,
				Index:        index,
				Topic:        Topic,
				Group:        Group,
			}
		} else {
			CompanyId := fmt.Sprintf("%v", arg[0])
			ProjectId := fmt.Sprintf("%v", arg[1])
			Topic := GetIndex(CompanyId, ProjectId)
			Group := GetIndex(CompanyId, ProjectId)
			return Indexinfo{
				CompanyId:    CompanyId,
				ProjectId:    ProjectId,
				CompanyIndex: index,
				Index:        index,
				Topic:        Topic,
				Group:        Group,
			}
		}
	} else if len(arg) == 3 {
		CompanyId := fmt.Sprintf("%v", arg[0])
		ProjectId := fmt.Sprintf("%v", arg[1])
		CompanyIndex := fmt.Sprintf("%v", arg[2])
		Topic := GetIndex(CompanyId, ProjectId)
		Group := GetIndex(CompanyId, ProjectId)
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
	if len(arg) >= 1 {
		return getCompanyIndex(arg[0])
	}
	return "anything"
}

// func IsClientIndex(id []string) bool {
// 	//return string.IsNullOrEmpty(CompanyId) == false || IsProjectIndex()

// 	if len(arg[0]) == 0 || IsProjectIndex(arg) {
// 		return true
// 	}

// 	return false
// }

// func (i Indexinfo) IsProjectIndex() bool {
// 	//return string.IsNullOrEmpty(ProjectId) == false

// 	if len(i.ProjectId) == 0 {
// 		return false
// 	}
// 	return true
// }
