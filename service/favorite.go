package service

import (
	"simple-demo/global"
	"simple-demo/model"
)

type FavoriteRequest struct {
	UserId     uint32 `form:"user_id" binding:"required"`
	Token      string `form:"token" binding:"required"`
	VideoId    uint32 `form:"video_id" binding:"required"`
	ActionType int    `form:"action_type" binding:"required, oneof= 1 2"`
}

type FavoriteListRequest struct {
	UserId uint32 `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

func GetFavoriteNum(ID uint) (count int64, err error) {
	result := global.DBEngine.Model(&model.Like{}).Where("video_id = ? AND action_type = ?", ID, "1")
	count = result.RowsAffected
	err = result.Error
	return
}

func SetOrUpdateFavorite(UserId uint, VideoID uint, ActionType uint8) (like model.Like, err error) {
	like.UserId = UserId
	like.VideoId = VideoID
	like.ActionType = ActionType
	result := global.DBEngine.Model(&model.Like{}).Where("user_id = ? AND video_id = ?", UserId, VideoID).First(&like)
	if result.RowsAffected == 0 {
		err = global.DBEngine.Create(&like).Error
	} else {
		like.ActionType = ActionType
		err = global.DBEngine.Model(&like).Select("action_type").Updates(map[string]interface{}{"action_type": ActionType}).Error
	}
	return
}
func GetLikeVideoID(ID uint) (users []model.Like, err error) {
	err = global.DBEngine.Model(&model.Like{}).Where("user_id = ? AND action_type = ?", ID, "1").Find(&users).Error
	return
}
func FavoriteList(userID, ID uint) (replys []model.ReplyVideo, err error) {
	var videos []model.Like
	videos, err = GetLikeVideoID(ID)
	for i := 0; i < len(videos); i++ {
		reply := model.ReplyVideo{
			ID: videos[i].VideoId,
		}
		var video model.Video
		err = global.DBEngine.Model(&model.Video{}).Where("id = ?", videos[i].VideoId).First(&video).Error
		reply.PlayUrl = video.PlayUrl
		reply.CoverUrl = video.CoverUrl
		reply.FavoriteCount, err = GetFavoriteNum(reply.ID)
		if err != nil {
			return
		}
		reply.CommentCount, err = GetCommentsNum(reply.ID)
		if err != nil {
			return
		}
		reply.IsFavorite = IsLike(userID, ID)
		replys = append(replys, reply)
	}
	return
}

func IsLike(FollowID uint, VideoID uint) (Islike bool) {
	var real model.Like
	global.DBEngine.Model(&model.Like{}).Where("user_id = ? AND video_id = ?", FollowID, VideoID).Find(&real)
	if real.ActionType != 1 {
		return false
	}
	return true
}
