package api

import (
	"github.com/gin-gonic/gin"
	"simple-demo/internal/dao"
	"simple-demo/internal/model"
	"simple-demo/internal/pkg/api"
	"simple-demo/internal/pkg/errcode"
	"simple-demo/internal/pkg/global"
	"simple-demo/internal/service"
	"simple-demo/pkg/app"
	"strconv"
)

const (
	ActionPublishComment = iota + 1
	ActionDeleteComment
)

type CommentActionRequest struct {
	Token       string `form:"token" binding:"required"`
	VideoId     int64  `form:"video_id" binding:"required"`
	ActionType  int    `form:"action_type" binding:"required"`
	CommentText string `form:"comment_text"`
	CommentId   int64  `form:"comment_id"`
}

type CommentListResponse struct {
	CommentList []*model.Comment `json:"comment_list"`
}

type CommentActionResponse struct {
	Comment *model.Comment `json:"comment"`
}

type CommentController struct {
	csrv service.CommentSrv
}

func MakeCommentController(f *dao.DaoFactory) *CommentController {
	return &CommentController{
		csrv: service.MakeCommentSrv(f.Comment()),
	}
}

func (cm *CommentController) CommentAction(c *gin.Context) {
	uid, exist := c.Get(service.UserId)
	if !exist {
		api.RespWithErr(c, errcode.ServerError.WithDetails("获取当前用户信息失败"))
		global.Logger.Error(c, "获取当前用户信息失败")
		return
	}
	var req CommentActionRequest
	if valid, err := app.BindAndValid(c, &req); !valid || err != nil {
		api.RespWithErr(c, errcode.InvalidParams)
		return
	}

	if req.ActionType == ActionPublishComment {
		comment, err := cm.csrv.Publish(req.VideoId, uid.(int64), req.CommentText)
		if err != nil {
			global.Logger.Errorf(c, "新增评论错误: %v\n", err.Details())
			api.RespWithErr(c, err)
			return
		}
		api.RespWithData(c, CommentActionResponse{
			comment,
		})
	} else if req.ActionType == ActionDeleteComment {
		err := cm.csrv.Delete(req.CommentId, uid.(int64))
		if err != nil {
			global.Logger.Errorf(c, "删除评论错误: %v\n", err.Details())
			api.RespWithErr(c, err)
			return
		}
		api.RespOK(c)
	} else {
		api.RespWithErr(c, errcode.InvalidParams)
	}
}

func (cm *CommentController) CommentList(c *gin.Context) {
	videoId, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		api.RespWithErr(c, errcode.InvalidParams)
		return
	}

	if list, err := cm.csrv.List(int64(videoId)); err != nil {
		api.RespWithErr(c, errcode.ServerError.WithDetails(err.Error()))
	} else {
		api.RespWithData(c, CommentListResponse{
			CommentList: list,
		})
	}
}
