package service

import (
	"errors"
	"mime/multipart"
	"os"
	"simple-demo/global"
	"simple-demo/internal/model"
	"simple-demo/pkg/app"
	"simple-demo/pkg/upload"
)

type PublishRequest struct {
	Token      string
	File       multipart.File
	FileHeader *multipart.FileHeader
	FileType   upload.FileType
}

type PublishListRequest struct {
	Token string `form:"token" binding:"required"`
}

type VideoListResponse struct {
	Response
	VideoList []model.Video `json:"video_list"`
}

func (svc *Service) Publish(params *PublishRequest) (*Response, error) {
	claims, err := app.ParseToken(params.Token)
	if err != nil {
		return nil, errors.New("token 不存在")
	}
	fileName := upload.GetFileName(params.FileHeader.Filename, claims.AppKey)
	if !upload.CheckContainExt(fileName) {
		return nil, errors.New("file suffix is not supported.")
	}
	if upload.CheckMaxSize(fileName, params.File) {
		return nil, errors.New("exceeded maximum file limit.")
	}

	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory.")
		}
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions.")
	}

	dst := uploadSavePath + "/" + fileName
	if err := upload.SaveFile(params.FileHeader, dst); err != nil {
		return nil, err
	}
	playUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	err = svc.dao.Publish(claims.AppKey, playUrl, "")
	if err != nil {
		return nil, err
	}
	return nil, nil
}
