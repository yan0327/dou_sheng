package service

type UserRegisterRequest struct {
	UserName string `json:"user_name,omitempty"`
	PassWord string `json:"password, omitempty`
}
