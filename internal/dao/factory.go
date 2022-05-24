package dao

import (
	"gorm.io/gorm"
	"simple-demo/internal/dao/db"
	"simple-demo/internal/dao/store"
)

type DaoFactory struct {
	user     db.UserDao
	video    db.VideoDao
	comment  db.CommentDao
	favorite db.FavoriteDao
	relation db.RelationDao
	store    store.Store
	kv       store.Store
}

func MakeDaoFactory(gdb *gorm.DB, storage store.Store /*, kv store.Store*/) *DaoFactory {
	return &DaoFactory{
		user:     db.MakeUsers(gdb),
		video:    db.MakeVideos(gdb),
		comment:  db.MakeComments(gdb),
		favorite: db.MakeFavorites(gdb),
		relation: db.MakeRelations(gdb),
		store:    storage,
		//kv:    storage,
	}
}

func (f *DaoFactory) User() db.UserDao {
	return f.user
}

func (f *DaoFactory) Video() db.VideoDao {
	return f.video
}

func (f *DaoFactory) Comment() db.CommentDao {
	return f.comment
}

func (f *DaoFactory) Favorite() db.FavoriteDao {
	return f.favorite
}

func (f *DaoFactory) Relation() db.RelationDao {
	return f.relation
}

func (f *DaoFactory) Store() store.Store {
	return f.store
}
