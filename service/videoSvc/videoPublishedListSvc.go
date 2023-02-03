package videoSvc

import (
	"douyin/model"
	"errors"
)

type PublishedVideoList struct {
	VideoList []*model.Video `json:"video_list,omitempty"`
}

func QueryPublishedVideoList(userId int64) (*PublishedVideoList, error) {
	return NewQueryPublishedVideoFlow(userId).Do()
}

type PublishedVideoFlow struct {
	userId    int64
	videos    []*model.Video
	videoList *PublishedVideoList
}

func NewQueryPublishedVideoFlow(userId int64) *PublishedVideoFlow {
	return &PublishedVideoFlow{
		userId:    userId,
		videos:    nil,
		videoList: nil,
	}
}

func (pf *PublishedVideoFlow) Do() (*PublishedVideoList, error) {
	if err := pf.checkParam(); err != nil {
		return nil, err
	}
	if err := pf.preparePublishedInfo(); err != nil {
		return nil, err
	}
	return pf.videoList, nil

}

func (pf *PublishedVideoFlow) checkParam() error {
	if pf.userId <= 0 {
		return errors.New("Published Videos UserId <= 0")
	}
	//TODO 判断userId是否在数据库中存在
	userInfo := new(model.UserInfo)
	err := model.QueryUserInfoByUserId(pf.userId, userInfo)
	if err != nil {
		return err
	}
	return nil
}

func (pf *PublishedVideoFlow) preparePublishedInfo() error {

	err := model.QueryVideoListByUserId(pf.userId, &pf.videos) //TODO 判断videoList是否存在
	if err != nil {
		return err
	}
	pf.videoList = &PublishedVideoList{pf.videos}

	return nil
}
