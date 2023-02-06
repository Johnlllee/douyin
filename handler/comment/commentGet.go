package comment

import (
	"douyin/handler"
	"douyin/service/commentSvc"
	"douyin/service/userSvc"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCommentHandler(c *gin.Context) {
	newProxyGetCommentHandler(c).process()
}

func newProxyGetCommentHandler(c *gin.Context) *ProxyGetCommentHandler {
	return &ProxyGetCommentHandler{Context: c}
}

type ProxyGetCommentHandler struct {
	*gin.Context

	videoId int64
	userId  int64
}

func (p *ProxyGetCommentHandler) process() {
	/**
	 * @Author jojoleee
	 * @Description POST评论接口函数
	 * @Param
	 * @return
	 **/
	err := p.ParseParam()
	if err != nil {
		p.SendErr(err.Error())
		return
	}

	//addition we need to check user
	userExist, err := userSvc.IsUserExistByUserId(p.userId)
	if err != nil {
		p.SendErr(err.Error())
		return
	}
	if !userExist {
		p.SendErr("用户不存在")
		return
	}
	commentList, err := commentSvc.GetCommentList(p.videoId)
	if err != nil {
		p.SendErr("failed to get comment list: " + err.Error())
		return
	}

	p.SendOK(commentList)
}

func (p *ProxyGetCommentHandler) ParseParam() error {
	userId, _ := p.Get("userid")
	userIdInt, ok := userId.(int64)
	if !ok {
		return errors.New("failed to get userId")
	}
	p.userId = userIdInt

	rawVideoId := p.Query("video_id")
	videoId, err := strconv.ParseInt(rawVideoId, 10, 64)
	if err != nil {
		return err
	}
	p.videoId = videoId

	return nil
}

func (p *ProxyGetCommentHandler) SendOK(commentList *commentSvc.FeedCommentList) {
	p.JSON(http.StatusOK, handler.GetCommentResponse{
		CommonResponse:  handler.CommonResponse{StatusCode: 0, StatusMsg: "获得评论列表"},
		FeedCommentList: commentList,
	})
}

func (p *ProxyGetCommentHandler) SendErr(msg string) {
	p.JSON(http.StatusOK, handler.CommonResponse{StatusCode: 1, StatusMsg: msg})
}
