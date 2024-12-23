/*
@File   : search.go
@Author : pan
@Time   : 2023-06-12 16:10:46
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

func Search(db *gorm.DB) (bool, *Student) {
	var student Student
	err := db.First(&student, 1) // find student with integer primary key
	// err := db.First(&student, "name = ?", "pan") // find student with name pan
	if err.Error != nil {
		fmt.Printf("search data err:%v", err.Error)
		return false, nil
	} else {
		fmt.Println("search data successfule.")
		return true, &student
	}
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func Searchexist() { // 判断数据，存在则更新，不存在则创建
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// 自动迁移
	db.AutoMigrate(&Product{})

	// 创建
	db.Create(&Product{Code: "D42", Price: 100})

	// 更新或创建
	var product Product
	db.First(&product, "code = ?", "D42") // 尝试找到记录
	if product.Code == "" {
		// 如果不存在，则创建
		db.Create(&Product{Code: "D42", Price: 200})
	} else {
		// 如果存在，则更新
		db.Model(&product).Update("Price", 200)
	}

	// Save 方法示例
	db.Save(&Product{Code: "D42", Price: 300}) // 如果存在则更新，不存在则创建
}

type User struct {
	ID         uint
	Field1     string
	Field2     string
	OtherField string
}

func Search2() {
	// 初始化数据库连接
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移模式
	db.AutoMigrate(&User{})

	// 假设我们要查询 Field1 等于 "value1" 或者 Field2 等于 "value2" 的用户
	value1 := "value1"
	value2 := "value2"

	var users []User
	db.Where("field1 = ? OR field2 = ?", value1, value2).Find(&users)

	// 打印查询结果
	for _, user := range users {
		fmt.Printf("ID: %d, Field1: %s, Field2: %s, OtherField: %s\n", user.ID, user.Field1, user.Field2, user.OtherField)
	}
}

func SearchRun() {
	db, err := SqliteConn()
	if err != nil {
		fmt.Printf("sqlite conn err:%v", err)
	} else {
		fmt.Printf("sqlite conn successful:%v", db)
	}
	ok, student := Search(db)
	if ok {
		fmt.Printf("student search successful:%v", student)
	}
}

func main() {

}
