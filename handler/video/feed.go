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
		fmt.Println("FeedVideoHandle GetFeedVideoList Error: ", err.Error())
		response := handler.FeedResponse{
			handler.CommonResponse{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
			nil,
		}
		c.JSON(http.StatusOK, response)
		return
	}
	response := handler.FeedResponse{
		handler.CommonResponse{
			StatusCode: 0,
		},
		videoList,
	}
	c.JSON(http.StatusOK, response)
	return

}

func GetFeedVideoList(c *gin.Context, isLogin bool) (*videoSvc.FeedVideoList, error) {
	if c == nil {
		return nil, errors.New("Context is nil")
	}

	//var videoList *video.FeedVideoList

	latestTimeRaw := c.Query("latest_time")
	var latestTime time.Time
	latestTimeInt, err := strconv.ParseInt(latestTimeRaw, 10, 64)
	if err != nil {
		return nil, err
	}
	latestTime = time.Unix(0, latestTimeInt*1e6) //前端传来的时间戳以ms为单位

	var videoList *videoSvc.FeedVideoList
	if isLogin == false { //未登陆状态
		videoList, err = videoSvc.QueryFeedVideoList(0, latestTime)
		if err != nil {
			return nil, errors.New("failed to QueryFeedVideoList (isLogin false)")
		}
	} else {
		token := c.Query("token")
		claim, ok := middleware.ParseToken(token)
		if !ok {
			return nil, errors.New("Feed Video Handler Parse Token Error")
		}

		if claim.ExpiresAt < time.Now().Unix() {
			return nil, errors.New("Token Expired")
		}
		//claim.

		videoList, err = videoSvc.QueryFeedVideoList(claim.UserId, latestTime)

	}

	return videoList, nil

}
