package router

import (
	"github.com/gin-gonic/gin"
	"yshop/middleware"
	"yshop/server/controller"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	// 使用go-zero的日志中间件
	r.Use(middleware.GoZeroLogger())
	// 注册路由
	r.GET("/ping", controller.Ping)

	return r
}
