package service

import (
	"context"
	"simple-demo/global"
	"simple-demo/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.NewDao(global.DBEngine)
	return svc
}
