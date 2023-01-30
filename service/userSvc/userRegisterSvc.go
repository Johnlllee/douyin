package userSvc

import (
	"douyin/model"
	"douyin/service"
	"errors"
	"fmt"
	"sync"
)

type UserSvcMgr struct {
	userIdx int64

	mutex sync.RWMutex
}

func (usm *UserSvcMgr) RequestUserId() int64 {
	usm.mutex.Lock()
	defer usm.mutex.Unlock()

	usm.userIdx += 1
	return usm.userIdx
}

func (usm *UserSvcMgr) InitUserSvcMgr() {
	var lastestId *int64

	err := model.QueryLastUserId(lastestId)

	if err != nil {
		fmt.Println(err.Error())
		panic("[UserSvc] UserServiceManager Init Userid Fail. Please Check Database.")
	}
	usm.userIdx = *lastestId

}

func CheckUserName(username string) error {
	userinfo := &model.UserInfo{}
	err := model.QueryUserInfoByName(username, userinfo)
	if userinfo != nil {
		return errors.New("[UserSvc]: Username Already Exist.")
	}
	return err
}

func InsertUserData(username string, password string, userId int64) error {

	service.UserServiceManager.mutex.Lock()
	defer service.UserServiceManager.mutex.Unlock()

	err := CheckUserName(username)
	if err != nil {
		return err
	}

	userinfo := &model.UserInfo{
		Id:            userId,
		Name:          username,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
		User: &model.UserLogin{
			Id:         userId,
			UserInfoId: userId,
			Password:   password,
			Username:   username,
		},
	}

	err = nil
	err = model.AddUserInfo(userinfo)

	if err != nil {
		return err
	}
	return nil
}
