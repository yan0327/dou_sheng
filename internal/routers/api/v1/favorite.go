package v1

import (
	"net/http"
	"simple-demo/global"
	"simple-demo/internal/service"
	"simple-demo/pkg/app"
	"simple-demo/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Favorite struct{}

func NewFavorite() *Favorite {
	return &Favorite{}
}

// FavoriteAction no practical effect, just check if token is valid
func (f *Favorite) FavoriteAction(c *gin.Context) {
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

// FavoriteList all users have same favorite video list
func (f *Favorite) FavoriteList(c *gin.Context) {
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

	svc := service.New(c)
	respond, err := svc.FavoriteList(&params)
	if err != nil {
		global.Logger.Errorf(c, "svc.FavoriteList err: %v", err)
		response.ToErrorResponse(errcode.FavoriteActionError)
		return
	}
	c.JSON(http.StatusOK, respond)
}
