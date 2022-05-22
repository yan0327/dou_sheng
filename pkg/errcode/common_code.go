package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服务内部错误")
	InvalidParams             = NewError(10000001, "入参错误")
	NotFound                  = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10000003, "鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "鉴权失败，Token错误")
	UnauthorizedTokenTimeout  = NewError(10000005, "鉴权失败，Token超时")
	UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败，Token生成失败")
	TooManyRequests           = NewError(10000007, "请求过多")

	UserRegisterError   = NewError(10000008, "用户注册失败")
	UserGetInfoError    = NewError(10000009, "获取用户消息失败")
	ReverseFeedError    = NewError(20000001, "Feed接口视频流查询失败")
	PublishError        = NewError(20000002, "发布视频错误")
	PublishListError    = NewError(20000003, "获取发布视频错误")
	FavoriteActionError = NewError(20000004, "点赞错误")
	FavoriteListError   = NewError(20000005, "获取点赞视频列表错误")
)