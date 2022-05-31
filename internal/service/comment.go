package service

import (
	"simple-demo/internal/dao/db"
	"simple-demo/internal/model"
	"simple-demo/internal/pkg/errcode"
)

type CommentSrv interface {
	Publish(videoId int64, userId int64, comment string) (*model.Comment, *errcode.Error)
	Delete(commentId int64, userId int64) *errcode.Error
	List(videoId int64) ([]*model.Comment, *errcode.Error)
}

type CommentService struct {
	db db.CommentDao
}

func MakeCommentSrv(db db.CommentDao) CommentSrv {
	return &CommentService{db}
}

func (srv *CommentService) Publish(videoId int64, userId int64, comment string) (*model.Comment, *errcode.Error) {
	c, err := srv.db.Create(&model.Comment{
		UserId:  userId,
		Content: comment,
		VideoId: videoId,
	})
	if err != nil {
		return nil, errcode.ServerError.WithDetails(err.Error())
	}
	return c, nil
}

func (srv *CommentService) Delete(commentId int64, userId int64) *errcode.Error {
	c, err := srv.db.FindByIdUser(commentId, userId)
	if err != nil {
		return errcode.ServerError.WithDetails(err.Error())
	}
	if c == nil {
		return errcode.PermissionDenied
	}
	err = srv.db.Delete(commentId)
	if err != nil {
		return errcode.ServerError.WithDetails(err.Error())
	}
	return nil
}

func (srv *CommentService) List(videoId int64) ([]*model.Comment, *errcode.Error) {
	c, err := srv.db.FindByVideo(videoId)
	if err != nil {
		return nil, errcode.ServerError.WithDetails(err.Error())
	}
	return c, nil
}
