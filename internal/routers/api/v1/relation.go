package v1

import (
	"net/http"
	"simple-demo/global"
	"simple-demo/internal/model"
	"simple-demo/internal/service"
	"simple-demo/pkg/app"
	"simple-demo/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	Response
	UserList []model.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	params := service.RelationRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	respond, err := svc.RelationAction(&params)
	if err != nil {
		global.Logger.Errorf(c, "svc.RelationAction err: %v", err)
		response.ToErrorResponse(errcode.UserGetInfoError)
		return
	}

	c.JSON(http.StatusOK, respond)
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	params := service.FollowListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	params.Token = c.Query("token")
	svc := service.New(c.Request.Context())
	respond, err := svc.FollowList(&params)
	if err != nil {
		global.Logger.Errorf(c, "svc.FollowList err: %v", err)
		response.ToErrorResponse(errcode.UserGetInfoError)
		return
	}

	c.JSON(http.StatusOK, respond)
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	params := service.FollowListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	respond, err := svc.FollowerList(&params)
	if err != nil {
		global.Logger.Errorf(c, "svc.FollowerList err: %v", err)
		response.ToErrorResponse(errcode.UserGetInfoError)
		return
	}

	c.JSON(http.StatusOK, respond)
}
