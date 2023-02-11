package model

import "gorm.io/gorm"

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

func IsUserExistByUserId(id int64) (bool, error) {
	var user UserInfo
	err := DB.Where("id=?", id).First(&user).Error
	if err != nil {
		return false, err
	}
	if user.Id == 0 {
		return false, nil
	}
	return true, nil
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

func AddUserFollow(userId, followId int64) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE user_infos SET follow_count=follow_count+1 WHERE id=?", userId).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE user_infos SET follower_count=follower_count+1 WHERE id=?", followId).Error; err != nil {
			return err
		}
		if err := tx.Exec("INSERT INTO `user_relations` (`user_info_id`, follow_id) VALUES (?,?)", userId, followId).Error; err != nil {
			return err
		}
		return nil
	})
}

func DeleteUserFollow(userId, followId int64) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE user_infos SET follow_count=follow_count-1 WHERE id=? AND follow_count>0", userId).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE user_infos SET follower_count=follower_count-1 WHERE id=? AND follower_count>0", followId).Error; err != nil {
			return err
		}
		if err := tx.Exec("DELETE FROM `user_relations` WHERE user_info_id=? AND follow_id=?", userId, followId).Error; err != nil {
			return err
		}
		return nil
	})
}

func GetFollowListByUserId(userId int64, userList *[]*UserInfo) error {
	if userList == nil {
		return ErrIvdPtr
	}
	if err := DB.Raw("SELECT u.* FROM user_relations r, user_infos u WHERE r.user_info_id=? AND r.follow_id=u.id", userId).Scan(userList).Error; err != nil {
		return err
	}
	if len(*userList) == 0 || (*userList)[0].Id == 0 {
		return ErrEmptyUserList
	}
	return nil
}

func GetFollowerListByUserId(userId int64, userList *[]*UserInfo) error {
	if userList == nil {
		return ErrIvdPtr
	}
	if err := DB.Raw("SELECT u.* FROM user_relations r, user_infos u WHERE r.follow_id=? AND r.user_info_id=u.id", userId).Scan(userList).Error; err != nil {
		return err
	}
	return nil
}
