package v1

type UserRegisterRespond struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}
