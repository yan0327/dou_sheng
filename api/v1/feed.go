package v1

import (
	"net/http"
	"simple-demo/model"
	"simple-demo/service"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []model.ReplyVideo `json:"video_list,omitempty"`
	NextTime  int64              `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	//LastTime := c.Query("latest_time")
	var videos []model.Video
	var author model.User
	var err error
	ReplyVideo := []model.ReplyVideo{}
	videos, err = service.FeedVideos()
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{StatusCode: 2},
		})
	}
	for i := range videos {
		replyvideo := model.ReplyVideo{
			ID:       videos[i].ID,
			PlayUrl:  videos[i].PlayUrl,
			CoverUrl: videos[i].CoverUrl,
		}
		author, err = service.FindAuthor(videos[i].AuthorId)
		if err != nil {
			c.JSON(http.StatusOK, FeedResponse{
				Response: Response{StatusCode: 2},
			})
		}
		replyvideo.Author = model.ReplyUser{
			ID:       author.ID,
			Username: author.Username,
		}
		replyvideo.Author.FollowCount, err = service.GetFollowNum(author.ID)
		if err != nil {
			c.JSON(http.StatusOK, FeedResponse{
				Response: Response{StatusCode: 2},
			})
		}
		replyvideo.Author.FollowerCount, err = service.GetFollowerNum(author.ID)
		if err != nil {
			c.JSON(http.StatusOK, FeedResponse{
				Response: Response{StatusCode: 2},
			})
		}
		//result = global.DBEngine.Model(&model.Realtion{}).Where("user_id = ? AND follower_id = ?", author.ID)
		replyvideo.FavoriteCount, err = service.GetFavoriteNum(videos[i].ID)
		if err != nil {
			c.JSON(http.StatusOK, FeedResponse{
				Response: Response{StatusCode: 2},
			})
		}

		replyvideo.CommentCount, err = service.GetCommentsNum(videos[i].ID)
		if err != nil {
			c.JSON(http.StatusOK, FeedResponse{
				Response: Response{StatusCode: 2},
			})
		}
		replyvideo.Title = videos[i].Title
		ReplyVideo = append(ReplyVideo, replyvideo)
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: ReplyVideo,
		NextTime:  time.Now().Unix(),
	})
}
