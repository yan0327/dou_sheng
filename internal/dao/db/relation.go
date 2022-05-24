package db

import (
	"gorm.io/gorm"
)

type RelationDao interface {
	IsFollower(userId int64, toUserId int64) (bool, error)
	Create(userId int64, toUserId int64) error
	Delete(userId int64, toUserId int64) error
	FollowList(userId int64) ([]int64, error)
	FollowerList(userId int64) ([]int64, error)
}

type relations struct {
	db *gorm.DB
}

func MakeRelations(db *gorm.DB) *relations {
	return &relations{db}
}

const relationTableName = "tiktok_relation"

func (r relations) IsFollower(userId int64, toUserId int64) (bool, error) {
	var cnt int64
	r.db.Table(relationTableName).
		Where("user_id = ? AND follower_id = ?", toUserId, userId).
		Count(&cnt)
	return cnt > 0, nil
}

func (r relations) Create(userId int64, toUserId int64) error {
	if b, _ := r.IsFollower(userId, toUserId); b {
		return nil
	}
	return r.db.Table(relationTableName).Create(map[string]interface{}{
		"user_id":     toUserId,
		"follower_id": userId,
	}).Error
}

func (r relations) Delete(userId int64, toUserId int64) error {
	return r.db.Table(relationTableName).
		Where("user_id = ? AND follower_id = ?", toUserId, userId).
		Delete(map[string]interface{}{}).Error
}

func (r relations) FollowList(userId int64) ([]int64, error) {
	var uid []int64
	err := r.db.Table(relationTableName).
		Select("user_id").
		Where("follower_id = ?", userId).
		Find(&uid).Error
	if err != nil {
		return nil, err
	}
	return uid, nil
}

func (r relations) FollowerList(userId int64) ([]int64, error) {
	var uid []int64
	err := r.db.Table(relationTableName).
		Select("follower_id").
		Where("user_id = ?", userId).
		Find(&uid).Error
	if err != nil {
		return nil, err
	}
	return uid, nil
}
