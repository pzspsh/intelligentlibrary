/*
@File   : create.go
@Author : pan
@Time   : 2023-06-12 16:09:41
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
	gorm.Model
	Id     int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name   string `gorm:"column:name;type:varchar(255);default:null"`
	Number string `gorm:"column:number;type:varchar(255);default:null"`
	Title  string `gorm:"column:title;type:varchar(255);default:null"`
}

// 设置表名称
func (Student) TableName() string {
	return "student"
}

func Create(db *gorm.DB) {
	dbMig := db.Migrator()
	if !dbMig.HasTable(&Student{}) { // 查看表是否存在，不存在这则执行创建表操作
		db.AutoMigrate(&Student{})
		fmt.Println("创建表 student成功。。。")
	}
}

func main() {
	db, err := SqliteConn()
	if err != nil {
		fmt.Printf("sqlite conn err:%v", err)
	} else {
		fmt.Printf("sqlite conn successful:%v", db)
	}
	Create(db)
}
