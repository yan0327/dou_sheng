package api

import (
	"github.com/agiledragon/gomonkey/v2"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"reflect"
	"simple-demo/internal/pkg/errcode"
	"testing"
)

func TestResp(t *testing.T) {
	var c *gin.Context
	type jsonResType struct {
		code int
		obj  interface{}
	}
	var jsonRes jsonResType
	p := gomonkey.ApplyMethod(reflect.TypeOf(c), "JSON", func(c *gin.Context, code int, obj interface{}) {
		jsonRes = jsonResType{
			code: code,
			obj:  obj,
		}
	})
	defer p.Reset()

	type args struct {
		ctx  *gin.Context
		e    *errcode.Error
		data gin.H
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"成功-无数据",
			args{
				ctx:  &gin.Context{},
				e:    errcode.Success,
				data: gin.H{},
			},
		},
		{
			"成功-有数据",
			args{
				ctx: &gin.Context{},
				e:   errcode.Success,
				data: gin.H{
					"abc": 1,
					"def": 2,
				},
			},
		},
		{
			"成功-Error nil",
			args{
				ctx:  &gin.Context{},
				e:    nil,
				data: nil,
			},
		},
		{
			"失败",
			args{
				ctx:  &gin.Context{},
				e:    errcode.ServerError,
				data: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Resp(tt.args.ctx, tt.args.e, tt.args.data)
			inData := tt.args.data
			inErr := tt.args.e
			outData := jsonRes.obj.(gin.H)

			if inErr == nil {
				inErr = errcode.Success
			}
			if inData == nil {
				inData = gin.H{}
			}
			assert.Equal(t, jsonRes.code, inErr.HTTPStatus())
			assert.Equal(t, inErr.Code(), outData[StatusCode])
			assert.Equal(t, inErr.Msg(), outData[StatusMsg])
			if inErr.Code() == errcode.Success.Code() {
				delete(outData, StatusCode)
				delete(outData, StatusMsg)
				assert.Equal(t, inData, outData)
			}
		})
	}
}
