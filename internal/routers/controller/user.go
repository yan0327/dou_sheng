package api

import (
	"github.com/gin-gonic/gin"
	"simple-demo/internal/dao"
	"simple-demo/internal/middleware/auth"
	"simple-demo/internal/model"
	"simple-demo/internal/pkg/api"
	"simple-demo/internal/pkg/errcode"
	"simple-demo/internal/pkg/global"
	"simple-demo/internal/service"
	"simple-demo/pkg/app"
)

const (
	ActionFollow = iota + 1
	ActionCancelFollow
)

type UserLoginResponse struct {
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserInfoRequest struct {
	UserId int64 `form:"user_id" binding:"required"`
}

type RelationActionRequest struct {
	ToUserId   int64 `form:"to_user_id" binding:"required"`
	ActionType int   `form:"action_type" binding:"required"`
}

type UserResponse struct {
	User *model.User `json:"user"`
}

type UserListResponse struct {
	UserList []*model.User `json:"user_list"`
}

type UserController struct {
	usrv service.UserSrv
}

func MakeUserController(factory *dao.DaoFactory) *UserController {
	return &UserController{
		usrv: service.MakeUserSrv(factory.User(), factory.Relation()),
	}
}

func (u *UserController) processIsFollow(c *gin.Context, list []*model.User) {
	userId, isLogin := auth.IsLogin(c)
	if !isLogin {
		return
	}
	// 已经登录
	if followList, err := u.usrv.FollowList(userId); err == nil {
		n := len(followList)
		m := len(list)
		i := 0
		j := 0
		for i < n && j < m {
			if followList[i].Id == list[j].Id {
				list[j].IsFollow = true
				i++
			} else if followList[i].Id > list[j].Id {
				i++
			} else {
				j++
			}
		}
	}
}

func (u *UserController) Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	user, token, err := u.usrv.Login(username, password)
	if err != nil {
		api.RespWithErr(c, err)
		return
	}
	api.RespWithData(c, UserLoginResponse{
		UserId: user.Id,
		Token:  token,
	})
}

func (u *UserController) Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	user, token, err := u.usrv.Register(username, password)
	if err != nil {
		api.RespWithErr(c, err)
		return
	}
	api.RespWithData(c, UserLoginResponse{
		UserId: user.Id,
		Token:  token,
	})
}

func (u *UserController) UserInfo(c *gin.Context) {
	req := UserInfoRequest{}
	if valid, errs := app.BindAndValid(c, &req); !valid {
		api.RespWithErr(c, errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user, err := u.usrv.GetById(req.UserId)
	if err != nil {
		api.RespWithErr(c, err)
		return
	}
	api.RespWithData(c, UserResponse{
		User: user,
	})
}

func (u *UserController) RelationAction(c *gin.Context) {
	uid, exist := c.Get(service.UserId)
	if !exist {
		api.RespWithErr(c, errcode.ServerError.WithDetails("获取当前用户信息失败"))
		global.Logger.Error(c, "获取当前用户信息失败")
		return
	}

	req := RelationActionRequest{}
	if valid, errs := app.BindAndValid(c, &req); !valid {
		api.RespWithErr(c, errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	if req.ActionType == ActionFollow || req.ActionType == ActionCancelFollow {
		err := u.usrv.Follow(uid.(int64), req.ToUserId)
		if err != nil {
			api.RespWithErr(c, err)
			return
		}
		api.RespOK(c)
	} else {
		api.RespWithErr(c, errcode.InvalidParams)
	}
}

func (u *UserController) FollowList(c *gin.Context) {
	req := UserInfoRequest{}
	if valid, errs := app.BindAndValid(c, &req); !valid {
		api.RespWithErr(c, errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	ret, err := u.usrv.FollowList(req.UserId)
	if err != nil {
		api.RespWithErr(c, err)
		global.Logger.Error(c, err.Details())
		return
	}
	for i := range ret {
		ret[i].IsFollow = true
	}

	api.RespWithData(c, UserListResponse{
		UserList: ret,
	})
}

func (u *UserController) FollowerList(c *gin.Context) {
	req := UserInfoRequest{}
	if valid, errs := app.BindAndValid(c, &req); !valid {
		api.RespWithErr(c, errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	ret, err := u.usrv.FollowerList(req.UserId)
	if err != nil {
		api.RespWithErr(c, err)
		global.Logger.Error(c, err.Details())
		return
	}
	u.processIsFollow(c, ret)

	api.RespWithData(c, UserListResponse{
		UserList: ret,
	})
}
