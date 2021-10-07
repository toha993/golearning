package corejob

import (
	"fmt"
	"strings"
)

var ClientIndexFormat string = "[IndexPrefix]-Client"
var ClientIndexFormatWithCompanyId string = "[IndexPrefix]-[CompanyId]"
var ClientIndexFormatWithDemoId string = "[IndexPrefix]-DEMO-[DemoId]"

type Indexinfo struct {
	CompanyId    string
	ProjectId    string
	CompanyIndex string
	Index        string
	Topic        string
	Group        string
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
		index := getCompanyIndex(fmt.Sprintf("%v", arg[0]))
		if fmt.Sprintf("%T", arg[1]) == "bool" {
			return Indexinfo{
				CompanyId:    fmt.Sprintf("%v", arg[0]),
				ProjectId:    "",
				CompanyIndex: index,
				Index:        index,
				Topic:        "",
				Group:        "",
			}
		} else {
			return Indexinfo{
				CompanyId:    fmt.Sprintf("%v", arg[0]),
				ProjectId:    fmt.Sprintf("%v", arg[1]),
				CompanyIndex: index,
				Index:        index,
				Topic:        "",
				Group:        "",
			}
		}
	} else if len(arg) == 3 {
		return Indexinfo{
			CompanyId:    fmt.Sprintf("%v", arg[0]),
			ProjectId:    fmt.Sprintf("%v", arg[1]),
			CompanyIndex: fmt.Sprintf("%v", arg[2]),
			Index:        fmt.Sprintf("%v", arg[2]),
			Topic:        "",
			Group:        "",
		}
	}

	return Indexinfo{}
}

func getCompanyIndex(companyId string) string {
	index := strings.Replace(ClientIndexFormat, "[IndexPrefix]", GetIndexPrefix(), -1)
	strings.Replace(index, "[CompanyId]", companyId, -1)
	return index
}

// func (i Indexinfo) GetIndex() string {
// 	if i.IsClientIndex() {
// 		return getCompanyIndex(i.CompanyId)
// 	}
// 	return i.Index
// }

// func (i Indexinfo) IsClientIndex() bool {
// 	//return string.IsNullOrEmpty(CompanyId) == false || IsProjectIndex()

// 	if len(i.CompanyId) == 0 || i.IsProjectIndex() {
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
