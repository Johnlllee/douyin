package userSvc

import (
	"douyin/model"
	"douyin/service"
	"errors"
)

func QueryUserLogin(username, password string) (int64, error) {

	service.UserServiceManager.mutex.Lock()
	defer service.UserServiceManager.mutex.Unlock()

	hasLogin := model.HasUserLoginByName(username)

	if hasLogin == true {
		return -1, errors.New("[UserSvc]: Username: [" + username + "] Already Login.")
	}

	userLogin := &model.UserLogin{}
	err := model.QueryUserLogin(username, password, userLogin)

	if err != nil {
		return -1, err
	}

	return userLogin.Id, nil
}
