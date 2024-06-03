/*
@File   : main.go
@Author : pan
@Time   : 2024-06-03 16:38:57
*/
package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type YourModel struct {
	ID        uint      `gorm:"primaryKey"`
	FieldName time.Time `gorm:"type:datetime"` // 根据您的实际字段名和类型进行修改
	Status    int
}

func main() {
	dsn := "your_database_connection_string" // 请替换为您的数据库连接字符串
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 获取三天前的日期
	threeDaysAgo := time.Now().AddDate(0, 0, -3)

	var results []YourModel
	err = db.Where("FieldName >= ? AND (Status != ? OR Status != ?)", threeDaysAgo, 0, 1).Find(&results).Error
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
