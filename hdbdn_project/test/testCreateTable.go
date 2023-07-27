package test

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func mainB() {
	// 连接到MySQL数据库
	dsn := "root:hdbdn@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 读取SQL文件内容
	sqlFile := "./init.sql"
	content, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		panic(err.Error())
	}

	// 将SQL文件内容拆分为单个SQL语句
	sqlStatements := strings.Split(string(content), ";")

	// 遍历并执行每个SQL语句
	for _, statement := range sqlStatements {
		// 忽略空语句
		if strings.TrimSpace(statement) == "" {
			continue
		}

		// 执行SQL语句
		_, err := db.Exec(statement)
		if err != nil {
			fmt.Printf("执行SQL语句时出错: %v\n", err)
			continue
		}
		fmt.Println("成功执行SQL语句: ", statement)
	}
}
