package service

import (
	"simple-demo/global"
	"simple-demo/model"
)

type FeedRequest struct {
	LatestTime uint32 `form:"latest_time"`
}

func FeedVideos() (videos []model.Video, err error) {
	err = global.DBEngine.Model(&model.Video{}).Order("create_time desc").Limit(30).Find(&videos).Error
	return
}
