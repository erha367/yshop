package utils

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
	"yshop/middleware"
)

var SQLite *gorm.DB

func InitSQLite(path string) *gorm.DB {
	// 使用 SQLite 驱动打开数据库连接
	dbConn, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		Logger:         middleware.NewSQLLog(),
	})
	if err != nil {
		panic("sqlite init error: " + err.Error())
	}
	sqlDB, err := dbConn.DB()
	if err != nil {
		panic("sqlite conn error: " + err.Error())
	}
	log.Println(`sqlite database connect success`, path)
	// SetMaxIdleCons 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(30)
	// SetMaxOpenCons 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(50)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return dbConn
}
