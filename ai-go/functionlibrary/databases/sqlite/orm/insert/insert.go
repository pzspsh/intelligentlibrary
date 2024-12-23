/*
@File   : insert.go
@Author : pan
@Time   : 2023-06-12 16:10:22
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

func (s *Student) Insert(db *gorm.DB) error {
	err := db.Create(s)
	if err.Error != nil {
		fmt.Printf("insert data err:%v", err.Error)
		return err.Error
	} else {
		fmt.Printf("insert data successful")
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
	s := &Student{Name: "pan", Number: "1001", Title: "hello world"}
	err = s.Insert(db)
	if err != nil {
		fmt.Printf("insert data fault:%v", err)
	} else {
		fmt.Printf("insert data successful")
	}
}
