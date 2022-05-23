package v1

import (
	"net/http"
	"simple-demo/middleware"
	"simple-demo/model"
	"simple-demo/model/response"
	"simple-demo/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	VideoID, _ := strconv.Atoi(c.Query("video_id"))
	var islike uint8 = 0
	action_type, _ := strconv.Atoi(c.Query("action_type"))

	var err error
	var user model.User
	if uint8(action_type) == 1 {
		islike = 1
	} else {
		islike = 0
	}
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
	user, err = service.FindUser(claims.Username)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
	_, err = service.SetOrUpdateFavorite(user.ID, uint(VideoID), islike)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	}

}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	userID, _ := strconv.Atoi(c.Query("user_id"))
	var err error
	var user model.User
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
	user, err = service.FindUser(claims.Username)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
	ReplyVideos := []model.ReplyVideo{}
	//ReplyVideos, err = service.FavoriteList(user_id)
	ReplyVideos, err = service.FavoriteList(user.ID, uint(userID))
	if err != nil {
		c.JSON(http.StatusOK, LikeListResponse{
			Response: Response{
				StatusCode: 1,
			},
			VideoList: []model.ReplyVideo{},
		})
		return
	}
	c.JSON(http.StatusOK, LikeListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: ReplyVideos,
	})
}
