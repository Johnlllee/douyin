package userSvc

import (
	"douyin/model"
	"errors"
)

func QueryUserLogin(username, password string) (int64, error) {

	hasLogin := model.HasUserLoginByName(username)

	if hasLogin == false {
		return -1, errors.New("[UserSvc]: Username: [" + username + "] Not Exist.")
	}

	userLogin := &model.UserLogin{}
	err := model.QueryUserLogin(username, password, userLogin)

	if err != nil {
		return -1, err
	}

	return userLogin.UserInfoId, nil
}
