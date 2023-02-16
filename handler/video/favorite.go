package video

import (
	"douyin/handler"
	"douyin/service/videoSvc"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddFavoriteHandler(c *gin.Context) {
	userIdRaw, ok := c.Get("userid")
	if !ok {
		SendAddFavoriteResponse(c, 1, "AddFavorite Fail to Get userId")
	}

	userIdInt, ok := userIdRaw.(int64)
	if !ok {
		SendAddFavoriteResponse(c, 1, "AddFavorite Fail to Parse userId")
		return
	}

	videoId, ok := c.GetQuery("video_id")
	if !ok {
		SendAddFavoriteResponse(c, 1, "AddFavorite Fail to Parse video_id")
		return
	}
	videoIdInt, err := strconv.ParseInt(videoId, 10, 64)
	if err != nil {
		SendAddFavoriteResponse(c, 1, err.Error())
		return
	}

	actionType, ok := c.GetQuery("action_type")
	if !ok {
		SendAddFavoriteResponse(c, 1, "AddFavorite Fail to Parse action_type")
		return
	}

	actionTypeInt, err := strconv.ParseInt(actionType, 10, 64)
	if err != nil {
		SendAddFavoriteResponse(c, 1, err.Error())
	}

	err = videoSvc.AddFavorite(userIdInt, videoIdInt, actionTypeInt)
	if err != nil {
		SendAddFavoriteResponse(c, 1, err.Error())
		return
	}

	SendAddFavoriteResponse(c, 0, "Successfully Conduct AddFavorite")
	return
}

func SendAddFavoriteResponse(c *gin.Context, statusCode int32, statusMessage string) {
	if c == nil {
		fmt.Println("SendAddFavoriteResponse Fail: Context == nil")
		return
	}
	c.JSON(http.StatusOK, handler.AddFavoriteResponse{
		CommonResponse: handler.CommonResponse{
			StatusCode: statusCode,
			StatusMsg:  statusMessage,
		},
	})
	return
}
