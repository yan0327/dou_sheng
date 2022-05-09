package service

type RelationRequest struct {
	UserId     uint32  `form:"user_id" binding:"required"`
	Token      string `form:"token" binding:"required"`
	ToUserId   uint32  `form:"to_user_id" binding:"required"`
	ActionType uint8    `form:"action_type" binding:"required, oneof= 1 2"`
}

//关注列表
type FollowListRequest struct {
	UserId uint32  `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

//粉丝列表
type FollowerListRequest struct {
	UserId uint32  `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}
