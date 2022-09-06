package models

type RequestCreateUser struct {
	UserName string `validate:"required" json:"user_name"`
	Password string `validate:"required" json:"password"`
}

type RequestLoginUser struct {
	UserName string `validate:"required" json:"user_name"`
	Password string `validate:"required" json:"password"`
}

type ResponseLoginUser struct {
	UserName    string `json:"user_name"`
	AccessToken string `json:"access_token"`
}
