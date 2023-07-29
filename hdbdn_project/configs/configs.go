package configs

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v2"
)

var(
	MySQL = initMySQLConfig("configs/mysql.yaml")
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
		fmt.Println("读取mysql配置文件失败: ", err)
		return &MySQLConfig{}
	}
  
	var cfg MySQLConfig
	if err := yaml.Unmarshal(content, &cfg); err != nil {
		fmt.Println("解析mysql配置文件失败")
		return &MySQLConfig{}
	}
	return &cfg
}