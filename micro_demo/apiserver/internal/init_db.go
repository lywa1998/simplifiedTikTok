package main

import (
	"github.com/micro/simplifiedTikTok/apiserver/internal/init_mysql"
)

func main() {
	init_mysql.CreateDB()
	init_mysql.CreateUserTable()
	init_mysql.CreateVideoTable()
	init_mysql.CreateCommentTable()
	init_mysql.CreateFavoriteTable()
}
