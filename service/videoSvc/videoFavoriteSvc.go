package videoSvc

import (
	"douyin/cache"
	"douyin/model"
	"errors"
	"sync"
)

type AddFavoriteFlow struct {
	userId     int64
	videoId    int64
	actionType int64
}

func AddFavorite(userId int64, videoId int64, actionType int64) error {
	return NewAddFavoriteFlow(userId, videoId, actionType).Do()
}

func NewAddFavoriteFlow(userId int64, videoId int64, actionType int64) *AddFavoriteFlow {
	return &AddFavoriteFlow{userId: userId, videoId: videoId, actionType: actionType}
}

func (af *AddFavoriteFlow) Do() error {
	if err := af.checkParam(); err != nil {
		return err
	}

	if err := af.prepareAddInfo(); err != nil {
		return err
	}

	return nil
}

func (af *AddFavoriteFlow) checkParam() error {
	if af.userId <= 0 {
		return errors.New("AddFavorite userId <= 0")
	}

	exist, err := model.IsUserExistByUserId(af.userId)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("favoriteSvc error: user not exist")
	}

	exist, err = model.IsVideoExistByVideoId(af.videoId)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("favoriteSvc error: video not exist")
	}

	if af.actionType != 1 && af.actionType != 2 {
		return errors.New("Unsupported ActionType")
	}

	return nil
}

func (af *AddFavoriteFlow) prepareAddInfo() error {
	var actionType bool
	if af.actionType == 1 {
		actionType = true
	} else {
		actionType = false
	}

	//TODO 将userId和videoId加入user_favor_videos表，改变video的获赞总数, 将 go程
	wg := sync.WaitGroup{}
	wg.Add(2)

	var setFavorErr error
	//缓存操作
	go func() {
		defer wg.Done()
		err := cache.NewRedisProxy().SetFavorite(af.userId, af.videoId, actionType)
		if err != nil {
			setFavorErr = err
			return
		}
	}()

	var addCommentErr error
	//修改video表
	go func() {
		defer wg.Done()
		//TODO mysql 数据库操作videos表

		if actionType {
			//加
			err := model.AddFavoriteCountByVideoId(af.videoId)
			if err != nil {
				addCommentErr = err
				return
			}
		} else {
			//减
			err := model.MinusFavoriteCountByVideoId(af.videoId)
			if err != nil {
				addCommentErr = err
				return
			}

		}

	}()
	wg.Wait()

	if setFavorErr != nil {
		return setFavorErr
	}
	if addCommentErr != nil {
		return addCommentErr
	}

	return nil

}
