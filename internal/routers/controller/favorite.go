package api

import (
	"github.com/gin-gonic/gin"
	"simple-demo/internal/dao"
	"simple-demo/internal/pkg/api"
	"simple-demo/internal/pkg/errcode"
	"simple-demo/internal/pkg/global"
	"simple-demo/internal/service"
	"simple-demo/pkg/app"
)

const (
	ActionLike = iota + 1
	ActionCancelLike
)

type FavoriteActionRequest struct {
	VideoId    int64 `form:"video_id" binding:"required"`
	ActionType int   `form:"action_type" binding:"required"`
}

type FavoriteListRequest struct {
	UserId int64 `form:"user_id" binding:"required"`
}

type FavoriteController struct {
	fsrv service.FavoriteSrv
}

func MakeFavoriteController(f *dao.DaoFactory) *FavoriteController {
	return &FavoriteController{
		fsrv: service.MakeFavoriteSrv(f.Favorite(), f.Video()),
	}
}

func (f *FavoriteController) FavoriteAction(c *gin.Context) {
	uid, exist := c.Get(service.UserId)
	if !exist {
		api.RespWithErr(c, errcode.ServerError.WithDetails("获取当前用户信息失败"))
		global.Logger.Error(c, "获取当前用户信息失败")
		return
	}
	var req FavoriteActionRequest
	if valid, err := app.BindAndValid(c, &req); !valid || err != nil {
		api.RespWithErr(c, errcode.InvalidParams)
		return
	}

	if req.ActionType == ActionLike || req.ActionType == ActionCancelLike {
		err := f.fsrv.Like(uid.(int64), req.VideoId)
		if err != nil {
			global.Logger.Error(c, err.Details())
			api.RespWithErr(c, errcode.ServerError.WithDetails(err.Error()))
			return
		}
		api.RespOK(c)
	} else {
		api.RespWithErr(c, errcode.InvalidParams)
	}

}

func (f *FavoriteController) FavoriteList(c *gin.Context) {
	var req FavoriteListRequest
	if valid, err := app.BindAndValid(c, &req); !valid || err != nil {
		api.RespWithErr(c, errcode.InvalidParams)
		return
	}

	list, err := f.fsrv.ListByUser(req.UserId)
	if err != nil {
		global.Logger.Error(c, err.Details())
		api.RespWithErr(c, errcode.ServerError.WithDetails(err.Error()))
		return
	}

	api.RespWithData(c, VideoListResponse{
		VideoList: list,
	})
}
