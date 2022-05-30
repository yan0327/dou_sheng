package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Video struct {
	Id            int64  `json:"id" gorm:"column:id"`
	AuthorId      int64  `json:"author_id,omitempty" gorm:"column:author_id"`
	Author        *User  `json:"author,omitempty"`
	PlayUrl       string `json:"play_url,omitempty" gorm:"column:play_url"`
	CoverUrl      string `json:"cover_url,omitempty" gorm:"column:cover_url"`
	Title         string `json:"title,omitempty" gorm:"title" `
	CreateTime    string `json:"create_time,omitempty" gorm:"column:create_time" `
	User          *User  `gorm:"-"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`

	UserName string
	LastTime int64
}

type VideoPush struct {
	AuthorId int64  `gorm:"column:author_id"`
	PlayUrl  string `gorm:"column:play_url"`
	CoverUrl string `gorm:"column:cover_url"`
	UserName string `gorm:"-"`
	Title    string `gorm:"column:title"`
}

type Favorite struct {
	UserId     int64  `gorm:"column:user_id"`
	VideoId    int64  `gorm:"column:video_id"`
	ActionType int    `gorm:"column:action_type"`
	UserName   string `gorm:"-"`
}

func (v Video) TableName() string {
	return "tiktok_video"
}
func (v *Video) ReverseFeed(db *gorm.DB, lastTime int64) ([]*Video, error) {
	videos := make([]*Video, 0)
	format := time.Unix(int64(time.Now().Unix()), 0).Format("2006-01-02 15:04:05")
	err := db.Table("tiktok_video").Where("create_time <= ?", format).Order("create_time desc").Limit(20).Find(&videos).Error
	for i := 0; i < len(videos); i++ {
		// videos[i] = &Video{}
		user := User{ID: videos[i].AuthorId}
		videos[i].Author = user.VideoGetUserInfo(db)
		db.Table("tiktok_video_like").Where("video_id = ?", videos[i].Id).Count(&videos[i].FavoriteCount)
		db.Table("tiktok_video_comment").Where("video_id = ?", videos[i].Id).Count(&videos[i].CommentCount)
	}
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (v *Video) PublishList(db *gorm.DB) ([]Video, error) {
	// db.Table("tiktok_user").Where("username = ?", this.User.UserName).Find(this.User)
	videos := make([]Video, 0)
	// db.Table("tiktok_video").Select("*").Joins("inner join tiktok_video_like on tiktok_video.id = tiktok_video_like.video_id").Where("tiktok_video_like.user_id = ?", this.User.ID).Find(&videos)
	db.Table("tiktok_video").Where("author_id = ?", v.User.ID).Find(&videos)
	user := User{ID: v.User.ID}
	author := user.VideoGetUserInfo(db)
	for i := 0; i < len(videos); i++ {
		videos[i].Author = author
		db.Table("tiktok_video_like").Where("video_id = ?", videos[i].Id).Count(&videos[i].FavoriteCount)
		db.Table("tiktok_video_comment").Where("video_id = ?", videos[i].Id).Count(&videos[i].CommentCount)
	}
	return videos, nil
}

func (v *VideoPush) Publish(db *gorm.DB) error {
	user := User{}
	err := db.Table("tiktok_user").Where("username = ?", v.UserName).Find(&user).Error
	if err != nil {
		return err
	}
	v.AuthorId = user.ID
	err = db.Table("tiktok_video").Create(v).Error
	if err != nil {
		return err
	}
	return nil
}

func (f *Favorite) FavoriteAction(db *gorm.DB) error {
	user := User{}
	db.Table("tiktok_user").Select("id").Where("username = ?", f.UserName).Find(&user)
	fav := Favorite{}
	f.UserId = user.ID
	err := db.Table("tiktok_video_like").Where("user_id = ? AND video_id = ?", user.ID, f.VideoId).First(&fav).Error
	if err == nil && fav.ActionType != f.ActionType {
		db.Table("tiktok_video_like").Where("user_id = ? AND video_id = ?", user.ID, f.VideoId).Update("action_type", f.ActionType)
	}
	if err == gorm.ErrRecordNotFound {
		err := db.Table("tiktok_video_like").Create(f).Error
		return err
	}
	return err
}

func (f *Favorite) FavoriteList(db *gorm.DB) ([]*Video, error) {
	user := User{}
	db.Table("tiktok_user").Select("id").Where("username = ?", f.UserName).Find(&user)
	f.UserId = user.ID
	favorites := []Favorite{}
	err := db.Table("tiktok_video_like").Where("user_id = ? AND action_type = ?", user.ID, 1).Find(&favorites).Error
	if err != nil {
		return nil, err
	}
	videos := make([]*Video, len(favorites))
	for i := range favorites {
		videos[i] = &Video{}
		db.Table("tiktok_video").Where("id = ?", favorites[i].VideoId).Find(&videos[i])
		user := &User{
			ID: videos[i].AuthorId,
		}
		user = user.VideoGetUserInfo(db)
		videos[i].Author = user
		db.Table("tiktok_video_like").Where("video_id = ? AND action_type = ?", favorites[i].VideoId, 1).Count(&videos[i].FavoriteCount)
		db.Table("tiktok_video_comment").Where("video_id = ?", favorites[i].VideoId).Count(&videos[i].CommentCount)
		var isFollow int
		videos[i].IsFavorite = true
		db.Table("tiktok_relation").Where("user_id = ? AND follower_id = ? AND action_type = ?", videos[i].AuthorId, user.ID, 1).Count(&isFollow)
		if isFollow >= 1 {
			videos[i].Author.IsFollow = true
		}
	}

	return videos, nil
}
