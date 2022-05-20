package service

import (
	"errors"
	"simple-demo/internal/model"
	"simple-demo/pkg/app"
)

type UserLoginRequest struct {
	UserName string `form:"username"`
	PassWord string `form:"password"`
}

type UserInfoRequest struct {
	UserId uint32 `form:"user_id"`
	Token  string `form:"token"`
}

type UserRegisterRequest struct {
	UserName string `form:"username"`
	PassWord string `form:"password"`
}

type UserLoginResponse struct {
	*Response
	UserId uint32 `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	*Response
	User model.User `json:"user"`
}

type UserRegisterRespond struct {
	*Response
	UserId uint32 `json:"user_id"`
	Token  string `json:"token"`
}

func (svc *Service) UserInfo(params *UserInfoRequest) (*UserInfoResponse, error) {
	claims, err := app.ParseToken(params.Token)
	if err != nil {
		return nil, errors.New("token 不存在")
	}
	user, err := svc.dao.GetUserInfo(claims.AppKey)
	if err != nil {
		return nil, err
	}
	return &UserInfoResponse{
		Response: &Response{StatusCode: 0, StatusMsg: "success"},
		User:     user,
	}, nil
}

func (svc *Service) UserRegister(params *UserRegisterRequest) (*UserRegisterRespond, error) {
	id, err := svc.dao.UserRegister(params.UserName, params.PassWord)
	if err != nil {
		return nil, err
	}
	respond := &UserRegisterRespond{
		UserId:   id,
		Response: &Response{StatusCode: 0, StatusMsg: "success"},
	}
	return respond, nil
}

func (svc *Service) UserLogin(params *UserLoginRequest) (*UserLoginResponse, error) {
	id, err := svc.dao.UserLogin(params.UserName, params.PassWord)
	if err != nil {
		return nil, err
	}
	respond := &UserLoginResponse{
		UserId:   id,
		Response: &Response{StatusCode: 0, StatusMsg: "success"},
	}
	return respond, nil
}
