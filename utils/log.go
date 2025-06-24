package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logx"
)

func init() {
	// 配置go-zero日志（可选）
	logx.MustSetup(logx.LogConf{
		Mode:     "file",   // 输出到控制台
		Encoding: "json",   // 格式化为JSON
		Path:     "./logs", // 日志文件路径
		Compress: true,     // 压缩旧日志
		Rotation: "daily",
		KeepDays: 7,
	})
}

// 封装日志记录函数
func Info(c *gin.Context, msg, key string, v ...interface{}) {
	// 从上下文中获取 trace_id
	traceID, ok := c.Get("trace_id")
	if !ok {
		traceID = "unknown"
	}
	// 记录日志
	logx.WithContext(c.Request.Context()).Infow(msg, logx.Field(`trace_id`, traceID), logx.Field(key, v))
}

func Error(c *gin.Context, msg, key string, v ...interface{}) {
	// 从上下文中获取 trace_id
	traceID, ok := c.Get("trace_id")
	if !ok {
		traceID = "unknown"
	}
	// 记录日志
	logx.WithContext(c.Request.Context()).Errorw(msg, logx.Field(`trace_id`, traceID), logx.Field(key, v))
}
