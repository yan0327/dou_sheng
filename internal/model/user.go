package model

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID            int64  `json:"id" gorm:"primary_key"`
	UserName      string `json:"name" gorm:"column:username"`
	PassWord      string `json:"-" gorm:"column:password"`
	FollowCount   int64  `json:"follow_count" gorm:"-"`
	FollowerCount int64  `json:"follower_count"  gorm:"-"`
	IsFollow      bool   `json:"is_follow"  gorm:"-"`
}

func (u User) TableName() string {
	return "tiktok_user"
}

func (u User) Register(db *gorm.DB) (*User, error) {
	err := db.Create(&u).Error
	return &u, err
}

func (u User) UserLogin(db *gorm.DB) (*User, error) {
	user := &User{}
	err := db.Where("username = ?", u.UserName).Find(&user).Error
	if err != nil {
		return nil, err
	}
	if u.PassWord != user.PassWord {
		return nil, errors.New("password err")
	}
	return user, nil
}

func (u User) GetUserInfo(db *gorm.DB) (*User, error) {
	user := User{}
	err := db.Where("username = ?", u.UserName).Find(&user).Error
	if err != nil {
		return &user, err
	}
	db.Table("tiktok_relation").Where("user_id = ? AND action_type = ?", user.ID, 1).Count(&user.FollowCount)
	db.Table("tiktok_relation").Where("follower_id = ? AND action_type = ?", user.ID, 1).Count(&user.FollowerCount)
	return &user, nil
}

func (u User) VideoGetUserInfo(db *gorm.DB) *User {
	user := User{ID: u.ID}
	db.Where("id = ?", u.ID).Find(&user)
	db.Table("tiktok_relation").Where("user_id = ? AND action_type = ?", u.ID, 1).Count(&user.FollowCount)
	db.Table("tiktok_relation").Where("follower_id = ? AND action_type = ?", u.ID, 1).Count(&user.FollowerCount)
	return &user
}

func (u User) CommentGetUserInfo(db *gorm.DB) *User {
	user := User{ID: u.ID}
	db.Where("id = ?", u.ID).Find(&user)
	db.Table("tiktok_relation").Where("user_id = ? AND action_type = ?", u.ID, 1).Count(&user.FollowCount)
	db.Table("tiktok_relation").Where("follower_id = ? AND action_type = ?", u.ID, 1).Count(&user.FollowerCount)
	return &user
}
