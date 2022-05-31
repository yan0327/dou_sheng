package model

type User struct {
	Id             int64  `json:"id,omitempty"`
	Name           string `json:"name,omitempty" gorm:"column:username"`
	Password       string `json:"-"`
	FollowCount    int64  `json:"follow_count"`
	FollowerCount  int64  `json:"follower_count"`
	IsFollow       bool   `json:"is_follow"`
	TotalFavorited int64  `json:"total_favorited"`
	FavoriteCount  int64  `json:"favorite_count"`
}

func (this User) TableName() string {
	return "tiktok_user"
}
