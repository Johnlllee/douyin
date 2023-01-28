package userSvc

import "sync/atomic"

type UserSvcMgr struct {
	userIdx int64
}

func (usm *UserSvcMgr) RequestUserId() int64 {
	atomic.AddInt64(&usm.userIdx, 1)
	return usm.userIdx
}

func (usm *UserSvcMgr) InitUserSvcMgr(idx int64) {
	usm.userIdx = idx
}

func InsertUserData(username string, password string, userId int64) error {

	return nil
}
