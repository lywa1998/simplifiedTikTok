package configs

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

var (
	MySQL = initMySQLConfig("configs/mysql.yaml")
	Jwt   = initJwtConfig("configs/jwt.yaml")
)

type MySQLConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
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

type JwtConfig struct {
	PrivKey string `yaml:"privKey"`
}

func initJwtConfig(filePath string) *JwtConfig {
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取mysql配置文件失败: ", err)
		return &JwtConfig{}
	}

	var cfg JwtConfig
	if err := yaml.Unmarshal(content, &cfg); err != nil {
		fmt.Println("解析mysql配置文件失败")
		return &JwtConfig{}
	}
	return &cfg
}
