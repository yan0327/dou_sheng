package service

import (
	"simple-demo/global"

	"simple-demo/model"
)

type RelationRequest struct {
	UserId     uint32 `form:"user_id" binding:"required"`
	Token      string `form:"token" binding:"required"`
	ToUserId   uint32 `form:"to_user_id" binding:"required"`
	ActionType uint8  `form:"action_type" binding:"required, oneof= 1 2"`
}

//关注列表
type FollowListRequest struct {
	UserId uint32 `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

//粉丝列表
type FollowerListRequest struct {
	UserId uint32 `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

func GetFollowNum(ID uint) (count int64, err error) {
	var users []model.Realtion
	err = global.DBEngine.Model(&model.Realtion{}).Where("user_id = ? AND is_effective = ?", ID, "1").Find(&users).Error
	return int64(len(users)), err
}
func GetFollowID(ID uint) (users []model.Realtion, err error) {
	err = global.DBEngine.Model(&model.Realtion{}).Where("user_id = ? AND is_effective = ?", ID, "1").Find(&users).Error
	return
}
func GetFollowerNum(ID uint) (count int64, err error) {
	var users []model.Realtion
	err = global.DBEngine.Model(&model.Realtion{}).Where("follower_id = ? AND is_effective = ?", ID, "1").Find(&users).Error
	return int64(len(users)), err
}
func GetFollowerID(ID uint) (users []model.Realtion, err error) {

	err = global.DBEngine.Model(&model.Realtion{}).Where("follower_id = ? AND is_effective = ?", ID, "1").Find(&users).Error
	return
}
func IsFollow(FollowID uint, FollowerID uint) (Isfollow bool, err error) {
	var real model.Realtion
	global.DBEngine.Model(&model.Realtion{}).Where("user_id = ? AND follower_id = ?", FollowID, FollowerID).Find(&real)
	if real.IsEffective != 1 {
		return false, err
	}
	return true, err
}

func SetOrUpdateRelation(FollowID uint, FollowerID uint, IsEffiective uint8) (real model.Realtion, err error) {
	real.UserId = FollowID
	real.FollowerId = FollowerID
	real.IsEffective = IsEffiective
	result := global.DBEngine.Model(&model.Realtion{}).Where("user_id = ? AND follower_id = ?", FollowID, FollowerID).First(&real)
	if result.RowsAffected == 0 {
		err = global.DBEngine.Create(&real).Error
	} else {
		real.IsEffective = IsEffiective
		err = global.DBEngine.Model(&real).Select("is_effective").Updates(map[string]interface{}{"is_effective": IsEffiective}).Error
	}
	return
}
