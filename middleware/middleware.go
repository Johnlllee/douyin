package middleware

import (
	"douyin/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWTMiddleware() gin.HandlerFunc {

	return func(context *gin.Context) {

		tokenStr := context.Query("token")

		if tokenStr == "" {
			tokenStr = context.PostForm("token")
		}

		if tokenStr == "" {
			context.JSON(http.StatusOK, &handler.CommonResponse{
				StatusCode: 1,
				StatusMsg:  "User Not Exist.",
			})
			context.Abort()
			return
		}

		token, ok := ParseToken(tokenStr)
		if ok == false {
			context.JSON(http.StatusOK, &handler.CommonResponse{
				StatusCode: 1,
				StatusMsg:  "Token Not Correct.",
			})
			context.Abort()
			return
		}

		if time.Now().Unix() > token.ExpiresAt {
			context.JSON(http.StatusOK, &handler.CommonResponse{
				StatusCode: 1,
				StatusMsg:  "Token Expired.",
			})
			context.Abort()
			return
		}

		context.Set("userid", token.UserId)
		context.Next()
	}

}

func SHA1Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		password := c.Query("password")
		if password == "" {
			password = c.PostForm("password")
		}
		c.Set("password", SHA1(password))
		c.Next()
	}
}
