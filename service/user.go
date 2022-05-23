package service

import (
	"simple-demo/global"
	"simple-demo/model"
)

type UserLoginRequest struct {
	UserName string `form:"username" binding:"required, max=32"`
	PassWord string `form:"password" binding:"required, max=32"`
}

type UserInfoRequest struct {
	UserId uint32 `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

func FindAuthor(ID uint) (author model.User, err error) {
	err = global.DBEngine.Model(&model.User{}).Where("id = ?", ID).First(&author).Error
	return
}

func FindUser(username string) (user model.User, err error) {
	err = global.DBEngine.Where("username =  ?", username).First(&user).Error
	return
}
func FindReplyUser(UserID uint, GoalID uint) (reply model.ReplyUser, err error) {
	var user model.User
	err = global.DBEngine.Model(&model.User{}).Where("id = ?", GoalID).Find(&user).Error
	if err != nil {
		return
	}
	reply.ID = user.ID
	reply.Username = user.Username
	reply.FollowCount, err = GetFollowNum(user.ID)
	if err != nil {
		return
	}
	reply.FollowerCount, err = GetFollowerNum(user.ID)
	if err != nil {
		return
	}
	reply.IsFollow = IsFollow(UserID, GoalID)
	return
}
