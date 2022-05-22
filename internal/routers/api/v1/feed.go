package v1

import (
	"net/http"
	"simple-demo/global"
	"simple-demo/internal/service"
	"simple-demo/pkg/app"
	"simple-demo/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Feed struct{}

func NewFeed() *Feed {
	return &Feed{}
}

// Feed same demo video list for every request
func (f *Feed) GetFeed(c *gin.Context) {
	response := app.NewResponse(c)
	param := service.FeedRequest{}
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	respond, err := svc.ReverseFeed(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.Login err: %v", err)
		response.ToErrorResponse(errcode.ReverseFeedError)
		return
	}
	c.JSON(http.StatusOK, respond)
}
