package model

import "time"

type Comment struct {
	Id        int64     `json:"id,omitempty"`
	UserId    int64     `json:"-"`
	User      *User     `json:"user"`
	VideoId   int64     `json:"-"`
	Video     *Video    `json:"-"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"create_date,string" gorm:"column:create_time"`
}

func (this Comment) TableName() string {
	return "tiktok_video_comment"
}
