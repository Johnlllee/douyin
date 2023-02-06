package comment

import (
	"douyin/handler"
	"douyin/service/commentSvc"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PostCommentHandler(c *gin.Context) {
	newProxyPostCommentHandler(c).process()
}

func newProxyPostCommentHandler(c *gin.Context) *ProxyPostCommentHandler {
	return &ProxyPostCommentHandler{Context: c}
}

type ProxyPostCommentHandler struct {
	*gin.Context

	videoId     int64
	userId      int64
	commentId   int64
	actionType  int64
	commentText string
}

func (p *ProxyPostCommentHandler) process() {
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

	var outComment commentSvc.Response

	switch p.actionType {
	case commentSvc.CREATE:
		err := commentSvc.CreateComment(p.videoId, p.userId, p.commentText, &outComment)
		if err != nil {
			p.SendErr("failed to create comment: " + err.Error())
			return
		}
	case commentSvc.DELETE:
		err := commentSvc.DeleteComment(p.commentId, p.videoId, &outComment)
		if err != nil {
			p.SendErr("failed to delete comment: " + err.Error())
			return
		}
	}

	p.SendOK(&outComment)
}

func (p *ProxyPostCommentHandler) ParseParam() error {
	/**
	 * @Author jojoleee
	 * @Description 解析POST COMMENT的参数，然后会传给commenthandler结构体
	 * @Param
	 * @return
	 **/
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

	rawActionType := p.Query("action_type")
	actionType, err := strconv.ParseInt(rawActionType, 10, 64)
	switch actionType {
	case commentSvc.CREATE:
		p.commentText = p.Query("comment_text")
	case commentSvc.DELETE:
		p.commentId, err = strconv.ParseInt(p.Query("comment_id"), 10, 64)
		if err != nil {
			return err
		}
	default:
		return errors.New("invalid action type")
	}
	p.actionType = actionType
	return nil
}

func (p *ProxyPostCommentHandler) SendOK(comment *commentSvc.Response) {
	p.JSON(http.StatusOK, handler.PostCommentResponse{
		CommonResponse: handler.CommonResponse{StatusCode: 0, StatusMsg: "成功添加/删除评论"},
		Response:       comment,
	})
}

func (p *ProxyPostCommentHandler) SendErr(msg string) {
	p.JSON(http.StatusOK, handler.CommonResponse{StatusCode: 1, StatusMsg: msg})
}
