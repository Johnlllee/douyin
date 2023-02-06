package userSvc

import "douyin/model"

func QueryUserInfo(userid int64) (*model.UserInfo, error) {

	userinfo := &model.UserInfo{}

	err := model.QueryUserInfoByUserId(userid, userinfo)

	if err != nil {
		return nil, err
	}

	return userinfo, nil

}

func IsUserExistByUserId(userid int64) (bool, error) {
	exist, err := model.IsUserExistByUserId(userid)
	if err != nil {
		return false, err
	}
	return exist, nil
}
