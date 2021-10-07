package corejob

func GetUserDetails() (UserInfo) {
	var claims = GetClaims()
	if claims == nil{
		return UserInfo{}
	}

	return UserInfo{
		UserName : claims.UserName,
		FirstName : claims.FirstName,
		LastName : claims.LastName,
		Email : claims.Email,
		CompanyId : claims.CompanyId,
		CompanyName : claims.CompanyName,
		ProductCode : claims.ProductCode,
		IsCheckPointUser : claims.IsCheckPointUser,
	}
}