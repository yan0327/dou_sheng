package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-demo/internal/dao"
	"simple-demo/internal/dao/store"
	"simple-demo/internal/middleware"
	"simple-demo/internal/middleware/auth"
	global2 "simple-demo/internal/pkg/global"
	"simple-demo/internal/routers/controller"
)

func NewRouter() *gin.Engine {
	factory := dao.MakeDaoFactory(
		global2.DBEngine,
		store.MakeFileStore(global2.AppSetting.UploadSavePath),
	)
	user := api.MakeUserController(factory)
	video := api.MakeVideoController(factory)
	comment := api.MakeCommentController(factory)
	favorite := api.MakeFavoriteController(factory)

	r := gin.Default()
	regisMiddleWare(r)

	r.StaticFS("/static", http.Dir(global2.AppSetting.UploadSavePath))

	apiRouter := r.Group("/douyin")
	{
		// basic apis
		apiRouter.GET("/feed/", video.Feed)
		apiRouter.GET("/user/", auth.AuthMiddleware, user.UserInfo)
		apiRouter.POST("/user/register/", user.Register)
		apiRouter.POST("/user/login/", user.Login)
		apiRouter.POST("/publish/action/", auth.AuthMiddleware, video.Publish)
		apiRouter.GET("/publish/list/", auth.AuthMiddleware, video.PublishList)

		// extra apis - I
		apiRouter.POST("/favorite/action/", auth.AuthMiddleware, favorite.FavoriteAction)
		apiRouter.GET("/favorite/list/", auth.AuthMiddleware, favorite.FavoriteList)
		apiRouter.POST("/comment/action/", auth.AuthMiddleware, comment.CommentAction)
		apiRouter.GET("/comment/list/", auth.AuthMiddleware, comment.CommentList)

		// extra apis - II
		apiRouter.POST("/relation/action/", auth.AuthMiddleware, user.RelationAction)
		apiRouter.GET("/relation/follow/list/", auth.AuthMiddleware, user.FollowList)
		apiRouter.GET("/relation/follower/list/", auth.AuthMiddleware, user.FollowerList)
	}

	return r
}

func regisMiddleWare(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//r.Use(middleware.Translations())
	r.Use(middleware.Tracing())
}
