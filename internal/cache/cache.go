package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	Ctx    context.Context
	Prefix string
	Client *redis.Client
}

func NewCache(prefix string, redisClient *redis.Client) *Cache {
	return &Cache{Prefix: prefix, Client: redisClient, Ctx: context.Background()}
}
func (r *Cache) CacheClient() *redis.Client {
	return r.Client
}
func (r *Cache) Check() (string, error) {
	return r.Client.Ping(r.Ctx).Result()
}

func (r *Cache) BgSave() (string, error) {
	return r.Client.BgSave(r.Ctx).Result()
}

//获取关注列表
func (r *Cache) GetFollowList(username string) ([]string, error) {
	cmd := r.Client.SMembers(r.Ctx, join(username, r.formatKey("followlist")))
	if cmd.Err() != nil {
		return []string{}, cmd.Err()
	}
	return cmd.Val(), nil
}

//获取粉丝列表
func (r *Cache) GetFollowerList(username string) ([]string, error) {
	cmd := r.Client.SMembers(r.Ctx, join(username, r.formatKey("followerlist")))
	if cmd.Err() != nil {
		return []string{}, cmd.Err()
	}
	return cmd.Val(), nil
}
