package service

import (
	"errors"
	"simple-demo/global"
	"simple-demo/model"

	"simple-demo/pkg/util"
)

func Register(u model.User) (err error, userInter model.User) {
	var user model.User
	total := global.DBEngine.Where("username = ?", u.Username).First(&user).RowsAffected
	if total > 0 {
		return errors.New("用户名已注册"), u
	}

	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = util.EncodeMD5(u.Password)
	err = global.DBEngine.Create(&u).Error
	return err, u
}

func Login(u *model.User) (err error, userInter *model.User) {
	var user model.User
	u.Password = util.EncodeMD5(u.Password)
	err = global.DBEngine.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return err, &user
}
