package routers

import (
	"simple-demo/internal/middleware"
	"simple-demo/internal/routers/api"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	regisMiddleWare(r)

	r.Static("/static", "./public/static")

	apiRouter := r.Group("/douyin")
	{
		// basic apis
		apiRouter.GET("/feed/", api.Feed)
		apiRouter.GET("/user/", api.UserInfo)
		apiRouter.POST("/user/register/", api.Register)
		apiRouter.POST("/user/login/", api.Login)
		apiRouter.POST("/publish/action/", api.Publish)
		apiRouter.GET("/publish/list/", api.PublishList)

		// extra apis - I
		apiRouter.POST("/favorite/action/", api.FavoriteAction)
		apiRouter.GET("/favorite/list/", api.FavoriteList)
		apiRouter.POST("/comment/action/", api.CommentAction)
		apiRouter.GET("/comment/list/", api.CommentList)

		// extra apis - II
		apiRouter.POST("/relation/action/", api.RelationAction)
		apiRouter.GET("/relation/follow/list/", api.FollowList)
		apiRouter.GET("/relation/follower/list/", api.FollowerList)
	}

	return r
}

func regisMiddleWare(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	r.Use(middleware.Tracing())
}
