package service

import (
	"errors"
	"simple-demo/internal/model"
	"simple-demo/pkg/app"
)

type CommentRequest struct {
	UserId      uint32 `form:"user_id"`
	VideoId     uint32 `form:"video_id"`
	ActionType  uint8  `form:"action_type"`
	CommentText string `form:"comment_text"`
	CommentId   uint32 `form:"comment_id"`
	Token       string `form:"token"`
}

type CommentListRequest struct {
	UserId  uint32 `form:"user_id"`
	VideoId uint32 `form:"video_id"`
	Token   string `form:"token"`
}

type CommentListResponse struct {
	*Response
	CommentList []model.Comment `json:"comment_list,omitempty"`
}

func (svc *Service) CreateComment(params *CommentRequest) (*Response, error) {
	if params.ActionType == 2 && params.CommentId == 0 {
		return &Response{StatusCode: 200, StatusMsg: "not comment"}, nil
	}
	if params.ActionType == 1 && params.CommentText == "" {
		return &Response{StatusCode: 200, StatusMsg: "not comment"}, nil
	}
	claims, err := app.ParseToken(params.Token)
	if err != nil {
		return nil, errors.New("token 不存在")
	}
	err = svc.dao.CreateComment(claims.AppKey, params.VideoId, params.ActionType, params.CommentText, params.CommentId)
	if err != nil {
		return nil, err
	}
	return &Response{StatusCode: 200, StatusMsg: "success"}, nil
}

func (svc *Service) GetCommentList(params *CommentListRequest) (*CommentListResponse, error) {
	claims, err := app.ParseToken(params.Token)
	if err != nil {
		return nil, errors.New("token 不存在")
	}

	comments, err := svc.dao.GetCommentList(claims.AppKey, params.VideoId)
	if err != nil {
		return nil, err
	}
	return &CommentListResponse{
		Response: &Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		CommentList: comments,
	}, nil
}
