package dao

import "simple-demo/internal/model"

func (d *Dao) RelationAction(userId uint32, toUserId uint32, acitonType uint8) error {
	relation := model.Relation{
		UserId:     toUserId,
		FollowerId: userId,
		ActionType: acitonType,
	}
	return relation.RelationAction(d.engine)
}

func (d *Dao) FollowList(userId uint32) ([]model.User, error) {
	relation := model.Relation{
		Id: userId,
	}
	return relation.FollowList(d.engine)
}

func (d *Dao) FollowerList(userId uint32) ([]model.User, error) {
	relation := model.Relation{
		Id: userId,
	}
	return relation.FollowerList(d.engine)
}
