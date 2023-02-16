package videoSvc

import (
	"douyin/cache"
	"douyin/model"
)

// 输入当前用户的userId，当前接口需要输出的videoList
func SetVideoListInfo(userId int64, videos *[]*model.Video, reverse bool) error {
	p := cache.NewRedisProxy()
	var i int
	var err error
	for i = 0; i < len(*videos)/2; i++ {
		j := len(*videos) - i - 1

		if userId > 0 {
			//设置点赞状态, 非登陆状态不需要设置
			(*videos)[i].IsFavorite, err = p.GetFavoriteStatus(userId, (*videos)[i].Id)
			if err != nil {
				return err
			}

			(*videos)[j].IsFavorite, err = p.GetFavoriteStatus(userId, (*videos)[j].Id)
			if err != nil {
				return err
			}
		}

		err = SetVideoAuthor(userId, (*videos)[i])
		if err != nil {
			return err
		}

		err = SetVideoAuthor(userId, (*videos)[j])
		if err != nil {
			return err
		}

		if reverse { // 翻转video
			(*videos)[i], (*videos)[j] = (*videos)[j], (*videos)[i]
		}
	}

	if len(*videos)%2 == 1 {
		if userId > 0 {
			(*videos)[i].IsFavorite, err = p.GetFavoriteStatus(userId, (*videos)[i].Id)
			if err != nil {
				return err
			}
		}

		err = SetVideoAuthor(userId, (*videos)[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func SetVideoAuthor(userId int64, video *model.Video) error {
	author := model.UserInfo{}
	err := model.QueryUserInfoByUserId((*video).UserInfoId, &author)
	if err != nil {
		return err
	}

	isFollow := false
	if userId > 0 {
		p := cache.NewRedisProxy()
		isFollow, err = p.GetFollowStatus(userId, (*video).UserInfoId)
		if err != nil {
			return err
		}
	}
	author.IsFollow = isFollow
	(*video).Author = author

	return nil
}
