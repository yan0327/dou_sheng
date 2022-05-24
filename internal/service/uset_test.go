package service

import (
	"github.com/stretchr/testify/assert"
	"os"
	"simple-demo/internal/model"
	"simple-demo/internal/pkg/errcode"
	"simple-demo/internal/pkg/global"
	"simple-demo/pkg/setting"
	"testing"
)

var srv UserSrv

type MockUserDao struct {
	u   map[int64]*model.User
	idx int64
}

func (m MockUserDao) Create(u *model.User) (*model.User, error) {
	if u.Id == 0 {
		u.Id = m.idx
		m.idx++
	}
	m.u[u.Id] = u
	return u, nil
}

func (m MockUserDao) FindById(uid int64) (*model.User, error) {
	return m.u[uid], nil
}

func (m MockUserDao) FindByName(username string) (*model.User, error) {
	for _, v := range m.u {
		if v.Name == username {
			return v, nil
		}
	}
	return nil, nil
}

func (m MockUserDao) FindByIds(uids []int64) ([]*model.User, error) {
	//TODO implement me
	panic("implement me")
}

type MockRelationDao struct {
	u   map[int64]*model.User
	idx int64
}

func (m MockRelationDao) IsFollower(userId int64, toUserId int64) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockRelationDao) Create(userId int64, toUserId int64) error {
	//TODO implement me
	panic("implement me")
}

func (m MockRelationDao) Delete(userId int64, toUserId int64) error {
	//TODO implement me
	panic("implement me")
}

func (m MockRelationDao) FollowList(userId int64) ([]int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockRelationDao) FollowerList(userId int64) ([]int64, error) {
	//TODO implement me
	panic("implement me")
}

func TestMain(m *testing.M) {
	global.JWTSetting = &setting.JWTSettingS{
		Secret: "abc",
		Issuer: "test",
		Expire: 3600,
	}
	os.Exit(m.Run())
}

func setup() {
	srv = MakeUserSrv(MockUserDao{
		u:   map[int64]*model.User{},
		idx: 1,
	}, MockRelationDao{
		u:   map[int64]*model.User{},
		idx: 1,
	})
}

func TestLogin(t *testing.T) {
	setup()
	u, token, e := srv.Login("a", "b")
	assert.Equal(t, e, errcode.ErrorUserNotExistFail)
	assert.Nil(t, u)
	assert.Empty(t, token)
	srv.Register("a", "b")
	u, token, e = srv.Login("a", "b")
	assert.Equal(t, u.Name, "a")
	assert.Equal(t, u.Password, "b")
	assert.NotEmpty(t, token)
	assert.Nil(t, e)
}

func TestRegister(t *testing.T) {
	setup()
	u, token, e := srv.Register("a", "b")
	assert.Equal(t, u.Name, "a")
	assert.Equal(t, u.Password, "b")
	assert.NotEmpty(t, token)
	assert.Nil(t, e)
	u, token, e = srv.Register("a", "b")
	assert.Empty(t, token)
	assert.Equal(t, e, errcode.ErrorUserExistFail)
}

func TestGetById(t *testing.T) {
	setup()
	u, e := srv.GetById(1)
	assert.Nil(t, u, e)
	u, _, _ = srv.Register("a", "b")
	u, e = srv.GetById(u.Id)
	assert.Equal(t, u.Name, "a")
	assert.Equal(t, u.Password, "b")
	assert.Nil(t, e)
}
