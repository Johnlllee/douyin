package action

import (
	"douyin/handler"
	"douyin/service/ActionSvc"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FollowListHandler(context *gin.Context) {
	userIdToQueryRaw := context.Query("user_id")
	if userIdToQueryRaw == "" {
		userIdToQueryRaw = context.PostForm("user_id")
	}

	if userIdToQueryRaw == "" {
		context.JSON(http.StatusOK, &handler.CommonResponse{
			StatusCode: 1,
			StatusMsg:  "[FollowListHanlder]: Userid To Query Not Exist.",
		})
		return
	}

	userIdToQuery, err := strconv.ParseInt(userIdToQueryRaw, 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, &handler.CommonResponse{
			StatusCode: 1,
			StatusMsg:  "[FollowListHandler]: Userid To Query Parse Error.",
		})
		return
	}

	list, err := ActionSvc.GetFollowList(userIdToQuery)
	if err != nil {
		context.JSON(http.StatusOK, &handler.CommonResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, &handler.FollowListResponse{
		CommonResponse: handler.CommonResponse{
			StatusCode: 0,
		},
		User_list: list,
	})
	return
}
