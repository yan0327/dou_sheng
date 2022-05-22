package routers

import (
	"net/http"
	"simple-demo/global"
	"simple-demo/internal/middleware"
	v1 "simple-demo/internal/routers/api/v1"
	"simple-demo/pkg/limiter"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	CommentController  *v1.Comment
	FavoriteController *v1.Favorite
	FeedController     *v1.Feed
	PublishController  *v1.Publish
	RelationController *v1.Relation
	UserController     *v1.User
)

func init() {
	CommentController = v1.NewComment()
	FavoriteController = v1.NewFavorite()
	FeedController = v1.NewFeed()
	PublishController = v1.NewPublish()
	RelationController = v1.NewRelation()
	UserController = v1.NewUser()
}

func NewRouter() *gin.Engine {
	r := gin.Default()
	regisMiddleWare(r)

	r.StaticFS("/static/video", http.Dir(global.AppSetting.UploadVideoSavePath))
	r.StaticFS("/static/image", http.Dir(global.AppSetting.UploadImageSavePath))
	apiRouter := r.Group("/douyin")
	{
		// basic apis
		apiRouter.GET("/feed/", FeedController.GetFeed)
		apiRouter.GET("/user/", middleware.JWT(), UserController.UserInfo) //token
		apiRouter.POST("/user/register/", UserController.Register)
		apiRouter.POST("/user/login/", UserController.Login)

		publishRouter := apiRouter.Group("/publish")
		// publishRouter.Use(middleware.JWT())
		{
			publishRouter.POST("/action/", PublishController.PublishVedio) //token
			publishRouter.GET("/list/", PublishController.PublishList)     //token
		}

		// extra apis - I
		//token
		favoriteRouter := apiRouter.Group("/favorite")
		favoriteRouter.Use(middleware.JWT())
		{
			favoriteRouter.POST("/action/", FavoriteController.FavoriteAction)
			favoriteRouter.GET("/list/", FavoriteController.FavoriteList)
		}

		commentRouter := apiRouter.Group("/comment")
		commentRouter.Use(middleware.JWT())
		{
			commentRouter.POST("/action/", CommentController.CommentAction)
			commentRouter.GET("/list/", CommentController.CommentList)
		}

		// extra apis - II
		//token
		relationRouter := apiRouter.Group("/relation")
		relationRouter.Use(middleware.JWT())
		{
			relationRouter.POST("/action/", RelationController.RelationAction)
			relationRouter.GET("/follow/list/", RelationController.FollowList)
			relationRouter.GET("/follower/list/", RelationController.FollowerList)
		}

	}

	return r
}

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/tiktok_auth", //限制请求的key
		FillInterval: 120 * time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func regisMiddleWare(r *gin.Engine) {
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	// r.Use(middleware.Translations())
	r.Use(middleware.Tracing())
}
