package routers

import (
	"net/http"
	v1 "simple-demo/api/v1"
	"simple-demo/global"
	"simple-demo/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	regisMiddleWare(r)

	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	privateGroup := r.Group("/douyin")
	privateGroup.Use(middleware.JWTAuth())
	{
		privateGroup.POST("/relation/action/", v1.RelationAction)
		privateGroup.GET("/relation/follow/list/", v1.FollowList)
		privateGroup.GET("/relation/follower/list/", v1.FollowerList)
		privateGroup.GET("/publish/list/", v1.PublishList)
		// extra apis - I
		privateGroup.POST("/favorite/action/", v1.FavoriteAction)
		privateGroup.GET("/favorite/list/", v1.FavoriteList)
		// extra apis - II
		privateGroup.POST("/comment/action/", v1.CommentAction)
		privateGroup.GET("/comment/list/", v1.CommentList)
	}
	publicGroup := r.Group("/douyin")
	{
		publicGroup.GET("/feed/", v1.Feed)
		publicGroup.POST("/user/register/", v1.Register)
		publicGroup.POST("/user/login/", v1.Login)
		publicGroup.POST("/publish/action/", v1.Publish)

		// basic apis

		publicGroup.GET("/user/", v1.UserInfo)

	}

	return r
}

func regisMiddleWare(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Tracing())
}
