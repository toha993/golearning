package models

type Country struct {
	Id            string
	CompanyId     string
	ProjectId     string
	Code          string
	Name          string
	ItrceCode     string
	ISO3          string
	ISO2          string
	IRS           string
	IsGroup       bool
	Region        string
	Tags          []string
	DisplayName   string
	EYDisplayName string
}
