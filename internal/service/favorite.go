package service

import (
	"simple-demo/internal/model"
)

type FavoriteRequest struct {
	UserId     uint32 `form:"user_id" binding:"required"`
	VideoId    uint32 `form:"video_id" binding:"required"`
	ActionType int    `form:"action_type" binding:"required,oneof= 1 2"`
}

type FavoriteListRequest struct {
	UserId uint32 `form:"user_id"  binding:"required"`
}

type FavoriteListRespond struct {
	*Response
	VideoList []model.Video `json:"video_list,omitempty"`
}



func (svc *Service) FavoriteAction(params *FavoriteRequest) (*Response, error) {
	err := svc.dao.FavoriteAction(params.UserId, params.VideoId, params.ActionType)
	if err != nil {
		return nil, err
	}
	return &Response{StatusCode: 0, StatusMsg: "success"}, nil
}

func (svc *Service) FavoriteList(params *FavoriteListRequest) (*FavoriteListRespond, error) {
	videos, err := svc.dao.FavoriteList(params.UserId)
	if err != nil {
		return nil, err
	}
	respond := &FavoriteListRespond{
		Response: &Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: videos,
	}

	return respond, nil
}
