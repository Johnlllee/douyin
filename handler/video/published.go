package video

import (
	"douyin/handler"
	"douyin/service/videoSvc"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PublishedVideoListHandler(c *gin.Context) {
	//userId, ok := c.GetQuery("user_id")
	userId, ok := c.Get("userid")

	if !ok {
		SendPublishedResponse(c, 1, "PublishedVideoListHandler Fails To Get userId", nil)
		return
	}

	userIdInt, ok := userId.(int64)
	if !ok {
		SendPublishedResponse(c, 1, "PublishedVideoListHandler Fails To Parse Id", nil)
		return
	}

	videoList, err := videoSvc.QueryPublishedVideoList(userIdInt)
	if err != nil {
		SendPublishedResponse(c, 1, err.Error(), nil)
		return
	}
	SendPublishedResponse(c, 0, "Successfully Get Published Video List", videoList)

	return

}

func SendPublishedResponse(c *gin.Context, statusCode int32, statusMessage string, videoList *videoSvc.PublishedVideoList) {
	if c == nil {
		fmt.Println("SendPublishedResponse Fail: Context == nil")
		return
	}
	c.JSON(http.StatusOK, handler.PublishedVideoResponse{
		handler.CommonResponse{
			statusCode,
			statusMessage,
		},
		videoList,
	})
	return
}
