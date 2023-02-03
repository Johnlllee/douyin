package service

import "douyin/service/userSvc"

var UserServiceManager userSvc.UserSvcMgr

func InitAllServiceMgr() {

	//todo 应该从数据库里面去拿到上一次的idx
	//UserServiceManager.InitUserSvcMgr(1)

}
