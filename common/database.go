package common

import (
	"TikTok/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _db *gorm.DB

func init() {
	// MySQL 配置信息
	username := "root"   // 账号
	password := "aaaaaa" // 密码
	host := "127.0.0.1"  // 地址
	port := 3306         // 端口
	DBname := "tiktok"   // 数据库名称
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, DBname)
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect mysql.")
	}
	// 迁移数据表
	db.AutoMigrate(&model.User{}, &model.Comment{}, model.Message{}, &model.Video{})
	_db = db
}

func GetDB() *gorm.DB {
	return _db
}
