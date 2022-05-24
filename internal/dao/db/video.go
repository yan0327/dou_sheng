package db

import (
	"gorm.io/gorm"
	"simple-demo/internal/model"
)

type VideoDao interface {
	Create(video *model.Video) (*model.Video, error)
	FindByTime(time int64) ([]*model.Video, error)
	FindByUser(userId int64) ([]*model.Video, error)
	FindFavoriteByUser(userId int64) ([]*model.Video, error)
}

type videos struct {
	db *gorm.DB
}

func MakeVideos(db *gorm.DB) *videos {
	return &videos{db}
}

func (v *videos) Create(video *model.Video) (*model.Video, error) {
	res := v.db.Create(video)
	return video, res.Error
}

func (v *videos) FindByTime(time int64) ([]*model.Video, error) {
	var res []*model.Video
	err := v.db.Model(&model.Video{}).
		Select(`*, 
			(SELECT COUNT(1) FROM tiktok_video_comment 
				WHERE tiktok_video_comment.video_id = tiktok_video.id) AS comment_count,
			(SELECT COUNT(1) FROM tiktok_video_like 
				WHERE tiktok_video_like.video_id = tiktok_video.id) AS favorite_count
		`).
		Where("unix_timestamp(create_time) < ?", time).
		Order("create_time DESC").
		Limit(30).
		Preload("Author", func(gdb *gorm.DB) *gorm.DB {
			return gdb.Select(`
				*,
				(SELECT COUNT(1) FROM tiktok_relation WHERE tiktok_relation.user_id = tiktok_user.id) AS follower_count,
				(SELECT COUNT(1) FROM tiktok_relation WHERE tiktok_relation.follower_id = tiktok_user.id) AS follow_count
			`)
		}).
		Find(&res).Error
	return res, err
}

func (v *videos) FindByUser(userId int64) ([]*model.Video, error) {
	var res []*model.Video
	err := v.db.Model(&model.Video{}).Where("author_id = ?", userId).
		Select(`*, 
			(SELECT COUNT(1) FROM tiktok_video_comment 
				WHERE tiktok_video_comment.video_id = tiktok_video.id) AS comment_count,
			(SELECT COUNT(1) FROM tiktok_video_like
				WHERE tiktok_video_like.video_id = tiktok_video.id) AS favorite_count
		`).
		Preload("Author", func(gdb *gorm.DB) *gorm.DB {
			return gdb.Select(`
				*,
				(SELECT COUNT(1) FROM tiktok_relation WHERE tiktok_relation.user_id = tiktok_user.id) AS follower_count,
				(SELECT COUNT(1) FROM tiktok_relation WHERE tiktok_relation.follower_id = tiktok_user.id) AS follow_count
			`)
		}).
		Order("create_time DESC").
		Find(&res).Error
	return res, err
}

func (v *videos) FindFavoriteByUser(userId int64) ([]*model.Video, error) {
	var vid []int64
	err := v.db.Table("tiktok_video_like").
		Select("video_id").
		Where("user_id = ?", userId).Find(&vid).Error
	if err != nil {
		return nil, err
	}
	if vid == nil || len(vid) == 0 {
		return []*model.Video{}, nil
	}
	var res []*model.Video
	err = v.db.Model(&model.Video{}).Where("id IN ?", vid).
		Select(`*, 
			(SELECT COUNT(1) FROM tiktok_video_comment 
				WHERE tiktok_video_comment.video_id = tiktok_video.id) AS comment_count,
			(SELECT COUNT(1) FROM tiktok_video_like
				WHERE tiktok_video_like.video_id = tiktok_video.id) AS favorite_count
		`).
		Preload("Author", func(gdb *gorm.DB) *gorm.DB {
			return gdb.Select(`
				*,
				(SELECT COUNT(1) FROM tiktok_relation WHERE tiktok_relation.user_id = tiktok_user.id) AS follower_count,
				(SELECT COUNT(1) FROM tiktok_relation WHERE tiktok_relation.follower_id = tiktok_user.id) AS follow_count
			`)
		}).
		Order("create_time DESC").
		Find(&res).Error
	return res, err
}
