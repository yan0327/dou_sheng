package model

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// type User struct {
// 	Id            int64  `json:"id,omitempty"`
// 	Name          string `json:"name,omitempty"`
// 	FollowCount   int64  `json:"follow_count,omitempty"`
// 	FollowerCount int64  `json:"follower_count,omitempty"`
// 	IsFollow      bool   `json:"is_follow,omitempty"`
// }
type User struct {
	ID            uint32 `json:"id" gorm:"primary_key"`
	UserName      string `json:"name" gorm:"column:username"`
	PassWord      string `json:"-" gorm:"column:password"`
	FollowCount   uint32 `json:"follow_count" gorm:"-"`
	FollowerCount uint32 `json:"follower_count"  gorm:"-"`
	IsFollow      bool   `json:"is_follow"  gorm:"-"`
}

func (this User) TableName() string {
	return "tiktok_user"
}

func (this User) Register(db *gorm.DB) (uint32, error) {
	err := db.Table("tiktok_user").Create(&this).Error
	return this.ID, err
}

func (this User) UserLogin(db *gorm.DB) (uint32, error) {
	user := &User{}
	err := db.Table("tiktok_user").Where("username = ?", this.UserName).Find(&user).Error
	if err != nil {
		return user.ID, err
	}
	if this.PassWord != user.PassWord {
		return user.ID, errors.New("password err")
	}
	return user.ID, nil
}

func (this User) GetUserInfo(db *gorm.DB) (User, error) {
	user := User{}
	err := db.Table("tiktok_user").Where("username = ?", this.UserName).Find(&user).Error
	if err != nil {
		return user, err
	}
	this.ID = user.ID
	var cnt uint32
	err = db.Table("tiktok_relation").Where("user_id = ? AND action_type = ?", this.ID, 1).Count(&cnt).Error
	if err != nil {
		return user, err
	}
	user.FollowCount = cnt
	err = db.Table("tiktok_relation").Where("follower_id = ? AND action_type = ?", this.ID, 1).Count(&cnt).Error
	if err != nil {
		return user, err
	}
	user.FollowerCount = cnt
	return user, nil
}

func (this User) VideoGetUserInfo(db *gorm.DB) *User {
	user := User{ID: this.ID}
	db.Table("tiktok_user").Where("id = ?", this.ID).First(&user)
	db.Table("tiktok_relation").Where("user_id = ? AND action_type = ?", this.ID, 1).Count(&user.FollowCount)
	db.Table("tiktok_relation").Where("follower_id = ? AND action_type = ?", this.ID, 1).Count(&user.FollowerCount)
	return &user
}

func (this User) CommentGetUserInfo(db *gorm.DB) *User {
	user := User{ID: this.ID}
	db.Table("tiktok_user").Where("id = ?", this.ID).First(&user)
	db.Table("tiktok_relation").Where("user_id = ? AND action_type = ?", this.ID, 1).Count(&user.FollowCount)
	db.Table("tiktok_relation").Where("follower_id = ? AND action_type = ?", this.ID, 1).Count(&user.FollowerCount)
	return &user
}
