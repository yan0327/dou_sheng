package v1

import (
	"net/http"
	"simple-demo/model"
	"simple-demo/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	Response
	CommentList []model.ReplyComment `json:"comment_list,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	VideoID, _ := strconv.Atoi(c.Query("video_id"))
	action_type, _ := strconv.Atoi(c.Query("action_type"))
	content_text := c.Query("comment_text")
	comment_id, _ := strconv.Atoi(c.Query("comment_id"))
	var user model.User
	var err error
	//user.Username = claims.Username
	cUser, ok := c.Get("token")
	if !ok {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
	user = cUser.(model.User)

	if uint8(action_type) == 1 {
		err = service.CreateComment(user.ID, uint(VideoID), content_text)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		}
	} else {
		err = service.DeleteComment(uint(comment_id))
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		}
	}
	c.JSON(http.StatusOK, Response{StatusCode: 0})

}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	VideoID, _ := strconv.Atoi(c.Query("video_id"))
	var user model.User
	var err error
	//user.Username = claims.Username
	cUser, ok := c.Get("token")
	if !ok {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
	user = cUser.(model.User)
	var comments []model.Comment
	ReplyComment := []model.ReplyComment{}

	comments, err = service.GetComments(uint(VideoID))
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
	for i := 0; i < len(comments); i++ {
		reply := model.ReplyComment{
			ID:         comments[i].ID,
			Content:    comments[i].Content,
			CreateTime: comments[i].CreateTime.Format("2006-01-02 15:04:05"),
		}
		replyuser, err := service.FindReplyUser(user.ID, comments[i].UserId)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		}
		reply.User = replyuser
		ReplyComment = append(ReplyComment, reply)
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: ReplyComment,
	})
}
