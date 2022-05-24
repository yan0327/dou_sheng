package service

import (
	"errors"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"simple-demo/internal/model"
	"simple-demo/pkg/app"
	"simple-demo/pkg/upload"
	"strings"
)

type PublishRequest struct {
	Token      string `form:"token"`
	File       multipart.File
	FileHeader *multipart.FileHeader
	Title      string `form:"title"`
}

type PublishListRequest struct {
	Token  string `form:"token"`
	UserId int64  `form:"user_id"`
}

type VideoListResponse struct {
	*Response
	VideoList []model.Video `json:"video_list"`
}

func (svc *Service) PublishList(params *PublishListRequest) (*VideoListResponse, error) {
	vedios, err := svc.dao.PublishList(params.UserId)
	if err != nil {
		return nil, err
	}
	respond := &VideoListResponse{
		Response: &Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: vedios,
	}
	return respond, nil
}

func (svc *Service) Publish(params *PublishRequest) (*Response, error) {
	claims, err := app.ParseToken(params.Token)
	if err != nil {
		return nil, errors.New("ParseToken is err")
	}
	username := claims.AppKey

	fileName := upload.GetFileName(params.FileHeader.Filename, username)
	if !upload.CheckContainExt(fileName) {
		return nil, errors.New("file suffix is not supported.")
	}
	if upload.CheckMaxSize(fileName, params.File) {
		return nil, errors.New("exceeded maximum file limit.")
	}

	videoSavePath, videoplayUrl := upload.GetSavaAndPlayPath(params.FileHeader.Filename)

	if upload.CheckSavePath(videoSavePath) {
		if err := upload.CreateSavePath(videoSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory.")
		}
	}
	if upload.CheckPermission(videoSavePath) {
		return nil, errors.New("insufficient file permissions.")
	}

	playUrl := videoplayUrl + "/" + fileName
	videodst := videoSavePath + "/" + fileName
	if err := upload.SaveFile(params.FileHeader, videodst); err != nil {
		return nil, err
	}
	imagedst := strings.Replace(videodst, ".mp4", ".jpg", 1)
	imagedst, _ = filepath.Abs(strings.Replace(imagedst, "uploadsVideo", "uploadsImage", 1))
	cmdArguments := []string{"-i", videodst, "-y", "-f", "image2", "-t", "1", "-s", "1364x900", imagedst}
	cmd := exec.Command("ffmpeg", cmdArguments...)
	_ = cmd.Run()
	imagePlayUrl := strings.Replace(playUrl, ".mp4", ".jpg", 1)
	imagePlayUrl = strings.Replace(imagePlayUrl, "video", "image", 1)

	err = svc.dao.Publish(username, playUrl, imagePlayUrl, params.Title)
	if err != nil {
		return nil, err
	}
	return &Response{StatusCode: 0, StatusMsg: "success"}, nil
}

// func (svc *Service) PublishList(params *PublishListRequest) (*VideoListResponse, error) {
// 	return nil, nil
// }
