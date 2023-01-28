package user

import (
	"douyin/config"
	"douyin/handler"
	"douyin/middleware"
	"douyin/service"
	"douyin/service/userSvc"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegisterHandler(c *gin.Context) {
	userName := c.Query("username")
	//不走sha1就这样写一下就行
	password := c.Query("password")

	uir, err := Register(userName, password)
	if err != nil {
		c.JSON(http.StatusOK, &handler.UserRegisterResponse{
			CommonResponse: handler.CommonResponse{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, &handler.UserRegisterResponse{
		CommonResponse: handler.CommonResponse{
			StatusCode: 0,
			StatusMsg:  "",
		},
		UserInfoResponse: handler.UserInfoResponse{
			UserId: uir.UserId,
			Token:  uir.Token,
		},
	})
	return

}

func Register(userName string, password string) (*handler.UserInfoResponse, error) {

	err := checkValid(userName, password)

	if err != nil {
		return nil, err
	}

	userid := service.UserServiceManager.RequestUserId()

	err = userSvc.InsertUserData(userName, password, userid)
	if err != nil {
		return nil, err
	}

	token, err := middleware.GenerateToken(userid)

	if err != nil {
		return nil, err
	}

	ret := &handler.UserInfoResponse{
		UserId: userid,
		Token:  token,
	}
	return ret, nil

}

// api
func checkValid(userName string, password string) error {
	if userName == "" {
		return errors.New("[UserRegisterHandler]: Empty Username.")
	}

	if len(userName) > config.MAX_USERNAME_LENGTH {
		return errors.New("[UserRegisterHandler]: Too Long Username. Username: " + userName)
	}

	if password == "" {
		return errors.New("[UserRegisterHandler]: Empty Password. Username: " + userName)
	}
	return nil
}
