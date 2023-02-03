package videoSvc

import (
	"douyin/model"
	"fmt"
	"time"
)

type FeedVideoList struct {
	VideoList []*model.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

func QueryFeedVideoList(userId int64, latestTime time.Time) (*FeedVideoList, error) {
	return NewFeedVideoList(userId, latestTime).Do()
}

func NewFeedVideoList(userId int64, latestTime time.Time) *QueryFeedVideoFlow {
	return &QueryFeedVideoFlow{userId: userId, latestTime: latestTime}
}

type QueryFeedVideoFlow struct {
	userId        int64
	latestTime    time.Time
	videos        []*model.Video
	nextTime      int64
	feedVideoList *FeedVideoList
}

func (qf *QueryFeedVideoFlow) Do() (*FeedVideoList, error) {
	isLogin := qf.checkParam()
	if err := qf.prepareFeedInfo(isLogin); err != nil {
		fmt.Println("prepareFeedInfo Error: ", err.Error())
		return nil, err
	}
	if err := qf.packageFeedInfo(isLogin); err != nil {
		fmt.Println("packageFeedInfo Error: ", err.Error())
		return nil, err
	}

	return qf.feedVideoList, nil
}

func (qf *QueryFeedVideoFlow) checkParam() bool {
	if qf.latestTime.IsZero() || time.Unix(0, 0).Equal(qf.latestTime) {
		qf.latestTime = time.Now()
	}
	if qf.userId <= 0 {
		return false
	}
	return true
}

func (qf *QueryFeedVideoFlow) prepareFeedInfo(isLogin bool) error {
	err := model.QueryVideoListByLimitAndTime(30, qf.latestTime, &qf.videos)
	if err != nil {
		return err
	}

	if !isLogin { //未登录
		qf.nextTime = time.Now().Unix() / 1e6
	} else { // 已登陆
		// TODO qf.nextTime = 最新视频的时间；根据用户id更新视频点赞状态
		/*size := len(qf.videos)
		nextTime := qf.videos[size-1].CreatedAt //有必要是这样的逻辑吗？
		qf.nextTime = nextTime.UnixNano() / 1e6
		*/
		qf.nextTime = time.Now().Unix() / 1e6
	}
	return nil
}

func (qf *QueryFeedVideoFlow) packageFeedInfo(isLogin bool) error {
	qf.feedVideoList = &FeedVideoList{qf.videos, qf.nextTime}
	return nil
}
