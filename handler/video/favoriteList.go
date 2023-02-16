package video

import (
	"douyin/handler"
	"douyin/service/videoSvc"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FavoriteListHandler(c *gin.Context) {
	userId, ok := c.Get("userid")
	if !ok {
		SendFavoriteListResponse(c, 1, "FavoriteListHandler get userId failed", nil)
		return
	}

	userIdInt, ok := userId.(int64)
	if !ok {
		SendFavoriteListResponse(c, 1, "FavoriteListHandler fail to parse userId", nil)
		return
	}

	videoList, err := videoSvc.QueryFavoriteList(userIdInt)
	if err != nil {
		SendFavoriteListResponse(c, 1, err.Error(), nil)
		return
	}

	SendFavoriteListResponse(c, 0, "Successfully send favorite list", videoList)

}

func SendFavoriteListResponse(c *gin.Context, statusCode int32, statusMessage string, videoList *videoSvc.FavoriteVideoList) {
	if c == nil {
		fmt.Println("SendFavorListResponse error: c == nil")
		return
	}
	c.JSON(http.StatusOK, handler.FavoriteListResponse{
		handler.CommonResponse{
			StatusCode: statusCode,
			StatusMsg:  statusMessage,
		},
		videoList,
	})
	return
}
