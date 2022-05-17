package model

type Like struct {
	ID         uint  `json:"id,omitempty" gorm:"primarykey"`
	UserId     uint  `json:"user_id"`
	VideoId    uint  `json:"video_id,omitempty"`
	ActionType uint8 `json:"action_type,omitempty"`
}

func (this Like) TableName() string {
	return "tiktok_video_like"
}
