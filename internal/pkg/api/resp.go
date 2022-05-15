package api

import (
	"github.com/gin-gonic/gin"
	"simple-demo/internal/pkg/errcode"
)

const (
	StatusCode = "status_code"
	StatusMsg  = "status_msg"
)

func Resp(ctx *gin.Context, e *errcode.Error, data gin.H) {
	if e == nil {
		e = errcode.Success
	}
	if e.Code() != errcode.Success.Code() {
		ctx.JSON(e.HTTPStatus(), gin.H{
			StatusCode: e.Code(),
			StatusMsg:  e.Msg(),
		})
		return
	}
	if data == nil {
		data = gin.H{}
	}
	data[StatusCode] = errcode.Success.Code()
	data[StatusMsg] = errcode.Success.Msg()
	ctx.JSON(errcode.Success.HTTPStatus(), data)
}

func RespWithErr(ctx *gin.Context, e *errcode.Error) {
	Resp(ctx, e, nil)
}

func RespWithData(ctx *gin.Context, data gin.H) {
	Resp(ctx, errcode.Success, data)
}

func RespOK(ctx *gin.Context) {
	Resp(ctx, errcode.Success, nil)
}
