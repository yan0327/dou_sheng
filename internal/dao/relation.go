package dao

import "simple-demo/internal/model"

func (d *Dao) RelationAction(username string, toUserId int64, acitonType uint8) error {
	relation := model.Relation{
		UserId:     toUserId,
		UserName:   username,
		ActionType: acitonType,
	}
	return relation.RelationAction(d.engine)
}

func (d *Dao) FollowList(userId int64) ([]*model.User, error) {
	relation := model.Relation{
		FollowerId: userId,
	}
	return relation.FollowList(d.engine)
}

func (d *Dao) FollowerList(userId int64) ([]*model.User, error) {
	relation := model.Relation{
		Id: userId,
	}
	return relation.FollowerList(d.engine)
}
