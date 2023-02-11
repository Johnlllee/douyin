package main

import (
	"douyin/cache"
	"douyin/model"
	"douyin/router"
	"douyin/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//初始化db
	model.InitDB()

	cache.InitCache()

	service.InitAllServiceMgr()

	//初始化router
	router.InitAllRouters(r)

	r.Run()
}
