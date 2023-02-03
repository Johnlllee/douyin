package model

import (
	"douyin/config"
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	ErrIvdPtr        = errors.New("[DB]空指针错误")
	ErrEmptyUserList = errors.New("[DB]用户列表为空")
	ErrUserNotExist  = errors.New("[DB]用户不存在")
)

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(config.DBConnectString()), &gorm.Config{
		PrepareStmt:            true, //缓存预编译命令
		SkipDefaultTransaction: true, //禁用默认事务操作
		//Logger:                 logger.Default.LogMode(logger.Info), //打印sql语句
	})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&UserInfo{}, UserLogin{}, &Video{})
	if err != nil {
		panic(err)
	}
}
