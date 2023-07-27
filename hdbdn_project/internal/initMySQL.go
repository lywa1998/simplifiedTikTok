package main

import (
	"github.com/hdbdn77/simplifiedTikTok/internal/initMySQL"
)

func main() {
	initMySQL.CreateDB()
	initMySQL.CreateUserTable()
	initMySQL.CreateVideoTable()
	initMySQL.CreateCommentTable()
	initMySQL.CreateFavoriteTable()
}