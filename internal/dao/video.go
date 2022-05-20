package dao

import (
	"simple-demo/internal/model"
)

func (d *Dao) PublishList(userId uint32) ([]model.Video, error) {
	video := model.Video{}
	video.User = &model.User{
		ID: userId,
	}
	return video.PublishList(d.engine)
}

func (d *Dao) ReverseFeed(lastTime int64) ([]model.Video, error) {
	video := model.Video{LastTime: lastTime}
	return video.ReverseFeed(d.engine, lastTime)
}

func (d *Dao) Publish(username string, playUrl string, coverUrl string) error {
	videoPush := model.VideoPush{
		UserName: username,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
	}
	return videoPush.Publish(d.engine)
}

func (d *Dao) FavoriteAction(username string, videoId uint32, actionType int) error {
	favorite := model.Favorite{
		UserName:   username,
		VideoId:    videoId,
		ActionType: actionType,
	}
	return favorite.FavoriteAction(d.engine)
}

func (d *Dao) FavoriteList(username string) ([]model.Video, error) {
	favorite := model.Favorite{
		UserName: username,
	}
	return favorite.FavoriteList(d.engine)
}
