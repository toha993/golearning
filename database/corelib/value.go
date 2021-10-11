package corelib

func Dummydata() UserInfo {
	v := UserInfo{
		UserName:         "tonoy.akando@orbitax.com",
		FirstName:        "Tonoy",
		LastName:         "Akando",
		Email:            "tonoy.akando@orbitax.com",
		CompanyId:        "orbitaxbdltd",
		CompanyName:      "",
		ProductCode:      "",
		IsCheckPointUser: "",
		IsAdmin:          "",
		IsPowerUser:      "",
		DemoId:           "",
	}
	return v
}
