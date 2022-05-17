package model

type Realtion struct {
	ID          uint  `json:"id,omitempty" gorm:"primarykey"`
	UserId      uint  `json:"user_id,omitempty"`
	FollowerId  uint  `json:"follower_id,omitempty"`
	IsEffective uint8 `json:"is_effective,omitempty"`
}

func (this Realtion) TableName() string {
	return "tiktok_relation"
}
