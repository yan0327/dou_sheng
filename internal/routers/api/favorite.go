package api

import (
	"github.com/gin-gonic/gin"
	"simple-demo/internal/pkg/api"
	"simple-demo/internal/pkg/errcode"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		api.RespOK(c)
	} else {
		api.RespWithErr(c, errcode.NotFound)
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	api.RespWithData(c, gin.H{
		"video_list": DemoVideos,
	})
}
