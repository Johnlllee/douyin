package handler

import "douyin/model"

//定义一下所有的respone

type CommonResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type UserInfoResponse struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserRegisterResponse struct {
	CommonResponse
	UserInfoResponse
}

type UserLoginResponse struct {
	CommonResponse
	UserInfoResponse
}

type UserInfoQueryResponse struct {
	CommonResponse
	User *model.UserInfo `json:"user"`
}
