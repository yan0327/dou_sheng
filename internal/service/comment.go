package service

import (
	"simple-demo/internal/model"
)

type CommentRequest struct {
	UserId      int64  `form:"user_id"`
	VideoId     int64  `form:"video_id"`
	ActionType  uint8  `form:"action_type"`
	CommentText string `form:"comment_text"`
	CommentId   int64  `form:"comment_id"`
	Token       string `form:"token"`
}

type CommentListRequest struct {
	UserId  int64  `form:"user_id"`
	VideoId int64  `form:"video_id"`
	Token   string `form:"token"`
}

type CommentListResponse struct {
	*Response
	CommentList []model.Comment `json:"comment_list,omitempty"`
}

func (svc *Service) CreateComment(params *CommentRequest) (*CommentListResponse, error) {
	if params.ActionType == 2 && params.CommentId == 0 {
		return &CommentListResponse{Response: &Response{StatusCode: 500, StatusMsg: "not comment"}}, nil
	}
	if params.ActionType == 1 && params.CommentText == "" {
		return &CommentListResponse{Response: &Response{StatusCode: 500, StatusMsg: "not comment"}}, nil
	}

	username, _ := svc.ctx.Get("username")
	repond, err := svc.dao.CreateComment(username.(string), params.VideoId, params.ActionType, params.CommentText, params.CommentId)
	if err != nil {
		return nil, err
	}
	return &CommentListResponse{Response: &Response{StatusCode: 0, StatusMsg: "not comment"},
		CommentList: []model.Comment{repond}}, nil
}

func (svc *Service) GetCommentList(params *CommentListRequest) (*CommentListResponse, error) {
	username, _ := svc.ctx.Get("username")
	comments, err := svc.dao.GetCommentList(username.(string), params.VideoId)
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
