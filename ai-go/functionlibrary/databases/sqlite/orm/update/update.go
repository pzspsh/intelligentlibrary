/*
@File   : update.go
@Author : pan
@Time   : 2023-06-12 16:11:10
*/
package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SqliteConn() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("../demo.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	} else {
		return db, nil
	}
}

type Student struct {
	Id     int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name   string `gorm:"column:name;type:varchar(255);default:null"`
	Number string `gorm:"column:number;type:varchar(255);default:null"`
	Title  string `gorm:"column:title;type:varchar(255);default:null"`
}

func Update(db *gorm.DB) error {
	// 更新操作： 更新单个字段
	var student Student
	err := db.Model(&student).Update("number", "10002")
	// 更新操作： 更新多个字段
	// err := db.Model(&student).Updates(Student{Number: "10003", Title: "hello pan"}) // non-zero fields
	// err := db.Model(&student).Updates(map[string]interface{}{"Number": "10004", "Title": "hello end"})
	if err != nil {
		return err.Error
	} else {
		return nil
	}
}

// User 模型
type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Age  int
}

func BatchUpdateMain() {
	// 连接数据库
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// 自动迁移模式
	db.AutoMigrate(&User{})

	// 假设这是我们要更新的用户数据
	updates := []struct {
		ID  uint
		Age int
	}{
		{ID: 1, Age: 30},
		{ID: 2, Age: 25},
		// ... 添加更多用户数据，‌直到1000条
	}

	// 批量更新的批次大小
	batchSize := 100

	// 分批更新用户数据
	for i := 0; i < len(updates); i += batchSize {
		end := i + batchSize
		if end > len(updates) {
			end = len(updates)
		}

		// 提取当前批次的用户ID
		var ids []uint
		for _, update := range updates[i:end] {
			ids = append(ids, update.ID)
		}

		// 批量更新操作
		db.Model(&User{}).Where("id IN ?", ids).Updates(map[string]interface{}{"age": 0})

		// 注意：‌由于Gorm的限制，‌我们不能直接在Updates中使用切片进行批量更新每个用户的不同年龄。‌
		// 因此，‌我们需要对每个用户单独进行更新操作。‌
		for _, update := range updates[i:end] {
			db.Model(&User{}).Where("id = ?", update.ID).Updates(User{Age: update.Age})
		}
	}

	fmt.Println("批量更新完成！")
}

func UpdateMain() {
	db, err := SqliteConn()
	if err != nil {
		fmt.Printf("sqlite conn err:%v", err)
	} else {
		fmt.Printf("sqlite conn successful:%v", db)
	}
	err = Update(db)
	if err != nil {
		fmt.Printf("update data err:%v", err)
	} else {
		fmt.Printf("update data successful.")
	}
}
