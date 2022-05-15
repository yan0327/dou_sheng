package dao

import (
	"simple-demo/internal/model"
)

func (d *Dao) ReverseFeed(lastTime int64) ([]model.Video, error) {
	vedio := model.Video{LastTime: lastTime}
	return vedio.ReverseFeed(d.engine, lastTime)
}

func (d *Dao) Publish(username string, playUrl string, coverUrl string) error {
	videoPush := model.VideoPush{
		UserName: username,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
	}
	return videoPush.Publish(d.engine)
}

func (d *Dao) FavoriteAction(userId uint32, videoId uint32, actionType int) error {
	favorite := model.Favorite{
		UserId:     userId,
		VideoId:    videoId,
		ActionType: actionType,
	}
	return favorite.FavoriteAction(d.engine)
}

func (d *Dao) FavoriteList(userId uint32) ([]model.Video, error) {
	favorite := model.Favorite{
		UserId: userId,
	}
	return favorite.FavoriteList(d.engine)
}
