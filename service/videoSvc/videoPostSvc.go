package videoSvc

import (
	"douyin/model"
	"errors"
)

type PostVideoFlow struct {
	userId    int64
	title     string
	videoPath string
	coverPath string
}

func PostVideo(userId int64, title string, videoPath string, coverPath string) error { // 参数来自data byte解析
	return NewPostVideoFlow(userId, title, videoPath, coverPath).Do()
}

func NewPostVideoFlow(userId int64, title string, videoPath string, coverPath string) *PostVideoFlow {
	return &PostVideoFlow{
		userId,
		title,
		videoPath,
		coverPath,
	}
}

func (pf *PostVideoFlow) Do() error {
	if err := pf.checkParam(); err != nil {
		return err
	}
	if err := pf.preparePostInfo(); err != nil {
		return err
	}
	if err := pf.postVideo(); err != nil {
		return err
	}
	return nil
}

func (pf *PostVideoFlow) checkParam() error {
	if pf.userId <= 0 {
		return errors.New("PostVideo userID <= 0")
	}
	return nil
}

func (pf *PostVideoFlow) preparePostInfo() error {
	pf.videoPath = GetFileUrl(pf.videoPath)
	pf.coverPath = GetFileUrl(pf.coverPath)
	return nil
}

func (pf *PostVideoFlow) postVideo() error {
	video := &model.Video{UserInfoId: pf.userId, PlayUrl: pf.videoPath, CoverUrl: pf.coverPath, Title: pf.title}
	err := model.AddVideo(video)
	if err != nil {
		return err
	}
	return nil
}

//TODO 完善config.Info

func GetFileUrl(fileName string) string {
	if fileName == "" { //TODO coverPath 尚未解决
		fileName = "Debug"
	}
	//base := fmt.Sprintf("http://%s:%d/static/%s", config.Info.IP, config.Info.Port, fileName)
	base := ""
	return base
}
