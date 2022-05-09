package service

type UserLoginRequest struct {
	UserName string `form:"username" binding:"required, max=32"`
	PassWord string `form:"password" binding:"required, max=32"`
}

type UserInfoRequest struct {
	UserId uint32 `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}
