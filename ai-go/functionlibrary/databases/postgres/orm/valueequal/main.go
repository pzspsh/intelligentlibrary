/*
@File   : main.go
@Author : pan
@Time   : 2024-03-20 16:09:31
*/
package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Balance int
	Target  int
}

func main() {
	// 连接到数据库
	dsn := "host=localhost user=your_user dbname=your_db password=your_password sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 假设我们要找到Balance字段，将其加10，然后检查新值是否等于Target字段的值
	var user User
	db.Model(&User{}).
		Where("id = ?", 1).                              // 假设我们更新ID为1的用户
		Update("Balance", gorm.Expr("balance + ?", 10)). // // 更新Balance字段的值，对balance字段加10
		First(&user)                                     // 重新查询用户以获取更新后的Balance值

	// 现在，user.Balance包含了更新后的值
	if user.Balance == user.Target {
		fmt.Printf("Updated user with ID: %d, new Balance: %d is equal to Target: %d\n", user.ID, user.Balance, user.Target)
	} else {
		fmt.Printf("Updated user with ID: %d, new Balance: %d is not equal to Target: %d\n", user.ID, user.Balance, user.Target)
	}
}
