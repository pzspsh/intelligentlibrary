/*
@File   : main.go
@Author : pan
@Time   : 2024-06-03 16:29:44
*/
package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MyModel struct {
	gorm.Model
	FieldName string // 假设你要查询的字段名是FieldName
	Status    int
	// 其他字段...
}

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移schema（如果还没有迁移的话）
	db.AutoMigrate(&MyModel{})

	// 计算近三天前的时间
	threeDaysAgo := time.Now().AddDate(0, 0, -3)

	// 查询近三天前到现在的记录，并且status不等于0且不等于1
	var records []MyModel
	result := db.Where("created_at BETWEEN ? AND ? AND status NOT IN (?)", threeDaysAgo, time.Now(), []int{0, 1}).Find(&records)
	if result.Error != nil {
		panic(result.Error)
	}

	// 输出查询结果
	for _, record := range records {
		fmt.Printf("ID: %d, FieldName: %s, Status: %d, CreatedAt: %s\n", record.ID, record.FieldName, record.Status, record.CreatedAt)
	}
}
