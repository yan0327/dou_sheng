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

// type CreateCommonParams struct {
// 	Id         uint32 `gorm: "column:id"`
// 	UserId     uint32 `gorm: "column:user_id"`
// 	VideoId    uint32 `gorm: "column:video_id"`
// 	User       *User
// 	ActionType int8
// 	Content    string `gorm: "column:content"`
// 	CommentId  uint32
// 	CreateDate string `gorm: "column:create_time"`
// }

func (this Comment) TableName() string {
	return "tiktok_video_comment"
}

func (this Comment) CreateComment(db *gorm.DB) (Comment, error) {
	var err error
	db.Table("tiktok_user").Where("username = ?", this.User.UserName).Find(this.User)
	this.UserId = this.User.ID
	if this.ActionType == 1 {
		err := db.Omit("CreateDate").Create(&this).Error
		if err != nil {
			return Comment{}, nil
		}
		user := User{
			ID: this.UserId,
		}
		db.Where("id = ?", this.Id).Find(&this)
		this.User = user.CommentGetUserInfo(db)
		var isFollow int
		db.Table("tiktok_relation").Where("user_id = ? AND follower_id = ? AND action_type = ?", this.User.ID, this.UserId, 1).Count(&isFollow)
		if isFollow >= 1 {
			this.User.IsFollow = true
		}
		return this, err
	} else {
		err = db.Delete(&this, this.CommentId).Error
	}
	return Comment{}, err
}

func (this Comment) GetCommentList(db *gorm.DB) ([]Comment, error) {
	db.Table("tiktok_user").Where("username = ?", this.User.UserName).Find(this.User)
	comments := []Comment{}
	err := db.Table("tiktok_video_comment").Where("video_id = ?", this.VideoId).Order("create_time desc").Find(&comments).Error
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(comments); i++ {
		user := User{
			ID: comments[i].UserId,
		}
		comments[i].User = user.CommentGetUserInfo(db)
		var isFollow int
		db.Table("tiktok_relation").Where("user_id = ? AND follower_id = ? AND action_type = ?", comments[i].User.ID, this.User.ID, 1).Count(&isFollow)
		if isFollow >= 1 {
			comments[i].User.IsFollow = true
		}
	}
	return comments, nil
}
