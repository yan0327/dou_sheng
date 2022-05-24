package service

import (
	"simple-demo/internal/model"
)

type UserLoginRequest struct {
	UserName string `form:"username"`
	PassWord string `form:"password"`
}

type UserInfoRequest struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}

type UserRegisterRequest struct {
	UserName string `form:"username"`
	PassWord string `form:"password"`
}

type UserLoginResponse struct {
	*Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	*Response
	User *model.User `json:"user"`
}

type UserRegisterRespond struct {
	*Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

func (svc *Service) UserInfo(params *UserInfoRequest) (*UserInfoResponse, error) {
	username, _ := svc.ctx.Get("username")
	user, err := svc.cache.GetUserStates(username.(string))
	if err == nil {
		return &UserInfoResponse{
			Response: &Response{StatusCode: 0, StatusMsg: "success"},
			User:     user,
		}, nil
	}
	user, err = svc.dao.GetUserInfo(username.(string))
	if err != nil {
		return nil, err
	}
	svc.cache.WriteUserState(user.ID, user.UserName, user.FollowCount, user.FollowerCount)

	return &UserInfoResponse{
		Response: &Response{StatusCode: 0, StatusMsg: "success"},
		User:     user,
	}, nil
}

func (svc *Service) UserRegister(params *UserRegisterRequest) (*UserRegisterRespond, error) {
	user, err := svc.dao.UserRegister(params.UserName, params.PassWord)
	if err != nil {
		return nil, err
	}
	svc.cache.WriteUserState(user.ID, user.UserName, user.FollowCount, user.FollowerCount)
	respond := &UserRegisterRespond{
		UserId:   user.ID,
		Response: &Response{StatusCode: 0, StatusMsg: "success"},
	}
	return respond, nil
}

func (svc *Service) UserLogin(params *UserLoginRequest) (*UserLoginResponse, error) {
	user, err := svc.dao.UserLogin(params.UserName, params.PassWord)
	if err != nil {
		return nil, err
	}
	svc.cache.WriteUserState(user.ID, user.UserName, user.FollowCount, user.FollowerCount)

	respond := &UserLoginResponse{
		UserId:   user.ID,
		Response: &Response{StatusCode: 0, StatusMsg: "success"},
	}
	return respond, nil
}
