package api

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"simple-demo/internal/dao"
	"simple-demo/internal/middleware/auth"
	"simple-demo/internal/model"
	"simple-demo/internal/pkg/api"
	"simple-demo/internal/pkg/errcode"
	"simple-demo/internal/pkg/global"
	"simple-demo/internal/service"
	"simple-demo/pkg/app"
	"sort"
	"time"
)

type FeedRequest struct {
	LatestTime int64 `form:"latest_time" json:",string"`
}

type PublishListRequest struct {
	UserId int64 `form:"user_id" json:"user_id,string"`
}

type FeedResponse struct {
	VideoList []*model.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty,string"`
}

type VideoListResponse struct {
	VideoList []*model.Video `json:"video_list"`
}

type VideoController struct {
	video    service.VideoSrv
	user     service.UserSrv
	favorite service.FavoriteSrv
}

func MakeVideoController(f *dao.DaoFactory) *VideoController {
	return &VideoController{
		service.MakeVideoSrv(f.Store(), f.Video()),
		service.MakeUserSrv(f.User(), f.Relation()),
		service.MakeFavoriteSrv(f.Favorite(), f.Video())}
}

func (v *VideoController) processIsFavorite(c *gin.Context, list []*model.Video) {
	userId, isLogin := auth.IsLogin(c)
	if !isLogin {
		return
	}
	// 已经登录
	if likeList, err := v.favorite.ListByUser(userId); err == nil {
		n := len(likeList)
		m := len(list)
		i := 0
		j := 0
		for i < n && j < m {
			if likeList[i].Id == list[j].Id {
				list[j].IsFavorite = true
				i++
			} else if likeList[i].Id > list[j].Id {
				i++
			} else {
				j++
			}
		}
	}
}

func (v *VideoController) processIsFollow(c *gin.Context, list []*model.Video) {
	userId, isLogin := auth.IsLogin(c)
	if !isLogin {
		return
	}
	// 已经登录
	if followList, err := v.user.FollowList(userId); err == nil {
		for i := range list {
			n := len(followList)
			if sort.Search(n, func(j int) bool {
				return followList[j].Id == list[i].Author.Id
			}) < n {
				list[i].Author.IsFollow = true
			}
		}
	}
}

func (v *VideoController) VideoData(c *gin.Context) {
	// TODO 视频格式问题
	r, e := v.video.DataStream(c.Param("id"))
	if e != nil {
		api.RespWithErr(c, e)
		return
	}
	c.Status(http.StatusOK)
	//io.Copy(c.Writer, r)
	c.Stream(func(w io.Writer) bool {
		buf := make([]byte, 1024*100)
		if n, _ := r.Read(buf); n == 0 {
			return false
		}
		w.Write(buf)
		return true
	})
}

func (v *VideoController) Feed(c *gin.Context) {
	arg := FeedRequest{}

	if valid, err := app.BindAndValid(c, &arg); !valid || err != nil {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", err)
		errRsp := errcode.InvalidParams.WithDetails(err.Errors()...)
		api.RespWithErr(c, errRsp)
		return
	}
	if arg.LatestTime == 0 {
		arg.LatestTime = time.Now().Unix()
	} else {
		// 客户端传的这玩意格式都不对 不是以秒为单位 而是以毫秒为单位的
		arg.LatestTime /= 1000
	}

	ret, err := v.video.Feed(arg.LatestTime)
	if err != nil {
		api.RespWithErr(c, errcode.ServerError)
		return
	}
	nextTime := int64(0)
	if len(ret) > 0 {
		nextTime = int64(ret[len(ret)-1].CreatedAt.Unix()) - 1
	}

	v.processIsFavorite(c, ret)
	v.processIsFollow(c, ret)
	api.RespWithData(c, FeedResponse{
		VideoList: ret,
		NextTime:  nextTime,
	})
}

func (v *VideoController) Publish(c *gin.Context) {
	uid, exist := c.Get(service.UserId)
	if !exist {
		api.RespWithErr(c, errcode.ServerError.WithDetails("获取当前用户信息失败"))
		global.Logger.Error(c, "获取当前用户信息失败")
		return
	}
	title, exist := c.GetPostForm("title")
	if !exist {
		api.RespWithErr(c, errcode.InvalidParams)
		return
	}

	file, _ := c.FormFile("data")
	if file == nil {
		api.RespWithErr(c, errcode.InvalidParams.WithDetails("文件为空"))
		global.Logger.Debug(c, "文件为空")
		return
	}

	f, err := file.Open()
	defer f.Close()
	if err != nil {
		api.RespWithErr(c, errcode.ServerError.WithDetails(err.Error()))
		global.Logger.Error(c, err.Error())
		return
	}

	if err := v.video.Publish(uid.(int64), title, bufio.NewReader(f)); err != nil {
		api.RespWithErr(c, err)
		global.Logger.Error(c, err.Details())
		return
	}
	api.RespOK(c)

}

func (v *VideoController) PublishList(c *gin.Context) {
	arg := PublishListRequest{}
	if valid, err := app.BindAndValid(c, &arg); !valid || err != nil {
		api.RespWithErr(c, errcode.InvalidParams)
		return
	}
	ret, err := v.video.FindByUser(arg.UserId)
	if err != nil {
		api.RespWithErr(c, err)
		return
	}

	v.processIsFavorite(c, ret)
	v.processIsFollow(c, ret)
	api.RespWithData(c, VideoListResponse{
		VideoList: ret,
	})
}
