package service

import (
	"github.com/google/uuid"
	"io"
	"simple-demo/internal/dao/db"
	"simple-demo/internal/dao/store"
	"simple-demo/internal/model"
	"simple-demo/internal/pkg/errcode"
)

type VideoSrv interface {
	DataStream(name string) (io.Reader, *errcode.Error)
	Feed(latestTime int64) ([]*model.Video, *errcode.Error)
	FindByUser(userId int64) ([]*model.Video, *errcode.Error)
	Publish(userId int64, title string, r io.Reader) *errcode.Error
}

type videoService struct {
	store store.Store
	db    db.VideoDao
}

func MakeVideoSrv(s store.Store, db db.VideoDao) VideoSrv {
	return &videoService{s, db}
}

func (srv *videoService) Feed(latestTime int64) ([]*model.Video, *errcode.Error) {
	v, e := srv.db.FindByTime(latestTime)
	if e != nil {
		return nil, errcode.ServerError.WithDetails(e.Error())
	}
	return v, nil
}

func (srv *videoService) FindByUser(userId int64) ([]*model.Video, *errcode.Error) {
	v, e := srv.db.FindByUser(userId)
	if e != nil {
		return nil, errcode.ServerError.WithDetails(e.Error())
	}
	return v, nil
}

func (srv *videoService) Publish(userId int64, title string, r io.Reader) *errcode.Error {
	id := uuid.New().String()
	if err := srv.store.Store(id, r); err != nil {
		return errcode.ServerError.WithDetails(err.Error())
	}
	if _, err := srv.db.Create(&model.Video{
		AuthorId: userId,
		PlayUrl:  "http://10.0.2.2:8080/douyin/video/" + id, //TODO 播放链接处理
		Title:    title,
	}); err != nil {
		srv.store.Delete(id)
		return errcode.ServerError.WithDetails(err.Error())
	}

	return nil
}

func (srv *videoService) DataStream(name string) (io.Reader, *errcode.Error) {
	r, err := srv.store.Get(name)
	if err != nil {
		return nil, errcode.ServerError.WithDetails(err.Error())
	}
	return r, nil
}
