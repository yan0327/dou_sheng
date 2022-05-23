package model

import "time"

type Comment struct {
	ID         uint      `json:"id" gorm:"primarykey"`
	UserId     uint      `json:"user_id"`
	VideoId    uint      `json:"video_id,omitempty"`
	Content    string    `json:"content,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
}
type ReplyComment struct {
	ID         uint      `json:"id" gorm:"primarykey"`
	User       ReplyUser `json:"user"`
	VideoId    uint      `json:"video_id,omitempty"`
	Content    string    `json:"content,omitempty"`
	CreateTime string    `json:"create_date,omitempty"`
}

func (this Comment) TableName() string {
	return "tiktok_video_comment"
}
