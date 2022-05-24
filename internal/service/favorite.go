package service

import (
	"simple-demo/internal/model"
)

type FavoriteRequest struct {
	UserId     int64  `form:"user_id"`
	VideoId    int64  `form:"video_id" `
	ActionType int    `form:"action_type" `
	Token      string `form:"token" `
}

type FavoriteListRequest struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}

type FavoriteListRespond struct {
	*Response
	VideoList []*model.Video `json:"video_list,omitempty"`
}

func (svc *Service) FavoriteAction(params *FavoriteRequest) (*Response, error) {
	username, _ := svc.ctx.Get("username")

	svc.cache.DeleteVideoFavoriteCount(params.UserId)
	svc.cache.DeleteVideoFavoriteList(username.(string))

	err := svc.dao.FavoriteAction(username.(string), params.VideoId, params.ActionType)
	if err != nil {
		return nil, err
	}
	return &Response{StatusCode: 0, StatusMsg: "success"}, nil
}

func (svc *Service) FavoriteList(params *FavoriteListRequest) (*FavoriteListRespond, error) {
	username, _ := svc.ctx.Get("username")
	videos, err := svc.cache.GetVideoFavoriteListInfo(username.(string)) //从缓存中读取
	if err == nil {
		respond := &FavoriteListRespond{
			Response: &Response{
				StatusCode: 0,
				StatusMsg:  "success",
			},
			VideoList: videos,
		}
		return respond, nil
	}
	videos, err = svc.dao.FavoriteList(username.(string)) //从数据库中读
	if err != nil {
		return nil, err
	}
	svc.cache.WriteVideoFavoriteListInfo(username.(string), videos) //写回缓存
	respond := &FavoriteListRespond{
		Response: &Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: videos,
	}
	return respond, nil
}
