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
