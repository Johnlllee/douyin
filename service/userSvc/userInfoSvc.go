package userSvc

import "douyin/model"

func QueryUserInfo(userid int64) (*model.UserInfo, error) {

	return &model.UserInfo{
		Id:       userid,
		Name:     "forTest",
		IsFollow: true,
	}, nil
}
