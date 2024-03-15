/*
@File   : main.go
@Author : pan
@Time   : 2024-03-15 17:34:31
*/
package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 假设你有这样一个模型
type User struct {
	gorm.Model
	Balance int // 用户余额
	Score   int // 用户积分
}

func main() {
	// 连接到数据库 (需要替换为你的数据库连接信息)
	dsn := "user=your_user dbname=your_db sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移模式
	db.AutoMigrate(&User{})

	// 假设我们要将某个用户的余额增加10，积分减少5
	// 这里我们假设用户的ID是1
	userId := 1
	balanceIncrement := 10
	scoreDecrement := 5

	// 使用GORM的Updates方法，结合SQL的EXCLUDED关键字
	// 注意：不同的数据库可能对更新前的值有不同的关键字，例如MySQL是`OLD`
	db.Model(&User{}).Where("id = ?", userId).Updates(map[string]interface{}{
		"balance": gorm.Expr("balance + ?", balanceIncrement),
		"score":   gorm.Expr("score - ?", scoreDecrement),
	})

	// 验证是否更新成功
	var user User
	if db.First(&user, userId).Error == nil {
		fmt.Printf("Updated user: %+v\n", user)
	} else {
		fmt.Println("Failed to update user")
	}
}
