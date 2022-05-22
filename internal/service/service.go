package service

import (
	"context"
	"simple-demo/global"
	"simple-demo/internal/cache"
	"simple-demo/internal/dao"
)

type Service struct {
	ctx   context.Context
	dao   *dao.Dao
	cache *cache.Cache
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.NewDao(global.DBEngine)
	svc.cache = cache.NewCache(global.RedisSetting.Prefix, global.RedisEngine)
	return svc
}
