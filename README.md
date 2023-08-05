##micro_demo已经完成基础接口并使用app证实可用（windows10环境）

##第一步
在每个项目文件下执行go mod tidy
例如~\micro_demo\apiserver go mod tidy

##第二步初始化数据库(记得检查项目文件下config文件里mysql、redis的配置是否与本地相同)
命令行或shell：~\micro_demo\apiserve go run .\internal\init_db.go

##第三步运行依赖项目
首先启动除apiserver外的其他项目(go run 项目路径下cmd中的文件)
例如：命令行或shell：~\micro_demo\apiserve go run .\cmd\user_service_server.go

##第四步运行apiserver
命令行或shell：~\micro_demo\apiserver go run .\cmd\apigateway.go