package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"reflect"
	"simple-demo/internal/pkg/errcode"
)

const (
	StatusCode = "status_code"
	StatusMsg  = "status_msg"
)

func Resp(ctx *gin.Context, e *errcode.Error, data interface{}) {
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
	var ret gin.H
	if data == nil {
		ret = gin.H{}
	} else {
		// 转成map
		v := reflect.TypeOf(data)
		if v.Kind() == reflect.Map {
			ret = data.(gin.H)
		} else {
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			if v.Kind() == reflect.Struct {
				// TODO 优化性能
				bytes, _ := json.Marshal(&data)
				json.Unmarshal(bytes, &ret)
			}
		}
	}
	ret[StatusCode] = errcode.Success.Code()
	ret[StatusMsg] = errcode.Success.Msg()
	ctx.JSON(errcode.Success.HTTPStatus(), ret)
}

func RespWithErr(ctx *gin.Context, e *errcode.Error) {
	Resp(ctx, e, nil)
}

func RespWithData(ctx *gin.Context, data interface{}) {
	Resp(ctx, errcode.Success, data)
}

func RespOK(ctx *gin.Context) {
	Resp(ctx, errcode.Success, nil)
}
