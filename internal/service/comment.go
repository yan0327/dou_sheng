package service

import (
	"simple-demo/internal/model"
)

type CommentRequest struct {
	UserId      uint32 `form:"user_id" binding:"required"`
	Token       string `form:"token"`
	VideoId     uint32 `form:"video_id" binding:"required"`
	ActionType  uint8  `form:"action_type" binding:"required"`
	CommentText string `form:"comment_text"`
	CommentId   uint32 `form:"comment_id"`
}

type CommentListRequest struct {
	UserId  uint32 `form:"user_id" binding:"required"`
	Token   string `form:"token"`
	VideoId uint32 `form:"video_id" binding:"required"`
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
	err := svc.dao.CreateComment(params.UserId, params.VideoId, params.ActionType, params.CommentText, params.CommentId)
	if err != nil {
		return nil, err
	}
	return &Response{StatusCode: 200, StatusMsg: "success"}, nil
}

func (svc *Service) GetCommentList(params *CommentListRequest) (*CommentListResponse, error) {
	comments, err := svc.dao.GetCommentList(params.UserId, params.VideoId)
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
