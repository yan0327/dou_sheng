package service

type CommentRequest struct {
	UserId      uint32 `form:"user_id" binding:"required"`
	Token       string `form:"token" binding:"required"`
	VideoId     uint32 `form:"video_id" binding:"required"`
	ActionType  uint8  `form:"action_type" binding:"required"`
	CommentText string `form:"comment_text"`
	CommentId   uint32 `form:"comment_id"`
}

type CommentListRequest struct {
	UserId  uint32 `form:"user_id" binding:"required"`
	Token   string `form:"token" binding:"required"`
	VideoId uint32 `form:"video_id" binding:"required"`
}
