package service

import "mime/multipart"

type PublishRequest struct {
	Token string                `form:"token" binding:"required"`
	Data  *multipart.FileHeader `form:"data"  binding:"required"`
}

type PublishListRequest struct {
	Token string `form:"token" binding:"required"`
}
