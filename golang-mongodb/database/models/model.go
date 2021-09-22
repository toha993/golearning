package models

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

type Currency struct {
	Id             string `bson:"_id"`
	CompanyId      string `json:"companyid"`
	ProjectId      string `json:"projectid"`
	CountryCode    string `json:"countrycode"`
	CurrencyName   string `json:"currencyname"`
	CurrencyCode   string `json:"currencycode"`
	DisplayName    string `json:"displayname"`
	CurrencySymbol string `json:"currencysymbol"`
}
