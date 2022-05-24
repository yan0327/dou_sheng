package cache

import (
	"simple-demo/internal/model"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

//视频信息不需要更新
func (r *Cache) WriteVideoInfo(videoId, author_id, playUrl, imagePlayUrl, title string, ts int64) error {
	key := r.formatKey("video", "info", videoId)
	_, err := r.Client.ZAdd(r.Ctx, key, &redis.Z{Score: float64(ts), Member: join(videoId, author_id, playUrl, imagePlayUrl, ts, title)}).Result()
	if err != nil {
		return err
	}
	r.Client.ExpireAt(r.Ctx, key, time.Now().Add(2*time.Hour))
	return nil
}

func (r *Cache) GetVideosInfo(userName string, ts int64) ([]*model.Video, error) {
	vals, err := r.Client.ZRevRangeByScore(r.Ctx, r.formatKey("video", "info", userName), &redis.ZRangeBy{
		Min: "-",
		Max: "+",
	}).Result()
	if err != nil {
		return nil, err
	}
	//id, author_id, playUrl, imagePlayUrl, ts, title)
	res := make([]*model.Video, len(vals))
	for i := 0; i < len(res); i++ {
		strs := strings.Split(vals[i], ":")
		id, _ := strconv.ParseInt(strs[0], 10, 64)
		author_id, _ := strconv.ParseInt(strs[1], 10, 64)
		playUrl := strs[2]
		coverUrl := strs[3]
		ts := strs[4]
		title := strs[5]
		res[i] = &model.Video{
			Id:         id,
			AuthorId:   author_id,
			PlayUrl:    playUrl,
			CoverUrl:   coverUrl,
			CreateTime: ts,
			Title:      title,
		}
	}
	return res, nil
}

func (r *Cache) WriteVideoFavoriteCount(videoId int64, count int64) error {
	key := r.formatKey("video", "favoriteCount", videoId)
	err := r.Client.SetEX(r.Ctx, key, count, time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Cache) GetVideoFavoriteCount(videoId int64) (int64, error) {
	val, err := r.Client.Get(r.Ctx, r.formatKey("video", "favoriteCount", videoId)).Result()
	if err != nil {
		return -1, err
	}
	res, _ := strconv.ParseInt(val, 10, 64)
	return res, nil
}

func (r *Cache) GetVideoCommentCount(videoId int64) (int64, error) {
	val, err := r.Client.Get(r.Ctx, r.formatKey("video", "commentCount", videoId)).Result()
	if err != nil {
		return -1, err
	}
	res, _ := strconv.ParseInt(val, 10, 64)
	return res, nil
}

func (r *Cache) DeleteVideoFavoriteCount(videoId int64) error {
	key := r.formatKey("video", "favoriteCount", videoId)
	return r.Client.Del(r.Ctx, key).Err()
}

func (r *Cache) WriteVideoFavoriteListInfo(userName string, videos []*model.Video) error {
	key := r.formatKey("video", "favoriteList", userName)
	for i := 0; i < len(videos); i++ {
		videoId := videos[i].Id
		author_id := videos[i].AuthorId
		playUrl := videos[i].PlayUrl
		imagePlayUrl := videos[i].CoverUrl
		title := videos[i].Title
		ts, _ := strconv.ParseInt(videos[i].CreateTime, 10, 64)
		r.Client.ZAdd(r.Ctx, key, &redis.Z{Score: float64(ts), Member: join(videoId, author_id, playUrl, imagePlayUrl, ts, title)}).Result()
	}
	r.Client.ExpireAt(r.Ctx, key, time.Now().Add(2*time.Hour))
	return nil
}

func (r *Cache) GetVideoFavoriteListInfo(userName string) ([]*model.Video, error) {
	vals, err := r.Client.ZRevRangeByScore(r.Ctx, r.formatKey("video", "favoriteList", userName), &redis.ZRangeBy{
		Min: "-",
		Max: "+",
	}).Result()

	if err != nil {
		return nil, err
	}
	//id, author_id, playUrl, imagePlayUrl, ts, title)
	res := make([]*model.Video, len(vals))
	for i := 0; i < len(res); i++ {
		strs := strings.Split(vals[i], ":")
		id, _ := strconv.ParseInt(strs[0], 10, 64)
		author_id, _ := strconv.ParseInt(strs[1], 10, 64)
		playUrl := strs[2]
		coverUrl := strs[3]
		ts := strs[4]
		title := strs[5]
		res[i] = &model.Video{
			Id:         id,
			AuthorId:   author_id,
			PlayUrl:    playUrl,
			CoverUrl:   coverUrl,
			CreateTime: ts,
			Title:      title,
		}
	}
	return res, nil
}

func (r *Cache) DeleteVideoFavoriteList(userName string) error {
	key := r.formatKey("video", "favoriteList", userName)
	return r.Client.Del(r.Ctx, key).Err()
}
