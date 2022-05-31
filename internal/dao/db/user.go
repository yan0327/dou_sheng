package db

import (
	"gorm.io/gorm"
	"simple-demo/internal/model"
)

type UserDao interface {
	Create(u *model.User) (*model.User, error)
	FindById(uid int64) (*model.User, error)
	FindByIds(uids []int64) ([]*model.User, error)
	FindByName(username string) (*model.User, error)
}

type users struct {
	db *gorm.DB
}

func MakeUsers(db *gorm.DB) *users {
	return &users{db}
}

func (user *users) Create(u *model.User) (*model.User, error) {
	res := user.db.
		Select("Id", "Name", "Password").
		Create(u)
	return u, res.Error
}

func (user *users) FindById(uid int64) (*model.User, error) {
	var res model.User
	err := user.db.
		Select(`
			*,
			(SELECT COUNT(1) FROM tiktok_relation WHERE tiktok_relation.user_id = tiktok_user.id) AS follower_count,
			(SELECT COUNT(1) FROM tiktok_relation WHERE tiktok_relation.follower_id = tiktok_user.id) AS follow_count,
			(SELECT COUNT(1) FROM tiktok_video_like WHERE tiktok_video_like.user_id = tiktok_user.id) AS favorite_count,
			(SELECT COUNT(1) FROM tiktok_video_like WHERE tiktok_video_like.video_id IN
				(SELECT tiktok_video.id FROM tiktok_video WHERE tiktok_video.author_id = tiktok_user.id)) AS total_favorited
		`).
		First(&res, uid).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &res, err
}

func (user *users) FindByIds(uids []int64) ([]*model.User, error) {
	var res []*model.User
	err := user.db.
		Select(`
			*,
			(SELECT COUNT(1) FROM tiktok_relation WHERE tiktok_relation.user_id = tiktok_user.id) AS follower_count,
			(SELECT COUNT(1) FROM tiktok_relation WHERE tiktok_relation.follower_id = tiktok_user.id) AS follow_count,
			(SELECT COUNT(1) FROM tiktok_video_like WHERE tiktok_video_like.user_id = tiktok_user.id) AS favorite_count,
			(SELECT COUNT(1) FROM tiktok_video_like WHERE tiktok_video_like.video_id IN
				(SELECT tiktok_video.id FROM tiktok_video WHERE tiktok_video.author_id = tiktok_user.id)) AS total_favorited
		`).
		Where("id IN ?", uids).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (user *users) FindByName(username string) (*model.User, error) {
	var res model.User
	err := user.db.
		Select(`
			*,
			(SELECT COUNT(1) FROM tiktok_relation WHERE tiktok_relation.user_id = tiktok_user.id) AS follower_count,
			(SELECT COUNT(1) FROM tiktok_relation WHERE tiktok_relation.follower_id = tiktok_user.id) AS follow_count,
			(SELECT COUNT(1) FROM tiktok_video_like WHERE tiktok_video_like.user_id = tiktok_user.id) AS favorite_count,
			(SELECT COUNT(1) FROM tiktok_video_like WHERE tiktok_video_like.video_id IN
				(SELECT tiktok_video.id FROM tiktok_video WHERE tiktok_video.author_id = tiktok_user.id)) AS total_favorited
		`).
		Where("username = ?", username).First(&res).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &res, err
}
