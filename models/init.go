package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var Db *gorm.DB
var err error

func init() {
	fmt.Println("链接数据库 \n\r")
	dsn := "root:root@tcp(127.0.0.1:3306)/gin_gorm_obj?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	defer func() {
		err2 := recover() //	recover 内置函数捕获异常
		if err2 != nil {  // 	nil 是 err 的零值
			fmt.Println("err = ", err)
			//runtime error: index out of range [3] with length 3
		}
	}()
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB, err1 := Db.DB()
	if err1 != nil {
		panic("failed to connect database")
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}
