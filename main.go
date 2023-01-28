package main

import (
	"douyin/model"
	"douyin/router"
	"douyin/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//初始化db
	model.InitDB()

	service.InitAllServiceMgr()

	//初始化router
	router.InitAllRouters(r)

	r.Run()
}
