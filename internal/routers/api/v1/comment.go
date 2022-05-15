package v1

import (
	"net/http"
	"simple-demo/global"
	"simple-demo/internal/service"
	"simple-demo/pkg/app"
	"simple-demo/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	params := service.CommentRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	respond, err := svc.CreateComment(&params)
	if err != nil {
		global.Logger.Errorf(c, "svc.UserInfo err: %v", err)
		response.ToErrorResponse(errcode.UserGetInfoError)
		return
	}

	c.JSON(http.StatusOK, respond)
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	params := service.CommentListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	respond, err := svc.GetCommentList(&params)
	if err != nil {
		global.Logger.Errorf(c, "svc.CommentList err: %v", err)
		response.ToErrorResponse(errcode.UserGetInfoError)
		return
	}
	c.JSON(http.StatusOK, respond)
}
