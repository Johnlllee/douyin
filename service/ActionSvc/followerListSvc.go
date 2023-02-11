package ActionSvc

import (
	"douyin/cache"
	"douyin/model"
	"errors"
)

func GetFollowerList(useridToQuery int64) (*UserInfoList, error) {

	exist, err := model.IsUserExistByUserId(useridToQuery)
	if exist == false {
		return nil, errors.New("[ActionSvc]: GetFollowerList Userid To Query Not Exist In Db.")
	}

	if err != nil {
		return nil, errors.New("[ActionSvc]: GetFollowerList " + err.Error())
	}

	var tempList []*model.UserInfo

	err = model.GetFollowerListByUserId(useridToQuery, &tempList)

	if err != nil {
		return nil, errors.New("[ActionSvc]: GetFollowerList " + err.Error())
	}

	for _, v := range tempList {
		flag, err := cache.NewRedisProxy().GetFollowStatus(useridToQuery, v.Id)
		if err != nil {
			return nil, err
		}
		v.IsFollow = flag
	}

	userInfoList := &UserInfoList{
		List: tempList,
	}

	return userInfoList, nil

}
