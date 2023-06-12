/*
@File   : search.go
@Author : pan
@Time   : 2023-06-12 16:10:46
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

func main() {
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
