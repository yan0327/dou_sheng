package dao

import "simple-demo/internal/model"

func (d *Dao) CreateComment(username string, videoId uint32, actionType uint8, commentText string, commentId uint32) error {
	comment := model.Comment{
		// UserId:     userId,
		VideoId:    videoId,
		ActionType: actionType,
		Content:    commentText,
		CommentId:  commentId,
		User: &model.User{
			UserName: username,
		},
	}
	return comment.CreateComment(d.engine)
}

func (d *Dao) GetCommentList(username string, videoId uint32) ([]model.Comment, error) {
	comment := model.Comment{
		VideoId: videoId,
		User: &model.User{
			UserName: username,
		},
	}
	return comment.GetCommentList(d.engine)
}
