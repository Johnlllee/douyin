package model

import "errors"

// jojoleee 用户登录表
type UserLogin struct {
	Id         int64 `gorm:"primary_key"`
	UserInfoId int64
	Username   string `gorm:"primary_key"`
	Password   string `gorm:"size:200;notnull"`
}

func QueryUserLogin(username, password string, login *UserLogin) error {
	/**
	 * @Author jojoleee
	 * @Description //TODO
	 * @Param [username, password, login]
	 * @return
	 **/
	if login == nil {
		return ErrIvdPtr
	}
	DB.Where("username=? and password=?", username, password).First(login)
	if login.Id == 0 {
		return errors.New("用户不存在，账号或密码出错")
	}
	return nil
}

// jojoleee 为什么不需要这个函数往数据表中插入用户已经登陆的这样的信息呢？
//func SetUserLogin(login *UserLogin) error {
//	if login == nil {
//		return ErrIvdPtr
//	}
//
//}

func HasUserLoginByName(username string) bool {
	var user UserLogin
	DB.Where("username=?", username).First(&user)
	if user.Id == 0 {
		return false
	}
	return true
}
