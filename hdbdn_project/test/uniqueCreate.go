package test

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"type:varchar(100);unique_index"`
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	// 自定义Email唯一性校验
	var count int64
	tx.Model(u).Where("email = ?", u.Email).Count(&count)
	if count > 0 {
		return errors.New("email already existed")
	}

	return nil
}

func CreateUser(user *User) error {
	db, err := gorm.Open(mysql.Open("root:hdbdn@tcp(127.0.0.1:3306)/test"), &gorm.Config{})
	if err != nil {
		panic("无法连接数据库")
	}
	return db.Create(user).Error
}

// type User struct {
// 	Id       uint   `gorm:"primaryKey"`
// 	Username string `gorm:"unique"`
// 	Email    string
// 	Age      int
// }

// func main() {
// 	// 建立数据库连接
// 	dsn := "root:hdbdn@tcp(127.0.0.1:3306)/test"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("无法连接数据库")
// 	}

// 	// 插入数据
// 	user := User{
//  	Username: "john_doe",
// 		Email:    "john@example.com",
// 		Age:      30,
// 	}

// 	result := db.Create(&user)
// 	if result.Error != nil {
// 		if mysqlErr, ok := result.Error.(*mysql.MySQLError); ok {
// 			// 检查MySQL错误码
// 			if mysqlErr.Number == 1062 {
// 				// 唯一性约束冲突
// 				err = errors.New("用户名已经存在")
// 			} else {
// 				// 其他MySQL错误
// 				err = errors.New("插入数据时出现MySQL错误")
// 			}
// 			// 处理错误
// 		} else {
// 			// 非MySQL错误
// 			err = errors.New("插入数据时出现其他错误")
// 		}
// 	}

// 	// 其他处理
// }
