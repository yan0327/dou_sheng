package routers

import (
	"simple-demo/internal/routers/api"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Static("/static", "./public")

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
