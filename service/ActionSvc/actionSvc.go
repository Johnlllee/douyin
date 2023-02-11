package ActionSvc

import (
	"douyin/cache"
	"douyin/model"
	"errors"
)

const FOLLOW int = 1
const UN_FOLLOW int = 2

func FollowOrUnfollowUser(userid, toUserid int64, action int) error {

	exist, err := model.IsUserExistByUserId(toUserid)
	if exist == false {
		return errors.New("[ActionSvc] ToUserid Not Exist.")
	}

	if err != nil {
		return err
	}

	if userid == toUserid {
		return errors.New("[ActionSvc] Userid = ToUserid.")
	}

	switch action {

	case FOLLOW:

		errFromDb := model.AddUserFollow(userid, toUserid)
		errFromRedis := cache.NewRedisProxy().SetFollow(userid, toUserid, true)

		if errFromDb != nil {
			return errors.New("[DB]: " + errFromDb.Error())
		}

		if errFromRedis != nil {
			return errors.New("[Redis]: " + errFromRedis.Error())
		}

		//if errFromDb != nil {
		//	return errors.New("[DB]: " + errFromDb.Error())
		//}
		return nil
	case UN_FOLLOW:
		errFromDb := model.DeleteUserFollow(userid, toUserid)
		errFromRedis := cache.NewRedisProxy().SetFollow(userid, toUserid, false)

		if errFromDb != nil {
			return errors.New("[DB]: " + errFromDb.Error())
		}

		if errFromRedis != nil {
			return errors.New("[Redis]: " + errFromRedis.Error())
		}

		//if errFromDb != nil {
		//	return errors.New("[DB]: " + errFromDb.Error())
		//}
		return nil
	}

	return nil
}
