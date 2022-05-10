package v1

import (
	"net/http"
	"simple-demo/global"
	"simple-demo/internal/model"
	"simple-demo/internal/service"
	"simple-demo/pkg/app"
	"simple-demo/pkg/errcode"
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

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	Response
	User model.User `json:"user"`
}

func Register(c *gin.Context) {
	response := app.NewResponse(c)

	username := c.Query("username")
	password := c.Query("password")

	/*
		1.写入数据库
		2. 签发token
	*/
	token, err := app.GenerateToken(username, password)
	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	atomic.AddInt64(&userIdSequence, 1)
	newUser := model.User{
		Id:   userIdSequence,
		Name: username,
	}
	usersLoginInfo[token] = newUser
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   userIdSequence,
		Token:    username + password,
	})

}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	/*
		1. 数据库校验密码是否正确
		2. 如果正确则签发token
		3. 根据username and password 签发token时需要对密码进行加密处理,JWT可以被反向破解
	*/
	// svc := service.New(c.Request.Context())
	// err := svc.CheckAuth(&param)
	// if err != nil {
	// 	global.Logger.Errorf(c, "svc.CheckAuth err: %v", err)
	// 	response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
	// 	return
	// }

	// token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	// if err != nil {
	// 	global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
	// 	response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
	// 	return
	// }
	token := username + password

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	req := service.UserInfoRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	if user, exist := usersLoginInfo[req.Token]; exist {
		response.ToResponse(UserInfoResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		response.ToResponse(UserInfoResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
