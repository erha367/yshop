package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"yshop/router"
	"yshop/utils"
)

// 定义命令行参数，用于接收配置文件路径，默认值为空字符串
var configPath = flag.String("f", "", "配置文件的路径，例如 /Users/yangsen13/waibao/yshop/config/test.yaml")

func main() {
	// 解析命令行参数
	flag.Parse()
	// 加载配置
	utils.InitConfig(*configPath)
	gin.SetMode(gin.DebugMode)
	utils.SQLite = utils.InitSQLite(utils.AppConfig.Sqlite)
	r := router.InitRouter()
	port := fmt.Sprintf(":%d", utils.AppConfig.App.Port)
	r.Run(port)
}
