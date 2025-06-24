package middleware

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm/logger"
	"time"
)

type SQLLog struct {
	LogLevel logger.LogLevel
}

func NewSQLLog() *SQLLog {
	return new(SQLLog)
}

func (l *SQLLog) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *SQLLog) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Info {
		return
	}
	// 从上下文中获取 trace_id
	traceID, ok := ctx.Value("trace_id").(string)
	if !ok {
		traceID = "unknown"
	}
	logx.WithContext(ctx).Debugf("TraceID: %s, %s", traceID, fmt.Sprintf(msg, data...))
}

func (l *SQLLog) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Warn {
		return
	}
	traceID, ok := ctx.Value("trace_id").(string)
	if !ok {
		traceID = "unknown"
	}
	logx.WithContext(ctx).Infof("TraceID: %s, %s", traceID, fmt.Sprintf(msg, data...))
}

func (l *SQLLog) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Error {
		return
	}
	traceID, ok := ctx.Value("trace_id").(string)
	if !ok {
		traceID = "unknown"
	}
	logx.WithContext(ctx).Errorf("TraceID: %s, %s", traceID, fmt.Sprintf(msg, data...))
}

func (l *SQLLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	//这块的逻辑可以自己根据业务情况修改
	fmt.Println(l.LogLevel)
	elapsed := time.Since(begin)
	sql, rows := fc()
	traceID, ok := ctx.Value("trace_id").(string)
	if !ok {
		traceID = "unknown"
	}
	logx.WithContext(ctx).WithDuration(elapsed).Slowf("TraceID: %s, Trace sql: %v  row： %v  err: %v", traceID, sql, rows, err)
}
