package ActionSvc

import (
	"douyin/model"
	"errors"
)

type UserInfoList struct {
	List []*model.UserInfo `json:"user_list"`
}

func GetFollowList(useridToQuery int64) (*UserInfoList, error) {

	exist, err := model.IsUserExistByUserId(useridToQuery)
	if exist == false {
		return nil, errors.New("[ActionSvc]: GetFollowList Userid To Query Not Exist In Db.")
	}

	if err != nil {
		return nil, errors.New("[ActionSvc]: GetFollowList" + err.Error())
	}

	var tempList []*model.UserInfo

	err = model.GetFollowListByUserId(useridToQuery, &tempList)

	if err != nil && err != model.ErrEmptyUserList {
		return nil, errors.New("[ActionSvc]: GetFollowList" + err.Error())
	}

	for _, v := range tempList {
		v.IsFollow = true
	}

	userInfoList := &UserInfoList{
		List: tempList,
	}

	return userInfoList, nil
}
