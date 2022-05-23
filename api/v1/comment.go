package v1

import (
	"net/http"
	"simple-demo/middleware"
	"simple-demo/model"
	"simple-demo/model/response"
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
	token := c.Query("token")
	VideoID, _ := strconv.Atoi(c.Query("video_id"))
	action_type, _ := strconv.Atoi(c.Query("action_type"))
	content_text := c.Query("comment_text")
	comment_id, _ := strconv.Atoi(c.Query("comment_id"))
	var err error
	var user model.User
	j := middleware.NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		if err == middleware.TokenExpired {
			response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
			c.Abort()
			return
		}
		response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		c.Abort()
		return
	}
	user, err = service.FindUser(claims.Username)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
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
	token := c.Query("token")
	VideoID, _ := strconv.Atoi(c.Query("video_id"))
	var err error
	var user model.User
	var comments []model.Comment
	ReplyComment := []model.ReplyComment{}
	j := middleware.NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		if err == middleware.TokenExpired {
			response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
			c.Abort()
			return
		}
		response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		c.Abort()
		return
	}
	user, err = service.FindUser(claims.Username)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	comments, err = service.GetComments(uint(VideoID))
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
