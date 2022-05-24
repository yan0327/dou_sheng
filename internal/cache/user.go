package cache

import (
	"errors"
	"simple-demo/internal/model"
	"strconv"
	"time"
)

func (r *Cache) WriteUserState(id int64, userName string, followCount int64, followerCount int64) error {
	key := r.formatKey("users", userName)
	err := r.Client.HMSet(r.Ctx, key, map[string]interface{}{
		"id":            id,
		"followCount":   followCount,
		"followerCount": followerCount,
	}).Err()
	if err != nil {
		return err
	}
	r.Client.ExpireAt(r.Ctx, key, time.Now().Add(60*time.Hour))
	return nil
}

// func (r *Cache) GetAllUserStates() ([]map[string]interface{}, error) {
// 	users := r.Client.HGetAll(r.Ctx, r.formatKey("users"))
// 	if users.Err() != nil {
// 		return nil, users.Err()
// 	}
// 	m := make(map[string]map[string]interface{})
// 	for key, value := range users.Val() {
// 		parts := strings.Split(key, ":")
// 		if val, ok := m[parts[0]]; ok {
// 			val[parts[1]] = value
// 		} else {
// 			node := make(map[string]interface{})
// 			node[parts[1]] = value
// 			m[parts[0]] = node
// 		}
// 	}
// 	v := make([]map[string]interface{}, len(m), len(m))
// 	for i, value := range m {
// 		id, _ := strconv.Atoi(i)
// 		v[id] = value
// 	}
// 	return v, nil
// }

func (r *Cache) GetUserStates(username string) (*model.User, error) {
	followCount, err := r.GetUserFollowCount(username)
	if err != nil {
		return nil, err
	}
	followerCount, err := r.GetUserFollowerCount(username)
	if err != nil {
		return nil, err
	}
	id, err := r.GetUserId(username)
	if err != nil {
		return nil, err
	}
	user := model.User{
		ID:            id,
		UserName:      username,
		FollowCount:   followCount,
		FollowerCount: followerCount,
	}
	return &user, nil
}

func (r *Cache) GetUserId(userName string) (int64, error) {
	cntStr, err := r.Client.HGet(r.Ctx, r.formatKey("users", userName), "id").Result()
	if err != nil {
		return 0, errors.New("key not find")
	}
	cnt, _ := strconv.ParseInt(cntStr, 10, 64)
	return cnt, nil
}

func (r *Cache) GetUserFollowCount(username string) (int64, error) {
	cntStr, err := r.Client.HGet(r.Ctx, r.formatKey("users", username), "followCount").Result()
	if err != nil {
		return 0, errors.New("key not find")
	}
	cnt, _ := strconv.Atoi(cntStr)
	return int64(cnt), nil
}

func (r *Cache) GetUserFollowerCount(username string) (int64, error) {
	cntStr, err := r.Client.HGet(r.Ctx, r.formatKey("users", username), "followerCount").Result()
	if err != nil {
		return 0, errors.New("key not find")
	}
	cnt, _ := strconv.Atoi(cntStr)
	return int64(cnt), nil
}
