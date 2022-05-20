package model

import "github.com/jinzhu/gorm"

type Relation struct {
	Id         uint32 `gorm:"column:id"`
	FollowerId uint32 `gorm:"column:follower_id"`
	UserId     uint32 `gorm:"column:user_id"`
	ActionType uint8  `gorm:"column:action_type"`
	UserName   string `gorm:"-"`
}

func (this Relation) TableName() string {
	return "tiktok_relation"
}

func (this Relation) RelationAction(db *gorm.DB) error {
	db.Table("tiktok_user").Select("id").Where("username = ?", this.UserName).Find(&this.FollowerId)
	relation := Relation{}
	err := db.Where("user_id = ? AND follower_id = ?", this.UserId, this.FollowerId).First(&relation).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err == gorm.ErrRecordNotFound {
		return db.Create(&this).Error
	} else {
		return db.Table("tiktok_relation").Where("user_id = ? AND follower_id = ?", this.UserId, this.FollowerId).Update("action_type", this.ActionType).Error
	}
}

func (this Relation) FollowList(db *gorm.DB) ([]User, error) {
	// db.Table("tiktok_user").Select("id").Where("username = ?", this.UserName).Find(&this.Id)
	users := []User{}
	db.Table("tiktok_user").Select("tiktok_user.id,tiktok_user.username").Joins("inner join tiktok_relation on tiktok_relation.user_id = tiktok_user.id").Where("tiktok_relation.follower_id = ?", this.Id).Scan(&users)
	for i := 0; i < len(users); i++ {
		var cnt uint32
		err := db.Table("tiktok_relation").Where("user_id = ? AND action_type = ?", users[i].ID, 1).Count(&cnt).Error
		if err != nil {
			return nil, err
		}
		users[i].FollowCount = cnt
		err = db.Table("tiktok_relation").Where("follower_id = ? AND action_type = ?", users[i].ID, 1).Count(&cnt).Error
		if err != nil {
			return nil, err
		}
		users[i].FollowerCount = cnt
	}

	return users, nil
}

func (this Relation) FollowerList(db *gorm.DB) ([]User, error) {
	users := []User{}
	db.Table("tiktok_user").Select("tiktok_user.id,tiktok_user.username").Joins("inner join tiktok_relation on tiktok_relation.follower_id = tiktok_user.id").Where("tiktok_relation.user_id = ?", this.Id).Scan(&users)
	for i := 0; i < len(users); i++ {
		var cnt uint32
		err := db.Table("tiktok_relation").Where("user_id = ? AND action_type = ?", users[i].ID, 1).Count(&cnt).Error
		if err != nil {
			return nil, err
		}
		users[i].FollowCount = cnt
		err = db.Table("tiktok_relation").Where("follower_id = ? AND action_type = ?", users[i].ID, 1).Count(&cnt).Error
		if err != nil {
			return nil, err
		}
		users[i].FollowerCount = cnt
	}

	return users, nil
}
