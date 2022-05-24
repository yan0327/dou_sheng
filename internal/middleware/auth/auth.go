package auth

import (
	"github.com/gin-gonic/gin"
	"simple-demo/internal/pkg/api"
	"simple-demo/internal/pkg/errcode"
	"simple-demo/internal/pkg/global"
	"simple-demo/internal/pkg/jwt"
	"simple-demo/internal/service"
	"strconv"
	"time"
)

func JWtAuth(token string) (int64, *errcode.Error) {
	claims, err := jwt.ParseJWT(token)
	if err != nil {
		return 0, errcode.UnauthorizedTokenError
	}

	illegal := false
	if v, has := claims["issuer"]; !has {
		illegal = true
	} else if issuer, ok := v.(string); !ok || issuer != global.JWTSetting.Issuer {
		illegal = true
	}
	if illegal {
		return 0, errcode.UnauthorizedTokenError
	}

	v, has := claims["expire"]
	if !has {
		return 0, errcode.UnauthorizedTokenError
	}
	expire, ok := v.(string)
	if !ok {
		return 0, errcode.UnauthorizedTokenError
	}
	if expireInt, err := strconv.Atoi(expire); err != nil {
		return 0, errcode.UnauthorizedTokenError
	} else if int64(expireInt) < time.Now().Unix() {
		return 0, errcode.UnauthorizedTokenTimeout
	}

	return int64(claims[service.UserId].(float64)), nil
}

func IsLogin(c *gin.Context) (int64, bool) {
	token := c.Query("token")
	if token == "" {
		token = c.PostForm("token")
	}
	userId, err := JWtAuth(token)
	if err != nil {
		return 0, false
	}
	return userId, true
}

func AuthMiddleware(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		token = c.PostForm("token")
		if token == "" {
			api.RespWithErr(c, errcode.UnauthorizedTokenError)
			c.Abort()
			return
		}
	}

	userId, err := JWtAuth(token)
	if err != nil {
		api.RespWithErr(c, err)
		c.Abort()
		return
	}

	c.Set(service.UserId, userId)
}
