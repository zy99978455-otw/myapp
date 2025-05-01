package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// 连接数据库
var dsn = "root:MYSQL@tcp(localhost:3306)/web?charset=utf8mb4"
var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("初始化数据库失败", err)
	}
}
