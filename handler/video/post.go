package video

import (
	"douyin/handler"
	"douyin/service/videoSvc"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/exec"
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
	//c.Set("user_id", int64(1))    //测试用
	userId, _ := c.Get("userid") //经过jwt鉴权解析的userid
	//userId, _ := c.GetQuery()

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
				"Unsupposed Video Type: " + "(" + suffix + ")",
			},
		})
		return
	}

	postTime := int(time.Now().Unix())
	videoPath := strconv.Itoa(int(userIdInt)) + "_" + strconv.Itoa(postTime) + suffix
	coverPath := strconv.Itoa(int(userIdInt)) + "_" + strconv.Itoa(postTime) + ".jpg"

	if exist := PathExists("./static"); !exist {
		err = os.Mkdir("./static", os.ModePerm)
		if err != nil {
			c.JSON(http.StatusOK, handler.PostVideoResponse{
				handler.CommonResponse{
					1,
					err.Error(),
				},
			})
			return
		}
	}

	saveVideoPath := filepath.Join("./static", videoPath) // TODO 路径写在config里
	saveCoverPath := filepath.Join("./static", coverPath)

	err = c.SaveUploadedFile(file, saveVideoPath)
	if err != nil {
		c.JSON(http.StatusOK, handler.PostVideoResponse{
			handler.CommonResponse{
				1,
				err.Error(),
			},
		})
		return
	}

	err = SaveImageFromVideo(saveVideoPath, saveCoverPath)
	if err != nil {
		c.JSON(http.StatusOK, handler.PostVideoResponse{
			handler.CommonResponse{
				1,
				err.Error(),
			},
		})
		return
	}

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
// TODO 工具放在util里
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}

func SaveImageFromVideo(inputPath string, outputPath string) error {
	//ffmpeg -i sd1671527681_2.mp4 -f image2 -frames:v 1 -y outputpath.jpg
	if inputPath == "" || outputPath == "" {
		return errors.New("Save Image From Video Error: path is invalid")
	}

	args := []string{"-i", inputPath, "-f", "image2", "-frames:v", "1", "-y", outputPath}

	cmd := exec.Command("ffmpeg", args...)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
