package video

import (
	"douyin/handler"
	"douyin/service/videoSvc"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PublishedVideoListHandler(c *gin.Context) {
	userId, ok := c.GetQuery("user_id")
	if !ok {
		c.JSON(http.StatusOK, handler.PublishedVideoResponse{
			handler.CommonResponse{
				1,
				"PublishedVideoListHandler Fails To Get userId",
			},
			nil,
		})
		return
	}
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, handler.PublishedVideoResponse{
			handler.CommonResponse{
				1,
				err.Error(),
			},
			nil,
		})
		return
	}

	videoList, err := videoSvc.QueryPublishedVideoList(userIdInt)
	if err != nil {
		c.JSON(http.StatusOK, handler.PublishedVideoResponse{
			handler.CommonResponse{
				1,
				err.Error(),
			},
			nil,
		})
		return
	}

	c.JSON(http.StatusOK, handler.PublishedVideoResponse{
		handler.CommonResponse{
			0,
			"Successfully Get Published Video List",
		},
		videoList,
	})

	return

}
