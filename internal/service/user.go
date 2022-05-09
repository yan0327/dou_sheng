package service

type UserLoginRequest struct {
	UserName string `json:"user_name,omitempty"`
	PassWord string `json:"password, omitempty`
}

type UserInfoRequest struct {
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token",omitempty"`
}
