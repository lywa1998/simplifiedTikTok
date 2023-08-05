<<<<<<< HEAD:hdbdn_project/internal/initMySQL/createUserTable.go
package initMySQL
=======
package init_mysql
>>>>>>> origin/hdbdn:micro_demo/apiserver/internal/init_mysql/create_user_table.go

// package main

import (
	"database/sql"
	"fmt"
<<<<<<< HEAD:hdbdn_project/internal/initMySQL/createUserTable.go
	"io/ioutil"
	"path/filepath"
	"strings"

=======
>>>>>>> origin/hdbdn:micro_demo/apiserver/internal/init_mysql/create_user_table.go
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
)

var (
<<<<<<< HEAD:hdbdn_project/internal/initMySQL/createUserTable.go
	userTable string = "initMySQL/UserTable.sql"
=======
	userTable string = "internal/init_mysql/UserTable.sql"
>>>>>>> origin/hdbdn:micro_demo/apiserver/internal/init_mysql/create_user_table.go
)

func CreateUserTable() {
	// 连接到MySQL数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 读取SQL文件内容
<<<<<<< HEAD:hdbdn_project/internal/initMySQL/createUserTable.go
	abs, _ := filepath.Abs(".")
	abs = filepath.Join(abs, userTable)
	content, err := ioutil.ReadFile(abs)
=======
	content, err := os.ReadFile(userTable)
>>>>>>> origin/hdbdn:micro_demo/apiserver/internal/init_mysql/create_user_table.go
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
