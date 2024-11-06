/*
@File   : main.go
@Author : pan
@Time   : 2024-11-06 15:04:03
*/
package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
}

func main() {
	// 初始化数据库连接（这里使用 SQLite 作为示例）
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移模式
	db.AutoMigrate(&User{})

	// 判断表是否有数据
	hasData := HasData(db, &User{})
	if hasData {
		fmt.Println("The table has data.")
	} else {
		fmt.Println("The table is empty.")
	}
}

// HasData 函数用于判断表是否有数据
func HasData(db *gorm.DB, model interface{}) bool {
	var count int64
	db.Model(model).Count(&count)
	return count > 0
}
