package service

import (
	"mime/multipart"
	"simple-demo/global"
	"simple-demo/model"
)

type PublishRequest struct {
	Token string                `form:"token" binding:"required"`
	Data  *multipart.FileHeader `form:"data"  binding:"required"`
}

type PublishListRequest struct {
	Token string `form:"token" binding:"required"`
}

func PublishVideo(id uint) (video []model.Video, err error) {
	err = global.DBEngine.Where("author_id = ?", id).Find(&video).Error
	return
}
