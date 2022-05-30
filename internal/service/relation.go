package service

import (
	"simple-demo/internal/model"
)

type RelationRequest struct {
	UserId     int64  `form:"user_id" `
	Token      string `form:"token"`
	ToUserId   int64  `form:"to_user_id" `
	ActionType uint8  `form:"action_type" `
}

//关注列表
type FollowListRequest struct {
	UserId int64  `form:"user_id" `
	Token  string `form:"token"`
}

type FollowListRepond struct {
	*Response
	UserList []*model.User `json:"user_list,omitempty"`
}

//粉丝列表
type FollowerListRequest struct {
	UserId int64 `form:"user_id"`
}

func (svc *Service) RelationAction(params *RelationRequest) (*Response, error) {
	username, _ := svc.ctx.Get("username")
	err := svc.dao.RelationAction(username.(string), params.ToUserId, params.ActionType)
	if err != nil {
		return nil, err
	}
	return &Response{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}

func (svc *Service) FollowList(params *FollowListRequest) (*FollowListRepond, error) {
	userList, err := svc.dao.FollowList(params.UserId)
	if err != nil {
		return nil, err
	}
	return &FollowListRepond{
		Response: &Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserList: userList,
	}, nil
}

func (svc *Service) FollowerList(params *FollowListRequest) (*FollowListRepond, error) {
	userList, err := svc.dao.FollowerList(params.UserId)
	if err != nil {
		return nil, err
	}
	return &FollowListRepond{
		Response: &Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserList: userList,
	}, nil
}
