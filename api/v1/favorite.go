package v1

import (
	"net/http"
	"simple-demo/model"
	"simple-demo/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	VideoID, _ := strconv.Atoi(c.Query("video_id"))
	var islike uint8 = 0
	action_type, _ := strconv.Atoi(c.Query("action_type"))
	if uint8(action_type) == 1 {
		islike = 1
	} else {
		islike = 0
	}
	var user model.User
	var err error
	//user.Username = claims.Username
	cUser, ok := c.Get("token")
	if !ok {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
	user = cUser.(model.User)
	_, err = service.SetOrUpdateFavorite(user.ID, uint(VideoID), islike)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	}

}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	var user model.User
	var err error
	//user.Username = claims.Username
	cUser, ok := c.Get("token")
	if !ok {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
	user = cUser.(model.User)
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
