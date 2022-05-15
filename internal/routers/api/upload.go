package api

import (
	"simple-demo/global"
	"simple-demo/internal/pkg/api"
	errcode2 "simple-demo/internal/pkg/errcode"
	"simple-demo/internal/service"
	"simple-demo/pkg/convert"
	"simple-demo/pkg/upload"

	"github.com/gin-gonic/gin"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		api.RespWithErr(c, errcode2.InvalidParams.WithDetails(err.Error()))
		return
	}

	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		api.RespWithErr(c, errcode2.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf(c, "svc.UploadFile err: %v", err)
		api.RespWithErr(c, errcode2.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	api.RespWithData(c, gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
