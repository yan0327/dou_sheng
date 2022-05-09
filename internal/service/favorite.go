package service

type FavoriteRequest struct {
	UserId     int64  `json:"user_id, omitempty"`
	Token      string `json:"token, omitempty"`
	VideoId    int64  `json:"video_id,omitempty"`
	ActionType int    `json:"action_type, omitempty"`
}

type FavoriteListRequest struct {
	UserId int64  `json:"user_id, omitempty"`
	Token  string `json:"token, omitempty"`
}

