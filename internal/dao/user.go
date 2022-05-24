package dao

import "simple-demo/internal/model"

func (d *Dao) UserRegister(username, password string) (*model.User, error) {
	user := model.User{UserName: username, PassWord: password}
	return user.Register(d.engine)
}

func (d *Dao) UserLogin(username, password string) (*model.User, error) {
	user := model.User{UserName: username, PassWord: password}
	return user.UserLogin(d.engine)
}

func (d *Dao) GetUserInfo(username string) (*model.User, error) {
	user := model.User{UserName: username}
	return user.GetUserInfo(d.engine)
}

// func (d *Dao) UserInfo(id uint32, username, password string)(model.User, error){
// 	user := model.User{UserName: username, PassWord: password}
// }
