/*
@File   : main.go
@Author : pan
@Time   : 2024-11-06 15:48:40
*/
package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
	Age  uint
}

func BatchUpdate1() {
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sql := "UPDATE users SET name = CASE id "
	sql += "WHEN 1 THEN 'NewName1' "
	sql += "WHEN 2 THEN 'NewName2' "
	sql += "WHEN 3 THEN 'NewName3' "
	sql += "ELSE name END WHERE id IN (1, 2, 3)"

	db.Exec(sql)

	fmt.Println("Batch update with SQL expression completed.")
}

func BatchUpdate2() {
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 假设我们要把所有年龄大于18的用户的名字改为"Adult"
	db.Model(&User{}).Where("age > ?", 18).Update("name", "Adult")

	// 或者使用Updates方法来更新多个字段
	db.Model(&User{}).Where("age > ?", 18).Updates(User{Name: "Adult", Age: 0}) // 注意：这会把年龄设为0，可能不是你想要的
	// 更安全的做法是指定要更新的字段
	db.Model(&User{}).Where("age > ?", 18).Updates(map[string]interface{}{"name": "Adult"})

	fmt.Println("Batch update completed.")
}

func BatchUpdate() {
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	tx := db.Begin()
	if tx.Error != nil {
		panic("failed to start transaction")
	}

	// 假设我们有一个用户ID和对应新名字的映射
	updates := map[uint]string{
		1: "NewName1",
		2: "NewName2",
		3: "NewName3",
	}

	for id, name := range updates {
		db.Model(&User{}).Where("id = ?", id).Update("name", name)
	}

	tx.Commit()
}

func main() {

}
