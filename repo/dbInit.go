package repo

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var Db *gorm.DB

func DbInit() {
	user := "root"
	pass := "3.1415926mkrgxy"
	host := "127.0.0.1"
	port := 3306
	dbname := "gorm_test"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 官网的配置
		logger.Config{
			SlowThreshold: time.Second,  // 慢 SQL 阈值
			LogLevel:      logger.Error, // Log level
			Colorful:      false,        // 彩色打印
		},
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbname)
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	err = Db.AutoMigrate(&User{}, &Video{}, &Comment{}, &Login{})
	if err != nil {
		log.Panicln("err:", err.Error())
	}
}
