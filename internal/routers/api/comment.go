package api

import (
	"simple-demo/internal/pkg/api"
	"simple-demo/internal/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		api.RespOK(c)
	} else {
		api.RespWithErr(c, errcode.UnauthorizedTokenError)
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	api.RespWithData(c, gin.H{
		"comment_list": DemoComments,
	})
}
