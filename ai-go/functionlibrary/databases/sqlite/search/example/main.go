/*
@File   : main.go
@Author : pan
@Time   : 2024-06-03 13:13:55
*/
package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MyModel struct {
	gorm.Model
	FieldName string // 假设这是你想要查询的字段名
	// ... 其他字段
}

//  golang gorm查询近三天前到现在的某个字段值

func main() {
	// 连接到数据库（这里使用SQLite作为例子）
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移schema
	db.AutoMigrate(&MyModel{})

	// 获取当前时间
	now := time.Now()

	// 计算三天前的时间
	threeDaysAgo := now.AddDate(0, 0, -3)

	// 使用GORM查询近三天前到现在的记录
	var results []MyModel
	err = db.Where("created_at BETWEEN ? AND ?", threeDaysAgo, now).Find(&results).Error
	if err != nil {
		panic("failed to query database")
	}

	// 输出查询结果中你感兴趣的字段值
	for _, result := range results {
		fmt.Println(result.FieldName)
	}
}
