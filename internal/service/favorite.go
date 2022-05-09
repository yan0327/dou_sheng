package service

type FavoriteRequest struct {
	UserId     uint32  `form:"user_id" binding:"required"`
	Token      string `form:"token" binding:"required"`
	VideoId    uint32  `form:"video_id" binding:"required"`
	ActionType int    `form:"action_type" binding:"required, oneof= 1 2"`
}

type FavoriteListRequest struct {
	UserId uint32  `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}
