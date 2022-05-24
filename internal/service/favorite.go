package service

import (
	"simple-demo/internal/dao/db"
	"simple-demo/internal/model"
	"simple-demo/internal/pkg/errcode"
)

type FavoriteSrv interface {
	Like(userId int64, videoId int64) *errcode.Error
	ListByUser(userId int64) ([]*model.Video, *errcode.Error)
}

type favoriteService struct {
	fdb db.FavoriteDao
	vdb db.VideoDao
}

func MakeFavoriteSrv(fdb db.FavoriteDao, vdb db.VideoDao) FavoriteSrv {
	return &favoriteService{fdb, vdb}
}

func (f *favoriteService) Like(userId int64, videoId int64) *errcode.Error {
	var err error
	if b, _ := f.fdb.IsFavorite(userId, videoId); b {
		err = f.fdb.Delete(userId, videoId)
	} else {
		err = f.fdb.Create(userId, videoId)
	}

	if err != nil {
		return errcode.ServerError.WithDetails(err.Error())
	}
	return nil
}

func (f *favoriteService) ListByUser(userId int64) ([]*model.Video, *errcode.Error) {
	v, err := f.vdb.FindFavoriteByUser(userId)
	if err != nil {
		return nil, errcode.ServerError.WithDetails(err.Error())
	}
	return v, nil
}
