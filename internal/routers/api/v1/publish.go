package v1

import (
	"net/http"
	"simple-demo/global"
	"simple-demo/internal/service"
	"simple-demo/pkg/app"
	"simple-demo/pkg/errcode"

	"github.com/gin-gonic/gin"
)

//Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	response := app.NewResponse(c)
	params := service.PublishRequest{}

	params.Token = c.Query("token")
	file, fileHeader, err := c.Request.FormFile("data")
	if err != nil {
		global.Logger.Errorf(c, "app.FormFile err: %v", err)
		response.ToErrorResponse(errcode.PublishError)
		return
	}

	if fileHeader == nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(c.Request.Context())
	params.File, params.FileHeader = file, fileHeader
	respond, err := svc.Publish(&params)
	if err != nil {
		global.Logger.Errorf(c, "svc.Publish err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}
	c.JSON(http.StatusOK, respond)
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	// params := service.PublishListRequest{}
	// response := app.NewResponse(c)
	// valid, errs := app.BindAndValid(c, &params)
	// if !valid {
	// 	global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
	// 	errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
	// 	response.ToErrorResponse(errRsp)
	// 	return
	// }

	// svc := service.New(c.Request.Context())
	// respond, err := svc.PublishList(&params)
	// if err != nil {
	// 	global.Logger.Errorf(c, "svc.PublishList err: %v", err)
	// 	response.ToErrorResponse(errcode.PublishListError)
	// 	return
	// }

	// c.JSON(http.StatusOK, respond)
}
