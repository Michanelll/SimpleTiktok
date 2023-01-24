package repo

import "errors"

type Login struct {
	Id         int64 `gorm:"primary_key"`
	UserInfoId int64
	Username   string `gorm:"primary_key"`
	Password   string `gorm:"size:100;notnull"`
}

//账号密码查询
func QueryUserLogin(username, password string, login *Login) error {
	if login == nil {
		return errors.New("Err:结构体空指针")
	}
	Db.Where("username=? and password=?", username, password).First(login)
	if login.Id == 0 {
		return errors.New("Err:账号或密码出错")
	}
	return nil
}

//账号查询
func IsUserExistByUsername(username string) bool {
	var userLogin Login
	Db.Where("username=?", username).First(&userLogin)
	if userLogin.Id == 0 {
		return false
	}
	return true
}
