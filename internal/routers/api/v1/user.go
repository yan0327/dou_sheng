package v1

import (
	"net/http"
	"simple-demo/global"
	"simple-demo/internal/service"
	"simple-demo/pkg/app"
	"simple-demo/pkg/errcode"
	"simple-demo/pkg/util"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	response := app.NewResponse(c)
	param := service.UserLoginRequest{}
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	param.PassWord = util.EncodeMD5(param.PassWord)
	svc := service.New(c.Request.Context())
	respond, err := svc.UserLogin(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.Login err: %v", err)
		response.ToErrorResponse(errcode.UserRegisterError)
		return
	}

	token, err := app.GenerateToken(param.UserName, param.PassWord)
	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	respond.Token = token
	c.JSON(http.StatusOK, respond)
}

func UserInfo(c *gin.Context) {
	params := service.UserInfoRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	respond, err := svc.UserInfo(&params)
	if err != nil {
		global.Logger.Errorf(c, "svc.UserInfo err: %v", err)
		response.ToErrorResponse(errcode.UserGetInfoError)
		return
	}

	c.JSON(http.StatusOK, respond)
}

func Register(c *gin.Context) {
	response := app.NewResponse(c)
	param := service.UserRegisterRequest{}
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	param.PassWord = util.EncodeMD5(param.PassWord)
	svc := service.New(c.Request.Context())
	respond, err := svc.UserRegister(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.UserRegisterError)
		return
	}

	token, err := app.GenerateToken(param.UserName, param.PassWord)
	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	respond.Token = token
	c.JSON(http.StatusOK, respond)
}
