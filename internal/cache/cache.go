package cache

import (
	"context"
	"math/big"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	Ctx         context.Context
	Prefix      string
	RedisClient *redis.Client
}

func NewCache(prefix string, redisClient *redis.Client) *Cache {
	return &Cache{Prefix: prefix, RedisClient: redisClient, Ctx: context.Background()}
}
func (r *Cache) Client() *redis.Client {
	return r.RedisClient
}
func (r *Cache) Check() (string, error) {
	return r.RedisClient.Ping(r.Ctx).Result()
}

func (r *Cache) BgSave() (string, error) {
	return r.RedisClient.BgSave(r.Ctx).Result()
}

//获取关注列表
func (r *Cache) GetFollowList(username string) ([]string, error) {
	cmd := r.RedisClient.SMembers(r.Ctx, join(username, r.formatKey("followlist")))
	if cmd.Err() != nil {
		return []string{}, cmd.Err()
	}
	return cmd.Val(), nil
}

//获取粉丝列表
func (r *Cache) GetFollowerList(username string) ([]string, error) {
	cmd := r.RedisClient.SMembers(r.Ctx, join(username, r.formatKey("followerlist")))
	if cmd.Err() != nil {
		return []string{}, cmd.Err()
	}
	return cmd.Val(), nil
}

// func (r *Cache) WriteUserState(username string) error {
// 	tx := r.RedisClient.Multi(r.Ctx)
// 	defer tx.Close()

// 	now := util.MakeTimestamp() / 1000

// 	_, err := tx.Exec(func() error {
// 		tx.HSet(r.formatKey("nodes"), join(id, "name"), id)
// 		tx.HSet(r.formatKey("nodes"), join(id, "height"), strconv.FormatUint(height, 10))
// 		tx.HSet(r.formatKey("nodes"), join(id, "difficulty"), diff.String())
// 		tx.HSet(r.formatKey("nodes"), join(id, "lastBeat"), strconv.FormatInt(now, 10))
// 		return nil
// 	})
// 	return err
// }

func (r *Cache) formatKey(args ...interface{}) string {
	return join(r.Prefix, join(args...))
}

func join(args ...interface{}) string {
	s := make([]string, len(args))
	for i, v := range args {
		switch v.(type) {
		case string:
			s[i] = v.(string)
		case int64:
			s[i] = strconv.FormatInt(v.(int64), 10)
		case uint64:
			s[i] = strconv.FormatUint(v.(uint64), 10)
		case float64:
			s[i] = strconv.FormatFloat(v.(float64), 'f', 0, 64)
		case bool:
			if v.(bool) {
				s[i] = "1"
			} else {
				s[i] = "0"
			}
		case *big.Int:
			n := v.(*big.Int)
			if n != nil {
				s[i] = n.String()
			} else {
				s[i] = "0"
			}
		default:
			panic("Invalid type specified for conversion")
		}
	}
	return strings.Join(s, ":")
}
