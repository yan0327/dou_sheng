package api

import (
	"encoding/json"
	"fmt"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"reflect"
	"simple-demo/internal/pkg/errcode"
	"testing"
)

// export GOARCH=amd64 && go test -gcflags=all=-l -run TestResp
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

	type TestType struct {
		A int
		B string
	}

	type args struct {
		ctx  *gin.Context
		e    *errcode.Error
		data interface{}
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
		{
			"成功-结构体数据",
			args{
				ctx: &gin.Context{},
				e:   errcode.Success,
				data: TestType{
					A: 100,
					B: "abc",
				},
			},
		},
		{
			"成功-嵌套结构体&tag",
			args{
				ctx: &gin.Context{},
				e:   errcode.Success,
				data: struct {
					TestType
					C int
					D int `json:"d"`
				}{
					TestType: TestType{
						A: 100,
						B: "abc",
					},
					C: 300,
					D: 400,
				},
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
			} else if _, ok := tt.args.data.(gin.H); !ok {
				bytes, _ := json.Marshal(tt.args.data)
				json.Unmarshal(bytes, &inData)
			}
			assert.Equal(t, jsonRes.code, inErr.HTTPStatus())
			assert.Equal(t, inErr.Code(), outData[StatusCode])
			assert.Equal(t, inErr.Msg(), outData[StatusMsg])
			if inErr.Code() == errcode.Success.Code() {
				delete(outData, StatusCode)
				delete(outData, StatusMsg)
				// 在map上失效：https://github.com/stretchr/testify/issues/143
				// assert.Equal(t, inData, outData)
				if fmt.Sprint(inData) != fmt.Sprint(outData) {
					t.Fatalf("got %v want %v\n", outData, inData)
				}
			}
		})
	}
}
