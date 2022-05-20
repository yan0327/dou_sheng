package service

import (
	"errors"
	"simple-demo/internal/model"
	"simple-demo/pkg/app"
)

type FavoriteRequest struct {
	UserId     uint32 `form:"user_id"`
	VideoId    uint32 `form:"video_id" `
	ActionType int    `form:"action_type" `
	Token      string `form:"token" `
}

type FavoriteListRequest struct {
	UserId uint32 `form:"user_id"`
	Token  string `form:"token"`
}

type FavoriteListRespond struct {
	*Response
	VideoList []model.Video `json:"video_list,omitempty"`
}

func (svc *Service) FavoriteAction(params *FavoriteRequest) (*Response, error) {
	claims, err := app.ParseToken(params.Token)
	if err != nil {
		return nil, errors.New("token 不存在")
	}
	err = svc.dao.FavoriteAction(claims.AppKey, params.VideoId, params.ActionType)
	if err != nil {
		return nil, err
	}
	return &Response{StatusCode: 0, StatusMsg: "success"}, nil
}

func (svc *Service) FavoriteList(params *FavoriteListRequest) (*FavoriteListRespond, error) {
	claims, err := app.ParseToken(params.Token)
	if err != nil {
		return nil, errors.New("token 不存在")
	}
	videos, err := svc.dao.FavoriteList(claims.AppKey)
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
