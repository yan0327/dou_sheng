package service

type RelationRequest struct {
	UserId     int64  `json:"user_id,omitempty"`
	Token      string `json:"token",omitempty"`
	ToUserId   int64  `json:"to_user_id",omitempty"`
	ActionType int    `json:"action_type",omitempty"`
}

//关注列表
type FollowListRequest struct {
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token",omitempty"`
}

//粉丝列表
type FollowerListRequest struct {
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token",omitempty"`
}
