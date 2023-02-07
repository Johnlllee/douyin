package video

import (
	"douyin/handler"
	"douyin/middleware"
	"douyin/service/videoSvc"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func FeedVideoHandler(c *gin.Context) {
	_, ok := c.GetQuery("token")
	var isLogin bool
	if !ok { // 非登陆状态
		isLogin = false
	} else { // 登陆状态
		isLogin = true
	}
	videoList, err := GetFeedVideoList(c, isLogin)
	if err != nil {
		StatusMsg := "FeedVideoHandle GetFeedVideoList Error: " + err.Error()
		SendFeedResponse(c, 1, StatusMsg, nil)
		return
	}

	SendFeedResponse(c, 0, "Successfully Get Feed Videolist", videoList)
	return

}

func GetFeedVideoList(c *gin.Context, isLogin bool) (*videoSvc.FeedVideoList, error) {
	if c == nil {
		return nil, errors.New("Context is nil")
	}

	latestTimeRaw := c.Query("latest_time")
	//fmt.Println("LTR: ", latestTimeRaw)
	var latestTime time.Time
	latestTimeInt, err := strconv.ParseInt(latestTimeRaw, 10, 64)
	if err != nil {
		//return nil, err
		latestTime = time.Unix(0, latestTimeInt*1e6) //前端传来的时间戳以ms为单位
	}

	var videoList *videoSvc.FeedVideoList
	if isLogin == false { //未登陆状态
		videoList, err = videoSvc.QueryFeedVideoList(0, latestTime)
		if err != nil {
			return nil, errors.New("failed to QueryFeedVideoList (isLogin false)")
		}
		return videoList, nil
	}

	token := c.Query("token")
	claim, ok := middleware.ParseToken(token)
	if !ok {
		return nil, errors.New("Feed Video Handler Parse Token Error")
	}

	if claim.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("Token Expired")
	}

	videoList, err = videoSvc.QueryFeedVideoList(claim.UserId, latestTime)
	//videoList, err = videoSvc.QueryFeedVideoList(1, latestTime)
	if err != nil {
		return nil, errors.New("failed to QueryFeedVideoList (isLogin true)")
	}

	return videoList, nil

}

func SendFeedResponse(c *gin.Context, statusCode int32, statusMessage string, videoList *videoSvc.FeedVideoList) {
	if c == nil {
		fmt.Println("SendFeedResponse Fail: Context == nil")
		return
	}
	c.JSON(http.StatusOK, handler.FeedResponse{
		CommonResponse: handler.CommonResponse{
			StatusCode: statusCode,
			StatusMsg:  statusMessage,
		},
		FeedVideoList: videoList,
	})
	return
}
