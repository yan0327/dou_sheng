package app

import (
	"github.com/gin-gonic/gin"
)

//Respond define
type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}
