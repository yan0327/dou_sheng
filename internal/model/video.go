package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Video struct {
	Id       int64  `json:"id" gorm:"column:id"`
	AuthorId int64  `json:"author_id,omitempty" gorm:"column:author_id"`
	Author   *User  `json:"author,omitempty"`
	PlayUrl  string `json:"play_url,omitempty" gorm:"column:play_url"`
	CoverUrl string `json:"cover_url,omitempty" gorm:"column:cover_url"`

	User          *User
	FavoriteCount int64 `json:"favorite_count,omitempty"`
	CommentCount  int64 `json:"comment_count,omitempty"`
	IsFavorite    bool  `json:"is_favorite,omitempty"`

	UserName string
	LastTime int64
}

type VideoPush struct {
	AuthorId uint32 `gorm:"column:author_id"`
	PlayUrl  string `gorm:"column:play_url"`
	CoverUrl string `gorm:"column:cover_url"`
	UserName string `gorm:"-"`
}

type Favorite struct {
	UserId     uint32 `gorm:"column:user_id"`
	VideoId    uint32 `gorm:"column:video_id"`
	ActionType int    `gorm:"column:action_type"`
}

func (this Video) TableName() string {
	return "Video"
}
func (this *Video) ReverseFeed(db *gorm.DB, lastTime int64) ([]Video, error) {
	videos := make([]Video, 0)
	format := time.Unix(int64(this.LastTime), 0).Format("2006-01-02 15:04:05")
	err := db.Table("tiktok_video").Where("create_time <= ?", format).Order("id").Limit(30).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (this *VideoPush) Publish(db *gorm.DB) error {
	user := User{}
	err := db.Table("tiktok_user").Where("username = ?", this.UserName).Find(&user).Error
	if err != nil {
		return err
	}
	this.AuthorId = user.ID
	err = db.Table("tiktok_video").Create(this).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *Favorite) FavoriteAction(db *gorm.DB) error {
	fav := Favorite{}
	err := db.Table("tiktok_video_like").Where("user_id = ? AND video_id = ?", this.UserId, this.VideoId).First(&fav).Error
	if err == nil && fav.ActionType != this.ActionType {
		db.Table("tiktok_video_like").Where("user_id = ? AND video_id = ?", this.UserId, this.VideoId).Update("action_type", this.ActionType)
	}
	if err == gorm.ErrRecordNotFound {
		err := db.Table("tiktok_video_like").Create(this).Error
		return err
	}
	return err
}
