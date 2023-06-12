/*
@File   : update.go
@Author : pan
@Time   : 2023-06-12 16:11:10
*/
package main

import (
	"fmt"

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

func main() {
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
