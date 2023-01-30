package service

import (
	"douyin/model"
	"fmt"
	"sync"
)

func InitAllServiceMgr() {

	//	UserServiceManager.InitUserSvcMgr()

}

type UserSvcMgr struct {
	userIdx int64

	mutex sync.RWMutex
}

var UserServiceManager UserSvcMgr

func (usm *UserSvcMgr) InitUserSvcMgr() {
	var lastestId *int64

	err := model.QueryLastUserId(lastestId)

	if err != nil {
		fmt.Println(err.Error())
		//panic("[UserSvc] UserServiceManager Init Userid Fail. Please Check Database.")
	}
	usm.userIdx = *lastestId

}
