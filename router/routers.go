package router

import (
	"douyin/handler/user"
	"douyin/handler/video"
	"douyin/middleware"
	"github.com/gin-gonic/gin"
)
import "net/http"

type rsp struct {
	StatusCode uint
	StatusMsg  string
}

func InitAllRouters(ge *gin.Engine) {

	ge.Static("static", "./static")

	baseGroup := ge.Group("/douyin")

	//ping接口只为了测试
	baseGroup.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, rsp{
			StatusCode: 0,
			StatusMsg:  "hello",
		})
	})

	//TODO 加上其他的router

	//中间件还没上
	baseGroup.POST("/user/register/", user.UserRegisterHandler)
	//中间件还没上
	baseGroup.POST("/user/login/", user.UserLoginHandler)
	//已经加上jwt中间件
	baseGroup.GET("/user/", middleware.JWTMiddleware(), user.UserInfoHandler)

	//video
	baseGroup.GET("/feed/", video.FeedVideoHandler)

	baseGroup.POST("/publish/action/", middleware.JWTMiddleware(), video.PostVideoHandler)

	baseGroup.GET("/publish/list/", middleware.JWTMiddleware(), video.PublishedVideoListHandler)

}
