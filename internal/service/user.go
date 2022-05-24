package service

import (
	j "github.com/golang-jwt/jwt/v4"
	"simple-demo/internal/dao/db"
	"simple-demo/internal/model"
	"simple-demo/internal/pkg/errcode"
	"simple-demo/internal/pkg/jwt"
)

const UserId = "user_id"

type UserSrv interface {
	Login(username string, password string) (*model.User, string, *errcode.Error)
	Register(username string, password string) (*model.User, string, *errcode.Error)
	GetById(userId int64) (*model.User, *errcode.Error)
	Follow(userId int64, toUserId int64) *errcode.Error
	FollowList(userId int64) ([]*model.User, *errcode.Error)
	FollowerList(userId int64) ([]*model.User, *errcode.Error)
}

type userService struct {
	udb db.UserDao
	rdb db.RelationDao
}

func MakeUserSrv(udb db.UserDao, rdb db.RelationDao) UserSrv {
	return &userService{udb, rdb}
}

func (srv *userService) Login(username string, password string) (*model.User, string, *errcode.Error) {
	u, err := srv.udb.FindByName(username)
	if err != nil {
		return nil, "", errcode.ServerError.WithDetails(err.Error())
	}
	if u == nil {
		return nil, "", errcode.ErrorUserNotExistFail
	}

	token, err := jwt.GenerateJWT(j.MapClaims{
		UserId: u.Id,
	})
	if err != nil {
		return nil, "", errcode.UnauthorizedTokenGenerate.WithDetails(err.Error())
	}
	return u, token, nil
}

func (srv *userService) Register(username string, password string) (*model.User, string, *errcode.Error) {
	// TODO 密码不明文存储
	u, err := srv.udb.FindByName(username)
	if err != nil {
		return nil, "", errcode.ServerError.WithDetails(err.Error())
	}
	if u != nil {
		return nil, "", errcode.ErrorUserExistFail
	}

	u, err = srv.udb.Create(&model.User{
		Name:     username,
		Password: password,
	})
	if err != nil {
		return nil, "", errcode.ServerError.WithDetails(err.Error())
	}

	token, err := jwt.GenerateJWT(j.MapClaims{
		UserId: u.Id,
	})
	if err != nil {
		return nil, "", errcode.UnauthorizedTokenGenerate.WithDetails(err.Error())
	}
	return u, token, nil
}

func (srv *userService) GetById(userId int64) (*model.User, *errcode.Error) {
	user, err := srv.udb.FindById(userId)
	if err != nil {
		return nil, errcode.ServerError.WithDetails(err.Error())
	}
	return user, nil
}

func (srv *userService) Follow(userId int64, toUserId int64) *errcode.Error {
	var err error
	if b, _ := srv.rdb.IsFollower(userId, toUserId); b {
		err = srv.rdb.Delete(userId, toUserId)
	} else {
		err = srv.rdb.Create(userId, toUserId)
	}
	if err != nil {
		return errcode.ServerError.WithDetails(err.Error())
	}
	return nil
}

func (srv *userService) FollowList(userId int64) ([]*model.User, *errcode.Error) {
	list, err := srv.rdb.FollowList(userId)
	if err != nil {
		return nil, errcode.ServerError.WithDetails(err.Error())
	}
	u, err := srv.udb.FindByIds(list)
	if err != nil {
		return nil, errcode.ServerError.WithDetails(err.Error())
	}
	return u, nil
}

func (srv *userService) FollowerList(userId int64) ([]*model.User, *errcode.Error) {
	list, err := srv.rdb.FollowerList(userId)
	if err != nil {
		return nil, errcode.ServerError.WithDetails(err.Error())
	}
	u, err := srv.udb.FindByIds(list)
	if err != nil {
		return nil, errcode.ServerError.WithDetails(err.Error())
	}
	return u, nil
}
