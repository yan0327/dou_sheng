package cache

import (
	"simple-demo/internal/model"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

func (r *Cache) WriteVideoCommentCount(videoId int64, count int64) error {
	key := r.formatKey("video", "commentCount", videoId)
	err := r.Client.SetEX(r.Ctx, key, count, time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Cache) DeleteVideoCommentCount(videoId int64) error {
	key := r.formatKey("video", "commentCount", videoId)
	return r.Client.Del(r.Ctx, key).Err()
}

func (r *Cache) WriteVideoComment(id int64, userId int64, videoId int64, commentText string, ts string) error {
	key := r.formatKey("video", "comment", videoId)
	_, err := r.Client.ZAdd(r.Ctx, key, &redis.Z{Score: float64(id), Member: join(id, videoId, userId, commentText, ts)}).Result()
	if err != nil {
		return err
	}
	r.Client.ExpireAt(r.Ctx, key, time.Now().Add(2*time.Hour))
	return nil
}

func (r *Cache) DeleteVideoComment(id int64, videoId int64, comment_id int64) error {
	key := r.formatKey("video", "comment", videoId)
	_, err := r.Client.ZRemRangeByScore(r.Ctx, key, "id", "id").Result()
	return err
}

//id, videoId, userId, commentText, ts
func (r *Cache) GetVideoCommentList(videoId int64) ([]*model.Comment, error) {
	key := r.formatKey("video", "comment", videoId)
	vals, err := r.Client.ZRevRangeByScore(r.Ctx, key, &redis.ZRangeBy{
		Min: "-",
		Max: "+",
	}).Result()
	if err != nil {
		return nil, err
	}
	res := make([]*model.Comment, len(vals))
	for i := 0; i < len(res); i++ {
		strs := strings.Split(vals[i], ":")
		id, _ := strconv.ParseInt(strs[0], 10, 64)
		videoId, _ := strconv.ParseInt(strs[1], 10, 64)
		userId, _ := strconv.ParseInt(strs[2], 10, 64)
		content := strs[3]
		ts := strs[4]
		res[i] = &model.Comment{
			Id:         id,
			VideoId:    videoId,
			UserId:     userId,
			Content:    content,
			CreateDate: ts,
		}
	}
	return res, nil
}
