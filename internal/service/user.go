package service

type UserLoginRequest struct {
	UserName string `json:"username" binding:"required, max=32"`
	PassWord string `json:"password" binding:"required, max=32"`
}

type UserInfoRequest struct {
	UserId uint32  `json:"user_id" binding:"required"`
	Token  string `json:"token" binding:"required"`
}
