package service

import "mime/multipart"

type PublishRequest struct {
	Token string                `json:"token",omitempty"`
	Data  *multipart.FileHeader `json:"data",omitempty"`
}

type PublishListRequest struct {
	Token string `json:"token",omitempty"`
}
