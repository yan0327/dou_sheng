package service

import (
	"simple-demo/global"
	"simple-demo/model"
	"time"
)

type CommentRequest struct {
	UserId      uint32 `form:"user_id" binding:"required"`
	Token       string `form:"token" binding:"required"`
	VideoId     uint32 `form:"video_id" binding:"required"`
	ActionType  uint8  `form:"action_type" binding:"required"`
	CommentText string `form:"comment_text"`
	CommentId   uint32 `form:"comment_id"`
}

type CommentListRequest struct {
	UserId  uint32 `form:"user_id" binding:"required"`
	Token   string `form:"token" binding:"required"`
	VideoId uint32 `form:"video_id" binding:"required"`
}

func (svc *Service) CreateComment(param *CommentRequest) error {
	return nil
}

func (svc *Service) GetCommentList(param *CommentListRequest) ([]*model.Comment, error) {
	return nil, nil
}

func GetComments(ID uint) (comments []model.Comment, err error) {
	err = global.DBEngine.Model(&model.Comment{}).Where("video_id = ?", ID).Find(&comments).Error
	return
}
func GetCommentsNum(ID uint) (count int64, err error) {
	var comments []model.Comment
	err = global.DBEngine.Model(&model.Comment{}).Where("video_id = ?", ID).Find(&comments).Error
	count = int64(len(comments))
	return
}

func CreateComment(useid uint, video uint, comment string) (err error) {
	com := &model.Comment{
		UserId:     useid,
		VideoId:    video,
		Content:    comment,
		CreateTime: time.Now(),
	}
	result := global.DBEngine.Create(&com)
	return result.Error
}

func DeleteComment(comment_id uint) (err error) {
	com := &model.Comment{
		ID: comment_id,
	}
	err = global.DBEngine.Delete(&com).Error
	return
}
