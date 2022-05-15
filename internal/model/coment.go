package model

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	Id         uint32 `gorm:"column:id"`
	UserId     uint32 `gorm:"column:user_id"`
	VideoId    uint32 `gorm:"column:video_id"`
	User       *User
	ActionType uint8  `gorm:"-"`
	Content    string `gorm:"column:content"`
	CommentId  uint32 `gorm:"-"`
	CreateDate string `gorm:"column:create_time"`
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

func (this Comment) CreateComment(db *gorm.DB) error {
	var err error
	if this.ActionType == 1 {
		err = db.Omit("CreateDate").Create(&this).Error
		return err
	} else {
		err = db.Delete(&this, this.CommentId).Error
	}
	return err
}

func (this Comment) GetCommentList(db *gorm.DB) ([]Comment, error) {
	comments := []Comment{}
	err := db.Where("video_id = ?", this.VideoId).Order("create_time desc").Find(&comments).Error
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(comments); i++ {
		user := User{
			ID: comments[i].UserId,
		}
		user, _ = user.GetUserInfo(db)
		comments[i].User = &user
		var isFollow int
		db.Table("tiktok_relation").Where("user_id = ? AND follower_id = ? AND action_type = ?", comments[i].User.ID, this.UserId, 1).Count(&isFollow)
		if isFollow >= 1 {
			comments[i].User.IsFollow = true
		}
	}
	return comments, nil
}
