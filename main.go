package main

import (
	"github.com/gin-gonic/gin"
	"yshop/router"
	"yshop/utils"
)

func main() {
	// 禁用Gin默认日志
	gin.SetMode(gin.DebugMode)
	utils.SQLite = utils.InitSQLite("/Users/yangsen13/waibao/galaxy/yshop.db")
	r := router.InitRouter()
	r.Run(":8088")
}
