package model

// jojoleee 用户信息表
type UserInfo struct {
	Id            int64      `json:"id"`
	Name          string     `json:"name"`
	FollowCount   int64      `json:"follow_count"`
	FollowerCount int64      `json:"follower_count"`
	IsFollow      bool       `json:"is_follow"`
	User          *UserLogin `json:"-"`
}

func AddUserInfo(userinfo *UserInfo) error {
	/**
	 * @Author jojoleee
	 * @Description 提供给用户注册
	 * @Param [userinfo]
	 * @return
	 **/
	if userinfo == nil {
		return ErrIvdPtr
	}
	return DB.Create(userinfo).Error
}

func QueryUserInfoByUserId(id int64, userinfo *UserInfo) error {
	/**
	 * @Author jojoleee
	 * @Description //TODO
	 * @Param [id, userinfo]
	 * @return
	 **/
	if userinfo == nil {
		return ErrIvdPtr
	}
	DB.Where("id=?", id).First(userinfo)
	if userinfo.Id == 0 {
		return ErrUserNotExist
	}
	return nil
}

func QueryUserInfoByName(username string, userinfo *UserInfo) error {
	/**
	 * @Author jojoleee
	 * @Description //TODO
	 * @Param [username, userinfo]
	 * @return
	 **/
	if userinfo == nil {
		return ErrIvdPtr
	}
	DB.Where("name=?", username).First(userinfo)
	if userinfo.Id == 0 {
		return ErrUserNotExist
	}
	return nil
}

func QueryLastUserId(id *int64) error {
	var userinfo UserInfo
	DB.Last(&userinfo)
	*id = userinfo.Id
	return nil
}
