package user

import (
	"douyin/handler"
	"douyin/middleware"
	"douyin/service/userSvc"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLoginHandler(context *gin.Context) {

	username := context.Query("username")
	password := context.Query("password")

	infoRsp, err := Login(username, password)

	if err != nil {
		context.JSON(http.StatusOK, &handler.CommonResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, &handler.UserLoginResponse{
		CommonResponse: handler.CommonResponse{
			StatusCode: 0,
		},
		UserInfoResponse: handler.UserInfoResponse{
			UserId: infoRsp.UserId,
			Token:  infoRsp.Token,
		},
	})

}

func Login(username, password string) (*handler.UserInfoResponse, error) {

	err := checkValid(username, password)
	if err != nil {
		return nil, err
	}
	err = nil

	userid, err := userSvc.QueryUserLogin(username, password)
	if err != nil {
		return nil, err
	}

	token, err := middleware.GenerateToken(userid)

	if err != nil {
		return nil, err
	}

	return &handler.UserInfoResponse{
		UserId: userid,
		Token:  token,
	}, nil

}
