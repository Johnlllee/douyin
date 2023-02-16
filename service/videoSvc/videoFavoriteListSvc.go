package videoSvc

import (
	"douyin/cache"
	"douyin/model"
	"errors"
)

type FavoriteVideoList struct {
	VideoList []*model.Video `json:"video_list,omitempty"`
}

type QueryFavoriteListFlow struct {
	userId int64
	videos []*model.Video

	favoriteVideoList *FavoriteVideoList
}

func QueryFavoriteList(userId int64) (*FavoriteVideoList, error) {
	return NewQueryFavoriteListFlow(userId).Do()
}

func NewQueryFavoriteListFlow(userId int64) *QueryFavoriteListFlow {
	return &QueryFavoriteListFlow{userId: userId}
}

func (qf *QueryFavoriteListFlow) Do() (*FavoriteVideoList, error) {
	if err := qf.checkParam(); err != nil {
		return nil, err
	}

	if err := qf.prepareInfo(); err != nil {
		return nil, err
	}

	if err := qf.packageInfo(); err != nil {
		return nil, err
	}

	return qf.favoriteVideoList, nil
}

func (qf *QueryFavoriteListFlow) checkParam() error {
	if qf.userId <= 0 {
		return errors.New("query favorite list flow error : user id <= 0")
	}

	exist, err := model.IsUserExistByUserId(qf.userId)
	if err != nil {
		return err
	}

	if !exist {
		return errors.New("query favorite list flow error : user does not exist")
	}

	return nil
}

func (qf *QueryFavoriteListFlow) prepareInfo() error {

	p := cache.NewRedisProxy()
	videoIds, err := p.GetFavoriteVideoIds(qf.userId)
	if err != nil {
		return err
	}

	err = model.QueryVideoListByVideoIds(videoIds, &qf.videos)
	if err != nil {
		return err
	}

	err = SetVideoListInfo(qf.userId, &qf.videos, false)
	if err != nil {
		return err
	}

	return nil
}

func (qf *QueryFavoriteListFlow) packageInfo() error {
	qf.favoriteVideoList = &FavoriteVideoList{
		VideoList: qf.videos,
	}
	return nil
}
