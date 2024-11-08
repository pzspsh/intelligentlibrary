/*
@File   : main.go
@Author : pan
@Time   : 2024-11-08 14:00:38
*/
package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	Data string
}

func main() {
	// 连接数据库
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// 确保表已创建
	db.AutoMigrate(&Model{})

	// 假设有40万条数据
	var models []Model
	// ... 这里填充models数据 ...

	// 批量更新数据
	db.Session(&gorm.Session{FullSaveAssociations: true}).Find(&models).Updates(Model{Data: "new_value"})
}
