package v1

import (
	"simple-demo/model"
	"time"
)

var DemoVideos = []model.Video{
	{
		ID:       1,
		PlayUrl:  "https://www.w3schools.com/html/movie.mp4",
		CoverUrl: "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
	},
}

var DemoComments = []model.Comment{
	{
		ID:         1,
		UserId:     1,
		Content:    "Test Comment",
		CreateTime: time.Now(),
	},
}

var DemoUser = model.User{
	ID:       1,
	Username: "abc",
	Password: "123",
}
