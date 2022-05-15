package api

import (
	"simple-demo/internal/pkg/api"
	"time"

	"github.com/gin-gonic/gin"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	api.RespWithData(c, gin.H{
		"video_list": DemoVideos,
		"next_time":  time.Now().Unix(),
	})
}
