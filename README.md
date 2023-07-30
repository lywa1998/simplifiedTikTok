# 注意
每个人创建自己的文件夹xxx_project，在自己的文件夹目录下开发项目
# 开始 
在自己的文件夹目录下simplifiedTikTok/xxx_project执行：go mod init github.com/xxx/simplifiedTikTok
然后在simplifiedTikTok/目录下执行：go work use xxx_project
# 接下来回到自己的项目路径下开发即可

# /internal/initMySQL.go
这是数据库初始化文件，在hdbdn_project文件夹下使用 go run .\internal\initMySQL.go 即可创建

# configs文件夹用于保存配置，可按需更改配置
其中configs.go文件用于初始化配置

# pkg/dao
其中的文件初始化各种数据库连接的

# pkg/model
定义不同mysql表的gorm模型及对模型的操作

# pkg/utils
jwt_token.go用于token生成与验证