package service

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"io"
	"simple-demo/internal/dao/db"
	"simple-demo/internal/dao/store"
	"simple-demo/internal/model"
	"simple-demo/internal/pkg/errcode"
	"simple-demo/internal/pkg/global"
	"simple-demo/pkg/util"
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
	// 复制一份流
	var r2 bytes.Buffer
	r1 := io.TeeReader(r, &r2)

	// 上传视频
	id := uuid.New().String()
	if err := srv.store.Store(id, r1); err != nil {
		return errcode.ServerError.
			WithDetails(err.Error()).
			WithDetails("视频上传失败")
	}

	// 生成封面并上传
	coverReader, err := util.Frame4Video(&r2)
	if err != nil {
		srv.store.Delete(id)
		return errcode.ServerError.
			WithDetails(err.Error()).
			WithDetails("封面获取失败")
	}
	if err := srv.store.Store(id+"_cover", coverReader); err != nil {
		srv.store.Delete(id)
		return errcode.ServerError.
			WithDetails(err.Error()).
			WithDetails("封面上传失败")
	}

	// 落库
	if _, err := srv.db.Create(&model.Video{
		AuthorId: userId,
		PlayUrl:  fmt.Sprintf(global.AppSetting.UploadServerUrl, id),
		CoverUrl: fmt.Sprintf(global.AppSetting.UploadServerUrl, id+"_cover"),
		Title:    title,
	}); err != nil {
		srv.store.Delete(id)
		srv.store.Delete(id + "_cover")
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
