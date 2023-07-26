package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func mainA() {
    // 连接数据库
    dsn := "root:hdbdn@tcp(127.0.0.1:3306)/"  //root:hdbdn 为mysql 用户名：密码
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // 创建新数据库
    databaseName := "new_database"
    _, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", databaseName))
    if err != nil {
        panic(err.Error())
    }
    fmt.Printf("成功创建数据库 %s\n", databaseName)
}