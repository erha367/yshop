package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

// GoZeroLogger 是一个使用go-zero日志的Gin中间件
func GoZeroLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 生成 trace_id
		traceID := uuid.New().String()
		// 将 trace_id 存入上下文
		c.Set("trace_id", traceID)
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		// 使用go-zero的日志记录请求信息
		logx.WithContext(c.Request.Context()).Infow(`request info:`,
			logx.Field("trace_id", traceID),
			logx.Field("status_code", c.Writer.Status()),
			logx.Field("latency", end.Sub(start)),
			logx.Field("client_ip", c.ClientIP()),
			logx.Field("method", c.Request.Method),
			logx.Field("path", c.Request.URL.Path),
		)
	}
}
