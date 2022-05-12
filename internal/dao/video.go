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
