package video

import (
	"douyin/handler"
	"douyin/service/videoSvc"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PublishedVideoListHandler(c *gin.Context) {
	//userId, ok := c.GetQuery("user_id")
	userId, ok := c.Get("userid")

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

	userIdInt, ok := userId.(int64)
	if !ok {
		c.JSON(http.StatusOK, handler.PublishedVideoResponse{
			handler.CommonResponse{
				1,
				"PublishedVideoListHandler Fails To Parse Id",
			},
			nil,
		})
		return
	}
	//userIdInt, err := strconv.ParseInt(userId, 10, 64)
	/*if err != nil {
		c.JSON(http.StatusOK, handler.PublishedVideoResponse{
			handler.CommonResponse{
				1,
				err.Error(),
			},
			nil,
		})
		return
	}*/

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
