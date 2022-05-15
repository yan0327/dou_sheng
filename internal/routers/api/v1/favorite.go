package v1

import (
	"fmt"
	"net/http"
	"simple-demo/global"
	"simple-demo/internal/service"
	"simple-demo/pkg/app"
	"simple-demo/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	params := service.FavoriteRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	respond, err := svc.FavoriteAction(&params)
	if err != nil {
		global.Logger.Errorf(c, "svc.FavoriteAction err: %v", err)
		response.ToErrorResponse(errcode.FavoriteActionError)
		return
	}

	c.JSON(http.StatusOK, respond)
}

func HandleGetAllData(c *gin.Context) {
	//log.Print("handle log")
	// body, _ := ioutil.ReadAll(c.Request.Body)
	// fmt.Println("---body/--- \r\n " + string(body))

	fmt.Println("---header/--- \r\n")
	for k, v := range c.Request.Header {
		fmt.Println(k, v)
	}
	//fmt.Println("header \r\n",c.Request.Header)
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	// HandleGetAllData(c)
	params := service.FavoriteListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &params)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	respond, err := svc.FavoriteList(&params)
	if err != nil {
		global.Logger.Errorf(c, "svc.FavoriteList err: %v", err)
		response.ToErrorResponse(errcode.FavoriteActionError)
		return
	}
	c.JSON(http.StatusOK, respond)
}
