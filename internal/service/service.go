package service

import (
	"simple-demo/global"
	"simple-demo/internal/cache"
	"simple-demo/internal/dao"

	"github.com/gin-gonic/gin"
)

type Service struct {
	ctx   *gin.Context
	dao   *dao.Dao
	cache *cache.Cache
}

func New(ctx *gin.Context) Service {
	svc := Service{ctx: ctx}
	svc.cache = cache.NewCache(global.RedisSetting.Prefix, global.RedisEngine)
	svc.dao = dao.NewDao(global.DBEngine, svc.cache)
	return svc
}
