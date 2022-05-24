package db

import "gorm.io/gorm"

type FavoriteDao interface {
	IsFavorite(userId int64, videoId int64) (bool, error)
	Create(userId int64, videoId int64) error
	Delete(userId int64, videoId int64) error
}

type favorites struct {
	db *gorm.DB
}

func MakeFavorites(db *gorm.DB) *favorites {
	return &favorites{db}
}

const favoriteTableName = "tiktok_video_like"

func (f *favorites) IsFavorite(userId int64, videoId int64) (bool, error) {
	var cnt int64
	f.db.Table(favoriteTableName).Where("user_id = ? AND video_id = ?", userId, videoId).Count(&cnt)
	return cnt > 0, nil
}

func (f *favorites) Create(userId int64, videoId int64) error {
	if b, _ := f.IsFavorite(userId, videoId); b {
		return nil
	}
	res := f.db.Table(favoriteTableName).Create(map[string]interface{}{
		"user_id":  userId,
		"video_id": videoId,
	})
	return res.Error
}

func (f *favorites) Delete(userId int64, videoId int64) error {
	res := f.db.Table(favoriteTableName).
		Where("user_id = ? AND video_id = ?", userId, videoId).
		Delete(map[string]interface{}{})
	return res.Error
}
