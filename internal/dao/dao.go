package dao

import (
	"simple-demo/internal/cache"

	"github.com/jinzhu/gorm"
)

type Dao struct {
	engine *gorm.DB
	cache  *cache.Cache
}

func NewDao(engine *gorm.DB, cache *cache.Cache) *Dao {
	return &Dao{
		engine: engine,
		cache:  cache,
	}
}
