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
		StatusMessage := "Invalid Title"
		SendPostResponse(c, 1, StatusMessage)
		return
	}

	file, err := c.FormFile("data")
	if err != nil {
		SendPostResponse(c, 1, err.Error())
		return
	}

	suffix := filepath.Ext(file.Filename)
	if _, ok := videoType[suffix]; !ok {
		StatusMessage := "Unsupposed Video Type: " + "(" + suffix + ")"
		SendPostResponse(c, 1, StatusMessage)
		return
	}

	videoPath, coverPath := GenerateStaticPath(userIdInt, suffix)
	if exist := PathExists("./static"); !exist {
		err = os.Mkdir("./static", os.ModePerm)
		if err != nil {
			SendPostResponse(c, 1, err.Error())
			return
		}
	}

	saveVideoPath := filepath.Join("./static", videoPath) // TODO 路径写在config里
	saveCoverPath := filepath.Join("./static", coverPath)

	err = c.SaveUploadedFile(file, saveVideoPath)
	if err != nil {
		SendPostResponse(c, 1, err.Error())
		return
	}

	err = SaveImageFromVideo(saveVideoPath, saveCoverPath)
	if err != nil {
		SendPostResponse(c, 1, err.Error())
		return
	}

	err = videoSvc.PostVideo(userIdInt, title, videoPath, coverPath)
	if err != nil {
		SendPostResponse(c, 1, err.Error())
		return
	}

	StausMessage := "Successfully Upload File " + file.Filename
	SendPostResponse(c, 0, StausMessage)
	return
}

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

func SendPostResponse(c *gin.Context, statusCode int32, statusMessage string) {
	if c == nil {
		fmt.Println("SendPostResponse Fail: Context is nil")
		return
	}
	c.JSON(http.StatusOK, handler.PostVideoResponse{
		handler.CommonResponse{
			StatusCode: statusCode,
			StatusMsg:  statusMessage,
		},
	})
	return
}

func GenerateStaticPath(userId int64, videoTypeSuffix string) (string, string) {
	postTime := int(time.Now().Unix())
	videoPath := strconv.Itoa(int(userId)) + "_" + strconv.Itoa(postTime) + videoTypeSuffix
	coverPath := strconv.Itoa(int(userId)) + "_" + strconv.Itoa(postTime) + ".jpg"
	return videoPath, coverPath
}
