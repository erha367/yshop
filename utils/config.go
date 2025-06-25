package utils

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name  string `json:"name"`
		Port  int    `json:"port"`
		Debug bool   `json:"debug"`
	} `json:"app"`
	Mysql   string `json:"mysql"`
	Sqlite  string `json:"sqlite"`
	Logging struct {
		Level  string `json:"level"`
		Format string `json:"format"`
	} `json:"logging"`
}

var AppConfig Config

func InitConfig(path string) {
	// 初始化 Viper
	viper.SetConfigFile(path) // 直接使用传入的配置文件路径

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}

	// 解析配置到结构体
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("解析配置失败: %v", err)
	}
	AppConfig = config
}
