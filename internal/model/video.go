package model

import (
	"time"
)

type Video struct {
	Id            int64     `json:"id,omitempty"`
	AuthorId      int64     `json:"-"`
	Author        *User     `json:"author"`
	Title         string    `json:"title"`
	PlayUrl       string    `json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	FavoriteCount int64     `json:"favorite_count"`
	CommentCount  int64     `json:"comment_count"`
	IsFavorite    bool      `json:"is_favorite"`
	CreatedAt     time.Time `json:"-" gorm:"column:create_time"`
}

func (this Video) TableName() string {
	return "tiktok_video"
}
