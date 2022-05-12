package dao

import "simple-demo/internal/model"

func (d *Dao) UserRegister(username, password string) (uint32, error) {
	user := model.User{UserName: username, PassWord: password}
	return user.Register(d.engine)
}

func (d *Dao) UserLogin(username, password string) (uint32, error) {
	user := model.User{UserName: username, PassWord: password}
	return user.UserLogin(d.engine)
}

func (d *Dao) GetUserInfo(id uint32) (model.User, error) {
	user := model.User{ID: id}
	return user.GetUserInfo(d.engine)
}

// func (d *Dao) UserInfo(id uint32, username, password string)(model.User, error){
// 	user := model.User{UserName: username, PassWord: password}
// }
