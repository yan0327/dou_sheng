package service

import (
	"fmt"
	"simple-demo/internal/model"
	"time"
)

type FeedRequest struct {
	LatestTime int64 `form:"latest_time"`
}
type FeedResponse struct {
	Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

//返回倒数的视频流
func (svc *Service) ReverseFeed(params *FeedRequest) (*FeedResponse, error) {
	t := params.LatestTime
	if t == 0 {
		t = time.Now().Unix()
	}
	vedios, err := svc.dao.ReverseFeed(t)
	if err != nil {
		return nil, err
	}
	fmt.Println("FeedResponse", vedios)
	respond := &FeedResponse{
		VideoList: vedios,
		NextTime:  time.Now().Unix(),
	}
	return respond, nil
}
