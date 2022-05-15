package api

import (
	"simple-demo/global"
	"simple-demo/internal/model"
	"simple-demo/internal/pkg/api"
	"simple-demo/internal/pkg/errcode"
	"simple-demo/internal/service"
	"simple-demo/pkg/app"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]model.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if _, exist := usersLoginInfo[token]; exist {
		api.RespWithErr(c, errcode.ErrorUserExistFail)
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		newUser := model.User{
			Id:   userIdSequence,
			Name: username,
		}
		usersLoginInfo[token] = newUser
		api.RespWithData(c, gin.H{
			"user_id": userIdSequence,
			"token":   username + password,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, exist := usersLoginInfo[token]; exist {
		api.RespWithData(c, gin.H{
			"user_id": user.Id,
			"token":   token,
		})
	} else {
		api.RespWithErr(c, errcode.NotFound)
	}
}

func UserInfo(c *gin.Context) {
	req := service.UserInfoRequest{}
	valid, errs := app.BindAndValid(c, &req)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		api.RespWithErr(c, errRsp)
		return
	}

	if user, exist := usersLoginInfo[req.Token]; exist {
		api.RespWithData(c, gin.H{
			"user": user,
		})
	} else {
		api.RespWithErr(c, errcode.NotFound)
	}
}
