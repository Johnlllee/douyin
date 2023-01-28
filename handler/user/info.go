package user

import (
	"douyin/handler"
	"douyin/service/userSvc"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserInfoHandler(context *gin.Context) {

	rawid, ok := context.Get("userid")

	if ok == false {
		context.JSON(http.StatusOK, &handler.CommonResponse{
			StatusCode: 1,
			StatusMsg:  "Get User Id Error",
		})
		return
	}
	userid := rawid.(int64)

	userinfo, err := userSvc.QueryUserInfo(userid)
	if err != nil {
		context.JSON(http.StatusOK, &handler.CommonResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, &handler.UserInfoQueryResponse{
		CommonResponse: handler.CommonResponse{
			StatusCode: 0,
		},
		User: userinfo,
	})
}
