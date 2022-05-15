package api

import (
	"fmt"
	"path/filepath"
	"simple-demo/internal/pkg/api"
	"simple-demo/internal/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; !exist {
		api.RespWithErr(c, errcode.UnauthorizedTokenError)
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		api.RespOK(c)
		return
	}

	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		fmt.Printf("%w", err)
		api.RespWithErr(c, errcode.ServerError)
		return
	}

	api.RespOK(c)
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	api.RespWithData(c, gin.H{
		"video_list": DemoVideos,
	})
}
