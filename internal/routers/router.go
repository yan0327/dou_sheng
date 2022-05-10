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

func NewRouter() *gin.Engine {
	r := gin.Default()
	regisMiddleWare(r)

	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	apiRouter := r.Group("/douyin")
	{
		// basic apis
		apiRouter.GET("/feed/", v1.Feed)
		apiRouter.GET("/user/", middleware.JWT(), v1.UserInfo) //token
		apiRouter.POST("/user/register/", v1.Register)
		apiRouter.POST("/user/login/", v1.Login)

		publishRouter := apiRouter.Group("/publish")
		publishRouter.Use(middleware.JWT())
		{
			publishRouter.POST("/action/", v1.Publish)  //token
			publishRouter.GET("/list/", v1.PublishList) //token
		}

		// extra apis - I
		//token
		favoriteRouter := apiRouter.Group("/favorite")
		favoriteRouter.Use(middleware.JWT())
		{
			favoriteRouter.POST("/action/", v1.FavoriteAction)
			favoriteRouter.GET("/list/", v1.FavoriteList)
		}

		commentRouter := apiRouter.Group("/comment")
		commentRouter.Use(middleware.JWT())
		{
			commentRouter.POST("/action/", v1.CommentAction)
			commentRouter.GET("/list/", v1.CommentList)
		}

		// extra apis - II
		//token
		relationRouter := apiRouter.Group("/relation")
		relationRouter.Use(middleware.JWT())
		{
			relationRouter.POST("/action/", v1.RelationAction)
			relationRouter.GET("/follow/list/", v1.FollowList)
			relationRouter.GET("/follower/list/", v1.FollowerList)
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

	r.Use(middleware.Translations())
	r.Use(middleware.Tracing())
}
