package service

import (
	"mime/multipart"
	"simple-demo/global"
	"simple-demo/model"
)

type PublishRequest struct {
	Token string                `form:"token" binding:"required"`
	Data  *multipart.FileHeader `form:"data"  binding:"required"`
}

type PublishListRequest struct {
	Token string `form:"token" binding:"required"`
}

func PublishVideo(userid, id uint) (replys []model.ReplyVideo, err error) {
	var videos []model.Video
	err = global.DBEngine.Where("author_id = ?", id).Find(&videos).Error
	for i := range videos {

		reply := model.ReplyVideo{
			ID: videos[i].ID,
		}
		var video model.Video
		err = global.DBEngine.Model(&model.Video{}).Where("id = ?", reply.ID).First(&video).Error
		reply.PlayUrl = video.PlayUrl
		reply.CoverUrl = video.CoverUrl
		reply.Author, err = FindReplyUser(id, videos[i].AuthorId)
		if err != nil {
			return
		}
		reply.FavoriteCount, err = GetFavoriteNum(reply.ID)
		if err != nil {
			return
		}
		reply.CommentCount, err = GetCommentsNum(reply.ID)
		if err != nil {
			return
		}
		reply.IsFavorite = IsLike(userid, video.ID)
		replys = append(replys, reply)
	}
	return
}
