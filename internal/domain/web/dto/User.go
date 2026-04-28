package dto

type UserRegistrationResponse struct {
}

type Token struct {
	AccessToken     string
	RefreshToken    string
	ExpAccessToken  int
	ExpRefreshToken int
}

type UserCredentials struct {
	Login    string
	Password string
	UserName string
}
