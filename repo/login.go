package repo

import (
	"errors"
	"sync"
)

type Login struct {
	Id         int64 `gorm:"primary_key"`
	UserInfoId int64
	Username   string `gorm:"primary_key"`
	Password   string `gorm:"size:100;notnull"`
}

type LoginDAO struct {
}

var (
	loginDao  *LoginDAO
	loginOnce sync.Once
)

func NewUserLoginDao() *LoginDAO {
	loginOnce.Do(func() {
		loginDao = new(LoginDAO)
	})
	return loginDao
}

//账号密码查询
func (u *LoginDAO) QueryUserLogin(username, password string, login *Login) error {
	if login == nil {
		return errors.New("空指针错误")
	}
	Db.Where("username=? and password=?", username, password).First(login)
	if login.Id == 0 {
		return errors.New("用户不存在，账号或密码出错")
	}
	return nil
}

//账号查询
func (u *LoginDAO) IsUserExistByUsername(username string) bool {
	var login Login
	Db.Where("username=?", username).First(&login)
	if login.Id == 0 {
		return false
	}
	return true
}
