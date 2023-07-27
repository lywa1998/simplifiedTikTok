package initMySQL
// package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var(
	username string = "root"
	password string = "hdbdn"
	host     string = "localhost"
	port     int = 3306
	Dbname   string = "simplifiedTikTok"
)

// func main() {
func CreateDB() {
	// 连接数据库
	// dsn := "root:hdbdn@tcp(127.0.0.1:3306)/" //root:hdbdn 为mysql 用户名：密码
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", username, password, host, port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 创建新数据库
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", Dbname))
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("成功创建数据库 %s\n", Dbname)
}