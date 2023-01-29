package video

import (
	"douyin/handler"
	"douyin/service/videoSvc"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

var videoType = map[string]struct{}{
	".mp4":  {},
	".avi":  {},
	".wmv":  {},
	".flv":  {},
	".mpeg": {},
	".mov":  {},
}

func PostVideoHandler(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userIdInt := userId.(int64)

	title := c.PostForm("title")
	if title == "" {
		c.JSON(http.StatusOK, handler.PostVideoResponse{
			handler.CommonResponse{
				1,
				"Invalid Title",
			},
		})
		return
	}

	file, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, handler.PostVideoResponse{
			handler.CommonResponse{
				1,
				err.Error(),
			},
		})
		return
	}

	suffix := filepath.Ext(file.Filename)
	if _, ok := videoType[suffix]; !ok {
		c.JSON(http.StatusOK, handler.PostVideoResponse{
			handler.CommonResponse{
				1,
				"Unsupposed Video Type" + suffix,
			},
		})
		return
	}

	postTime := int(time.Now().Unix())
	videoPath := strconv.Itoa(int(userIdInt)) + strconv.Itoa(postTime) + suffix
	savePath := filepath.Join("./static", videoPath)
	err = c.SaveUploadedFile(file, savePath)
	if err != nil {
		c.JSON(http.StatusOK, handler.PostVideoResponse{
			handler.CommonResponse{
				1,
				err.Error(),
			},
		})
		return
	}
	coverPath := "" //TODO 从上传的视频中选中一帧作为封面，并将其存在本地
	err = videoSvc.PostVideo(userIdInt, title, videoPath, coverPath)
	if err != nil {
		c.JSON(http.StatusOK, handler.PostVideoResponse{
			handler.CommonResponse{
				1,
				err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, handler.PostVideoResponse{
		handler.CommonResponse{
			0,
			"Successfully Upload File " + file.Filename,
		},
	})
	return
}

/*func SaveImageFromVideo(name string, isDebug bool) error {
	v2i := NewVideo2Image()
	if isDebug {
		v2i.Debug()
	}
	v2i.InputPath = filepath.Join(config.Info.StaticSourcePath, name+defaultVideoSuffix)
	v2i.OutputPath = filepath.Join(config.Info.StaticSourcePath, name+defaultImageSuffix)
	v2i.FrameCount = 1
	queryString, err := v2i.GetQueryString()
	if err != nil {
		return err
	}
	return v2i.ExecCommand(queryString)
}*/
