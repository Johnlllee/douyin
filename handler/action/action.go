package action

import (
	"douyin/handler"
	"douyin/service/ActionSvc"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ActionHandler(context *gin.Context) {

	useridraw, exists := context.Get("userid")
	userid := useridraw.(int64)

	if exists == false {
		context.JSON(http.StatusOK, &handler.CommonResponse{
			StatusCode: 1,
			StatusMsg:  "[ActionHandler]Parse UserId Error. UserId Not Exist.",
		})
		return
	}

	toUserIdStr := context.Query("to_user_id")

	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusOK, &handler.CommonResponse{
			StatusCode: 1,
			StatusMsg:  "[ActionHandler]Parse ToUserId Error." + err.Error(),
		})
		return
	}

	actionStr := context.Query("action_type")
	action, err := strconv.Atoi(actionStr)

	if (action != 1 && action != 2) || err != nil {
		context.JSON(http.StatusOK, &handler.CommonResponse{
			StatusCode: 1,
			StatusMsg:  "[ActionHandler]Parse Action Type Error.",
		})
		return
	}

	err = ActionSvc.FollowOrUnfollowUser(userid, toUserId, action)
	if err != nil {
		context.JSON(http.StatusOK, &handler.CommonResponse{
			StatusCode: 1,
			StatusMsg:  "[ActionHandler]error: " + err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, &handler.CommonResponse{
		StatusCode: 0,
		StatusMsg:  "",
	})

}
