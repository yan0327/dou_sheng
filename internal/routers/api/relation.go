package api

import (
	"simple-demo/internal/model"
	"simple-demo/internal/pkg/api"
	"simple-demo/internal/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		api.RespOK(c)
	} else {
		api.RespWithErr(c, errcode.UnauthorizedTokenError)
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	api.RespWithData(c, gin.H{
		"user_list": []model.User{DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	api.RespWithData(c, gin.H{
		"user_list": []model.User{DemoUser},
	})
}
