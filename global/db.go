package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type RedisClient struct {
	prefix string
}

var (
	DBEngine    *gorm.DB
	RedisEngine *redis.Client
)
