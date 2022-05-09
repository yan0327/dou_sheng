package service

type UserRegisterRequest struct {
	UserName string `form:"username" binding:"required, max=32"`
	PassWord string `form:"password" binding:"required, max=32"`
}
