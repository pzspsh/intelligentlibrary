/*
@File   : main.go
@Author : pan
@Time   : 2024-10-11 16:23:02
*/
package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
	Age  int
}

func main() {
	dsn := "host=localhost user=youruser dbname=yourdb sslmode=disable password=yourpassword"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移
	db.AutoMigrate(&User{})

	// 插入或更新数据
	user := User{Name: "John", Age: 30}
	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoUpdates: clause.AssignmentColumns([]string{"age"}),
	}).Create(&user)

	// 检查错误
	if result.Error != nil {
		fmt.Println("Error:", result.Error)
	} else {
		fmt.Println("Inserted/Updated User:", user)
	}
}

/*
Clauses方法结合ON CONFLICT子句来插入或更新数据。
如果插入的数据的Name字段与数据库中已存在的记录冲突，则会更新该记录的Age字段。
*/

/*

// 对于SQLite，使用REPLACE语句
db.Exec("REPLACE INTO users (name, age) VALUES (?, ?)", "John", 30)

// 或者使用INSERT OR REPLACE语句
db.Exec("INSERT OR REPLACE INTO users (name, age) VALUES (?, ?)", "John", 30)
*/
