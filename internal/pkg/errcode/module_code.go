package errcode

var (
	ErrorGetCommentListFail = NewError(20010001, "获取评论列表失败")
	ErrorCreateCommentFail  = NewError(20010002, "创建评论失败")
	ErrorDeleteCommentFail  = NewError(20010003, "删除评论失败")
	ErrorCountCommentFail   = NewError(20010004, "统计评论失败")

	ErrorUserExistFail    = NewError(20020001, "用户名已存在")
	ErrorUserNotExistFail = NewError(20020002, "用户名不存在")

	ErrorUploadFileFail = NewError(20030001, "上传文件失败")
)
