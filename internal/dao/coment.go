package dao

import "simple-demo/internal/model"

func (d *Dao) CreateComment(userId uint32, videoId uint32, actionType uint8, commentText string, commentId uint32) error {
	comment := model.Comment{
		UserId:     userId,
		VideoId:    videoId,
		ActionType: actionType,
		Content:    commentText,
		CommentId:  commentId,
	}
	return comment.CreateComment(d.engine)
}

func (d *Dao) GetCommentList(userId, videoId uint32) ([]model.Comment, error) {
	comment := model.Comment{
		VideoId: videoId,
		UserId:  userId,
	}
	return comment.GetCommentList(d.engine)
}
