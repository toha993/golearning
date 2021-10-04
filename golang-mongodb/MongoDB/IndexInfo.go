package mongodb

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type IndexInfo struct {
	CompanyId    string
	ProjectId    string
	CompanyIndex string
	Index        string
	Topic        string
	Group        string
}

func GetIndex(companyId string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error in environment file")
	}

	index := os.Getenv("IndexPrefix")
	result := index + "-" + strings.ToLower(companyId)

	return result
}

func GetIndexWithProjectId(companyId string, projectId string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error in environment file")
	}

	index := os.Getenv("IndexPrefix")
	result := index + "-" + strings.ToLower(companyId) + "-" + strings.ToLower(projectId)
	// "projects-orbitax-1231241414"
	// "projects-google-1231241414"
	// _toha_projects - DEV
	// _toha = devprefix
	// projects = IndexPrefix
	// dev

	return result
}

func GetIndexInfo(index string) IndexInfo {
	v := IndexInfo{}
	v.Index = index
	return v
}

func GetDevIndexInfo() IndexInfo {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error in environment file")
	}

	version := strings.ToUpper(os.Getenv("Version"))
	index := os.Getenv("GlobalPrefix") + os.Getenv("IndexPrefix") + "-" + version

	return GetIndexInfo(index)

}

// func GetIndexInfoWithCompanyAndProject(companyId string, projectId string) IndexInfo {
// 	v := IndexInfo{
// 		CompanyId:    companyId,
// 		ProjectId:    projectId,
// 		CompanyIndex: GetIndex(companyId),
// 		Index:        GetIndex(companyId),
// 		Topic:        GetIndexWithProjectId(companyId, projectId),
// 		Group:        GetIndexWithProjectId(companyId, projectId),
// 	}

// 	return v
// }

// public static IndexInfo GetIndex(this IConfigurationHelper source)
//         {
//             var index = source.GetIndexPrefix();
//             if (!String.IsNullOrEmpty(source["GlobalPrefix"]))
//             {
//                 index = source["GlobalPrefix"] + source["IndexPrefix"];
//             }

//             var useXversionIndex = source["UseXversionIndex"];
//             var version = source["Version"];
//             var isDisableXversionIndex = IsDisableXversionIndex(source);
//             if (isDisableXversionIndex == false && !String.IsNullOrEmpty(useXversionIndex) && !String.IsNullOrEmpty(version) && useXversionIndex.ToBool())
//             {
//                 index += $"-{version.ToUpper()}";
//             }
//             return new IndexInfo(index);
//         }
