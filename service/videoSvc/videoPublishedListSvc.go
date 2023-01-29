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
	return nil
}

func (pf *PublishedVideoFlow) preparePublishedInfo() error {

	err := model.QueryVideoListByUserId(pf.userId, &pf.videos)
	if err != nil {
		return err
	}
	pf.videoList = &PublishedVideoList{pf.videos}

	return nil
}
