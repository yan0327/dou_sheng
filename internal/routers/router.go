package routers

import (
	"net/http"
	"simple-demo/global"
	"simple-demo/internal/middleware"
	v1 "simple-demo/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	regisMiddleWare(r)

	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	apiRouter := r.Group("/douyin")
	{
		// basic apis
		apiRouter.GET("/feed/", v1.Feed)
		apiRouter.GET("/user/", v1.UserInfo)
		apiRouter.POST("/user/register/", v1.Register)
		apiRouter.POST("/user/login/", v1.Login)
		apiRouter.POST("/publish/action/", v1.Publish)
		apiRouter.GET("/publish/list/", v1.PublishList)

		// extra apis - I
		apiRouter.POST("/favorite/action/", v1.FavoriteAction)
		apiRouter.GET("/favorite/list/", v1.FavoriteList)
		apiRouter.POST("/comment/action/", v1.CommentAction)
		apiRouter.GET("/comment/list/", v1.CommentList)

		// extra apis - II
		apiRouter.POST("/relation/action/", v1.RelationAction)
		apiRouter.GET("/relation/follow/list/", v1.FollowList)
		apiRouter.GET("/relation/follower/list/", v1.FollowerList)
	}

	return r
}

func regisMiddleWare(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	r.Use(middleware.Tracing())
}
