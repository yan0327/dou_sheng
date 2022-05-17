package v1

import (
	"net/http"
	"simple-demo/global"
	"simple-demo/model"
	"simple-demo/service"
	"time"

	"simple-demo/middleware"
	"simple-demo/model/request"
	"simple-demo/model/response"

	"simple-demo/pkg/app"
	"simple-demo/pkg/errcode"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]model.User{
	"zhangleidouyin": {
		ID:       1,
		Username: "qqq",
		Password: "123",
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
	User model.ReplyUser `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	u := model.User{Username: username, Password: password}
	err, user := service.Register(u)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	}
	TokenNext(c, &user)

}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	u := model.User{Username: username, Password: password}
	err, user := service.Login(&u)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
	TokenNext(c, user)

}

func UserInfo(c *gin.Context) {
	req := service.UserInfoRequest{}
	res := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		res.ToErrorResponse(errRsp)
		return
	}
	token := req.Token
	if token == "" {
		res.ToResponse(UserInfoResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
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
	var user model.User
	var reply model.ReplyUser
	user, err = service.FindUser(claims.Username)
	if err != nil {
		res.ToResponse(UserInfoResponse{
			Response: Response{StatusCode: 1},
			User:     reply,
		})
	}

	//user.Username = claims.Username
	reply, err = service.FindReplyUser(user.ID, user.ID)
	if err != nil {
		res.ToResponse(UserInfoResponse{
			Response: Response{StatusCode: 1},
			User:     reply,
		})
	}
	res.ToResponse(UserInfoResponse{
		Response: Response{StatusCode: 0},
		User:     reply,
	})

}

//it will offer JWT token to client (When user try to register and login)
func TokenNext(c *gin.Context, user *model.User) {
	j := &middleware.JWT{SigningKey: []byte(global.JWTSetting.Secret)} // 唯一签名
	claims := request.CustomClaims{
		ID:       user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                     // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.JWTSetting.Expire, // 过期时间 7天  配置文件
			Issuer:    "ym",                                         // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.Logger.Errorf(c, "token create fail errs: %v", err)
		response.FailWithMessage("获取token失败", c)
		return
	}

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   int64(user.ID),
		Token:    token,
	})

}
