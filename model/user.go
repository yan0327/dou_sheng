package model

type User struct {
	ID       uint   `json:"id,omitempty" gorm:"primarykey"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type ReplyUser struct {
	ID            uint   `json:"id,omitempty" gorm:"primarykey"`
	Username      string `json:"username,omitempty"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

func (this User) TableName() string {
	return "tiktok_user"
}
