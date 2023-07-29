package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
 
var(
	username string = "root"
	password string = "hdbdn"
	host     string = "localhost"
	port     int = 3306
	Dbname   string = "simplifiedTikTok"

	db *gorm.DB
)

func init() {
	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, Dbname)
	con, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("无法连接数据库")
	}

	sqlDB, _ := con.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)

	db = con
}

func GetDB() *gorm.DB {
	return db
}