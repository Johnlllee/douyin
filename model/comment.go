package model

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	Id         int64     `json:"id"`
	UserInfoId int64     `json:"-"`
	VideoId    int64     `json:"-"`
	User       UserInfo  `json:"user" gorm:"-"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"-"`
}

func QueryCommentByCommentId(id int64, comment *Comment) error {
	/**
	 * @Author jojoleee
	 * @Description 根据评论Id寻找评论
	 * @Param [id, comment]
	 * @return
	 **/
	if comment == nil {
		return ErrIvdPtr
	}
	return DB.Model(&Comment{}).Where("id=?", id).First(comment).Error
}

func IsCommentExistByCommentId(id int64) (bool, error) {
	var comment Comment
	err := QueryCommentByCommentId(id, &comment)
	if err != nil {
		return false, err
	}
	if comment.Id == 0 {
		return false, nil
	}
	return true, nil
}

func QueryCommentListByVideoId(videoId int64, comments *[]*Comment) error {
	if comments == nil {
		return ErrIvdPtr
	}
	return DB.Model(&Comment{}).Where("video_id=?", videoId).
		//Select([]string{"id", "user", "content", "created_at"}).
		Find(comments).Error
}

func CreateCommentAndUpdateCount(comment *Comment) error {
	/**
	 * @Author jojoleee
	 * @Description 新增评论并使视频评论数加一
	 * @Param [comment]
	 * @return
	 **/
	if comment == nil {
		return ErrIvdPtr
	}
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(comment).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE videos v SET v.comment_count = v.comment_count+1 WHERE v.id=?", comment.VideoId).Error; err != nil {
			return err
		}
		return nil
	})
}

func DeleteCommentAndUpdateCount(commentId, videoId int64) error {
	/**
	 * @Author jojoleee
	 * @Description 删除评论并使视频评论数减一
	 * @Param [comment]
	 * @return
	 **/
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE FROM comments WHERE id=?", commentId).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE videos v SET v.comment_count = v.comment_count-1 WHERE v.id=? AND v.comment_count>0", videoId).Error; err != nil {
			return err
		}
		return nil
	})
}
