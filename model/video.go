package model

import "time"

type Video struct {
	ID         uint      `json:"id" gorm:"primarykey"`
	AuthorId   uint      `json:"author_id"`
	PlayUrl    string    `json:"play_url,omitempty"`
	CoverUrl   string    `json:"cover_url,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	Title      string    `json:"title"`
}

type ReplyVideo struct {
	ID            uint      `json:"id" gorm:"primarykey"`
	Author        ReplyUser `json:"author"`
	PlayUrl       string    `json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	FavoriteCount int64     `json:"favorite_count"`
	CommentCount  int64     `json:"comment_count"`
	IsFavorite    bool      `json:"is_favorite"`
	Title         string    `json:"title"`
}

func (this Video) TableName() string {
	return "tiktok_video"
}
