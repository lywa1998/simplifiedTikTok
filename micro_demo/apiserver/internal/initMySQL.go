package main

import (
	"github.com/micro/simplifiedTikTok/apiserver/internal/initMySQL"
)

func main() {
	initMySQL.CreateDB()
	initMySQL.CreateUserTable()
	initMySQL.CreateVideoTable()
	initMySQL.CreateCommentTable()
	initMySQL.CreateFavoriteTable()
}
