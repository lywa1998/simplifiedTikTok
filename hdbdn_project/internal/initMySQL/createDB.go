package initMySQL
// package main

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

var(
	mysqlConfig = initMySQLConfig("configs/mysql.yaml")
)

type MySQLConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	DBName string `yaml:"dbname"`
}

func initMySQLConfig(filePath string) *MySQLConfig {
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取mysql配置文件失败")
		return &MySQLConfig{}
	}
  
	var cfg MySQLConfig
	if err := yaml.Unmarshal(content, &cfg); err != nil {
		fmt.Println("解析mysql配置文件失败")
		return &MySQLConfig{}
	}
	return &cfg
  }

func CreateDB() {
	if mysqlConfig.DBName == "" {
		panic("mysql config error")
	}

	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 创建新数据库
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", mysqlConfig.DBName))
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("成功创建数据库 %s\n", mysqlConfig.DBName)
}