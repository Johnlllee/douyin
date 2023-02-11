package cache

import (
	"bytes"
	"context"
	"douyin/config"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
)

var rdb *redis.Client
var c = context.Background()

func InitCache() {
	initRDB()
}

func initRDB() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Info.RDB.IP, config.Info.RDB.Port),
		Password: "",
		DB:       0,
	})
}

type RedisProxy struct {
}

func NewRedisProxy() *RedisProxy {
	return &RedisProxy{}
}

// 设置缓存中的关注状态 userId为操作者Id， videoId为被关注者Id， isFavor为操作类型，true则为点赞，false为取消点赞
func (p *RedisProxy) SetFavorite(userId int64, videoId int64, isFavor bool) error {
	if userId <= 0 || videoId <= 0 {
		return errors.New("Redis SetFavorite Error: Id <= 0")
	}

	userIdString := strconv.Itoa(int(userId))
	var setName bytes.Buffer
	setName.WriteString(userIdString)
	setName.WriteString("Likes")
	if isFavor { //一个用户可以喜欢多个视频, Set名：userIdLikes
		err := rdb.SAdd(c, setName.String(), videoId)
		if err.Err() != nil {
			return err.Err()
		}
	} else {
		err := rdb.SRem(c, setName.String(), videoId)
		if err.Err() != nil {
			return err.Err()
		}
	}
	return nil
}

// 从缓存中取出关注状态，若userId 点赞了 videoId，则返回true，否则返回false
func (p *RedisProxy) GetFavoriteStatus(userId int64, videoId int64) (bool, error) {
	if userId <= 0 || videoId <= 0 {
		return false, errors.New("Redis GetFavoriteStatus Error: Id <= 0")
	}

	userIdString := strconv.Itoa(int(userId))
	var setName bytes.Buffer
	setName.WriteString(userIdString)
	setName.WriteString("Likes")
	res := rdb.SIsMember(c, setName.String(), videoId)

	return res.Val(), nil
}

// 设置缓存中的关注状态 userId为操作者Id， followUserId为被关注者Id， isFollow为操作类型，true则为关注，false为取消关注
func (p *RedisProxy) SetFollow(userId int64, followUserId int64, isFollow bool) error {
	if userId <= 0 || followUserId <= 0 {
		return errors.New("Redis SetFollow Error: Id <= 0")
	}
	if userId == followUserId {
		return errors.New("redis SetFollow error: you can't follow yourself")
	}
	userIdString := strconv.Itoa(int(userId))
	var setName bytes.Buffer
	setName.WriteString(userIdString)
	setName.WriteString("Follows")
	if isFollow { //一个用户可以关注多个用户, Set名：userIdFollows
		err := rdb.SAdd(c, setName.String(), followUserId)
		if err.Err() != nil {
			return err.Err()
		}
	} else {
		err := rdb.SRem(c, setName.String(), followUserId)
		if err.Err() != nil {
			return err.Err()
		}
	}
	return nil
}

// 从缓存中取出关注状态，若userId 关注了 followUserId，则返回true，否则返回false
func (p *RedisProxy) GetFollowStatus(userId int64, followUserId int64) (bool, error) {
	if userId <= 0 || followUserId <= 0 {
		return false, errors.New("Redis GetFavoriteStatus Error: Id <= 0")
	}

	userIdString := strconv.Itoa(int(userId))
	var setName bytes.Buffer
	setName.WriteString(userIdString)
	setName.WriteString("Follows")
	res := rdb.SIsMember(c, setName.String(), followUserId)

	return res.Val(), nil
}
