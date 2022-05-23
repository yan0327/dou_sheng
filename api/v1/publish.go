package v1

import (
	"fmt"
	"net/http"
	"path/filepath"
	"simple-demo/global"
	"simple-demo/middleware"
	"simple-demo/model"
	"simple-demo/model/response"
	"simple-demo/pkg/util"
	"strconv"
	"strings"
	"time"

	"simple-demo/service"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []model.ReplyVideo `json:"video_list"`
}
type LikeListResponse struct {
	Response
	VideoList []model.ReplyVideo `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")
	j := middleware.NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		if err == middleware.TokenExpired {
			response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
			c.Abort()
			return
		}
		response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		c.Abort()
		return
	}
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	var video model.Video
	var user model.User
	filename := filepath.Base(data.Filename)
	user, err = service.FindUser(claims.Username)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	video.Title = title
	finalName := fmt.Sprintf("%d_%s_%s", user.ID, title, filename)
	saveFile := filepath.Join("./public/upload", finalName)
	video.PlayUrl = "http://192.168.30.128:8080/static/" + finalName
	video.AuthorId = user.ID
	video.CreateTime = time.Now()
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//input VideoName output JPGName
	util.Video2JPG(saveFile, strings.Replace(saveFile, ".mp4", "_cover.jpg", 1))
	video.CoverUrl = strings.Replace(video.PlayUrl, ".mp4", "_cover.jpg", 1)
	global.DBEngine.Create(&video)
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	var user model.User
	var err error
	//user.Username = claims.Username
	cUser, ok := c.Get("token")
	if !ok {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
	user = cUser.(model.User)
	userID, _ := strconv.Atoi(c.Query("user_id"))

	var video []model.ReplyVideo

	video, err = service.PublishVideo(user.ID, uint(userID))
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: video,
	})
}
