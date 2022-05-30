package model

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	Id         int64  `json:"id" gorm:"column:id"`
	UserId     int64  `json:"user_id" gorm:"column:user_id"`
	VideoId    int64  `json:"video_id"  gorm:"column:video_id"`
	User       *User  `gorm:"-" json:"user"`
	ActionType uint8  `gorm:"-"`
	Content    string `json:"content" gorm:"column:content"`
	CommentId  int64  `gorm:"-"`
	CreateDate string `json:"create_date" gorm:"column:create_time"`
}

func (c Comment) TableName() string {
	return "tiktok_video_comment"
}

func (c Comment) CreateComment(db *gorm.DB) (Comment, error) {
	var err error
	db.Table("tiktok_user").Where("username = ?", c.User.UserName).Find(c.User)
	c.UserId = c.User.ID
	if c.ActionType == 1 {
		err := db.Omit("CreateDate").Create(&c).Error
		if err != nil {
			return Comment{}, nil
		}
		user := User{
			ID: c.UserId,
		}
		db.Where("id = ?", c.Id).Find(&c)
		c.User = user.CommentGetUserInfo(db)
		var isFollow int
		db.Table("tiktok_relation").Where("user_id = ? AND follower_id = ? AND action_type = ?", c.User.ID, c.UserId, 1).Count(&isFollow)
		if isFollow >= 1 {
			c.User.IsFollow = true
		}
		return c, err
	} else {
		err = db.Delete(&c, c.CommentId).Error
	}
	return Comment{}, err
}

func (c Comment) GetCommentList(db *gorm.DB) ([]Comment, error) {
	db.Table("tiktok_user").Where("username = ?", c.User.UserName).Find(c.User)
	comments := []Comment{}
	err := db.Table("tiktok_video_comment").Where("video_id = ?", c.VideoId).Order("create_time desc").Find(&comments).Error
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(comments); i++ {
		user := User{
			ID: comments[i].UserId,
		}
		comments[i].User = user.CommentGetUserInfo(db)
		var isFollow int
		db.Table("tiktok_relation").Where("user_id = ? AND follower_id = ? AND action_type = ?", comments[i].User.ID, c.User.ID, 1).Count(&isFollow)
		if isFollow >= 1 {
			comments[i].User.IsFollow = true
		}
	}
	return comments, nil
}
