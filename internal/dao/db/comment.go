package db

import (
	"gorm.io/gorm"
	"simple-demo/internal/model"
)

type CommentDao interface {
	Create(comment *model.Comment) (*model.Comment, error)
	Delete(commentId int64) error
	FindByVideo(videoId int64) ([]*model.Comment, error)
}

type comments struct {
	db *gorm.DB
}

func MakeComments(db *gorm.DB) *comments {
	return &comments{db}
}

func (c comments) Create(comment *model.Comment) (*model.Comment, error) {
	res := c.db.Create(comment)
	return comment, res.Error
}

func (c comments) Delete(commentId int64) error {
	res := c.db.Delete(&model.Comment{}, commentId)
	return res.Error
}

func (c comments) FindByVideo(videoId int64) ([]*model.Comment, error) {
	var list []*model.Comment
	res := c.db.Where("video_id = ?", videoId).
		Preload("User").
		Find(&list)
	return list, res.Error
}
