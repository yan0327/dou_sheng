package model

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty" gorm:"column:username"`
	Password      string `json:"-"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

func (this User) TableName() string {
	return "tiktok_user"
}
