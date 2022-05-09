package app

import (
	"net/http"
	"simple-demo/pkg/errcode"

	"github.com/gin-gonic/gin"
)

//Respond define
type Response struct {
	Ctx        *gin.Context
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}
func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"StatusCode": err.Code(), "StatusMsg": err.Msg()}
	r.Ctx.JSON(err.StatusCode(), response)
}
