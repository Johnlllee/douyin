package userSvc

import (
	"douyin/model"
	"errors"
	"fmt"
	"sync"
)

type UserSvcMgr struct {
	userIdx int64

	mutex sync.RWMutex
}

//无需分配
//func (usm *UserSvcMgr) RequestUserId() int64 {
//
//	usm.userIdx += 1
//	return usm.userIdx
//}

func (usm *UserSvcMgr) InitUserSvcMgr() {
	var lastestId *int64

	err := model.QueryLastUserId(lastestId)

	if err != nil {
		fmt.Println(err.Error())
		//panic("[UserSvc] UserServiceManager Init Userid Fail. Please Check Database.")
	}
	usm.userIdx = *lastestId

}

func CheckUserName(username string) error {
	userinfo := &model.UserInfo{}
	err := model.QueryUserInfoByName(username, userinfo)
	if err != model.ErrUserNotExist {
		return errors.New("[UserSvc]: Username Already Exist.")
	} else {
		return nil
	}
}

func InsertUserData(username string, password string) (int64, error) {

	err := CheckUserName(username)
	if err != nil {
		return -1, errors.New("User Already Exist.")
	}

	userinfo := &model.UserInfo{
		Name:          username,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
		User: &model.UserLogin{
			Password: password,
			Username: username,
		},
	}

	err = nil
	err = model.AddUserInfo(userinfo)

	if err != nil {
		return -1, err
	}
	return userinfo.Id, nil
}
