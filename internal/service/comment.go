package service

type CommentRequest struct {
	UserId      int64  `json:"user_id, omitempty"`
	Token       string `json:"token, omitempty"`
	VideoId     int64  `json:"video_id, omitempty"`
	ActionType  int    `json:"action_type, omitempty"`
	CommentText string `json:"comment_text"`
	CommentId   int64  `json:"comment_id"`
}

type CommentListRequest struct {
	UserId  int64  `json:"user_id, omitempty"`
	Token   string `json:"token, omitempty"`
	VideoId int64  `json:"video_id, omitempty"`
}
