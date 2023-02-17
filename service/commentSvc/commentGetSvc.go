package commentSvc

import "douyin/model"

type FeedCommentList struct {
	CommentList []*model.Comment `json:"comment_list,omitempty"`
}

func GetCommentList(videoId int64) (*FeedCommentList, error) {
	videoExist, _ := model.IsVideoExistByVideoId(videoId)
	if !videoExist {
		return nil, ErrObjNotExist
	}

	comments := make([]*model.Comment, 0)
	err := model.QueryCommentListByVideoId(videoId, &comments)
	if err != nil {
		return nil, err
	}
	err = QueryUserByComment(&comments)
	if err != nil {
		return nil, err
	}
	//for i:=0; i<len(comments); i++ {
	//	feed
	//}

	//[lzy] packData 不知道怎么吧comments-->Response,目前就用的comments

	return &FeedCommentList{CommentList: comments}, nil
}

func QueryUserByComment(comments *[]*model.Comment) error {
	n := len(*comments)
	if comments == nil || n == 0 {
		return ErrCommentListEmpty
	}
	for _, v := range *comments {
		_ = model.QueryUserInfoByUserId(v.UserInfoId, &v.User)
		v.CreateDate = v.CreatedAt.Format("1-2")
	}
	return nil
}
