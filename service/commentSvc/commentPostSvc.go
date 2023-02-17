package commentSvc

import (
	"douyin/model"
	"errors"
)

const (
	CREATE = 1
	DELETE = 2
)

var (
	ErrObjNotExist      = errors.New("[Svc]对象不存在")
	ErrCommentListEmpty = errors.New("[Svc]评论列表为空")
)

type Response struct {
	//个人觉得需要从数据库中拿到相应的数据[lzy]
	Id         int64          `json:"id"`
	User       model.UserInfo `json:"user" gorm:"-"`
	Content    string         `json:"content"`
	CreateDate string         `json:"create_date"`
}

func CreateComment(videoId, userId int64, commentText string, comment *Response) error {
	videoExist, _ := model.IsVideoExistByVideoId(videoId)
	userExist, _ := model.IsUserExistByUserId(userId)
	if !userExist || !videoExist {
		return ErrObjNotExist
	}

	err := model.QueryUserInfoByUserId(userId, &comment.User)
	if err != nil {
		return err
	}

	commentToDB := model.Comment{UserInfoId: userId, VideoId: videoId, User: comment.User, Content: commentText}
	err = model.CreateCommentAndUpdateCount(&commentToDB)
	if err != nil {
		return err
	}

	comment.Id = commentToDB.Id
	comment.Content = commentText
	comment.CreateDate = commentToDB.CreatedAt.String()

	return nil
}

func DeleteComment(commentId, videoId int64, comment *Response) error {
	commentExist, _ := model.IsCommentExistByCommentId(commentId)
	videoExist, _ := model.IsVideoExistByVideoId(videoId)
	if !commentExist || !videoExist {
		return ErrObjNotExist
	}

	commentFromDB := model.Comment{}
	err := model.QueryCommentByCommentId(commentId, &commentFromDB)
	if err != nil {
		return err
	}
	comment.Content = commentFromDB.Content
	//[lzy]不知道为什么user信息不能显示出来,但其实可能也不需要返回这个User信息，因为是Optional
	comment.User = commentFromDB.User
	comment.CreateDate = commentFromDB.CreatedAt.String()
	comment.Id = commentId

	//[lzy]貌似删除不需要区分用户，因为可能在APP中看到的如果是别人的评论，并不会允许做删除
	err = model.DeleteCommentAndUpdateCount(commentId, videoId)
	if err != nil {
		return err
	}

	return nil
}
