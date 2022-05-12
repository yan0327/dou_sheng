package service

import (
	"simple-demo/internal/model"
	"simple-demo/pkg/app"
)

type UserLoginRequest struct {
	UserName string `form:"username" binding:"required,max=32"`
	PassWord string `form:"password" binding:"required,max=32"`
}

type UserInfoRequest struct {
	UserId uint32 `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

type UserRegisterRequest struct {
	UserName string `form:"username" binding:"required,max=32"`
	PassWord string `form:"password" binding:"required,max=32"`
}

type UserLoginResponse struct {
	Response
	UserId uint32 `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	Response
	User model.User `json:"user"`
}

type UserRegisterRespond struct {
	Response
	UserId uint32 `json:"user_id"`
	Token  string `json:"token"`
}

func (svc *Service) UserInfo(params *UserInfoRequest) (*UserInfoResponse, error) {
	respond := &UserInfoResponse{}
	token := params.Token
	_, err := app.ParseToken(token)
	if err != nil {
		return nil, err
	}
	respond.User, err = svc.dao.GetUserInfo(params.UserId)
	if err != nil {
		return nil, err
	}
	return respond, nil
}

func (svc *Service) UserRegister(params *UserRegisterRequest) (*UserRegisterRespond, error) {
	id, err := svc.dao.UserRegister(params.UserName, params.PassWord)
	if err != nil {
		return nil, err
	}
	respond := &UserRegisterRespond{
		UserId: id,
	}
	return respond, nil
}

func (svc *Service) UserLogin(params *UserLoginRequest) (*UserLoginResponse, error) {
	id, err := svc.dao.UserLogin(params.UserName, params.PassWord)
	if err != nil {
		return nil, err
	}
	respond := &UserLoginResponse{
		UserId: id,
	}
	return respond, nil
}
