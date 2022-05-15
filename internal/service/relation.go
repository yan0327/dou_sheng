package service

import "simple-demo/internal/model"

type RelationRequest struct {
	UserId     uint32 `form:"user_id" binding:"required"`
	Token      string `form:"token"`
	ToUserId   uint32 `form:"to_user_id" binding:"required"`
	ActionType uint8  `form:"action_type" binding:"required,oneof=1 2"`
}

//关注列表
type FollowListRequest struct {
	UserId uint32 `form:"user_id" binding:"required"`
}

type FollowListRepond struct {
	*Response
	UserList []model.User `json:"user_list,omitempty"`
}

//粉丝列表
type FollowerListRequest struct {
	UserId uint32 `form:"user_id" binding:"required"`
	// Token  string `form:"token" binding:"required"`
}

func (svc *Service) RelationAction(params *RelationRequest) (*Response, error) {
	err := svc.dao.RelationAction(params.UserId, params.ToUserId, params.ActionType)
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
