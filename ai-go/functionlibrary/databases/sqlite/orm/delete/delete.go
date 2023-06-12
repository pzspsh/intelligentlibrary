/*
@File   : delete.go
@Author : pan
@Time   : 2023-06-12 16:10:02
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

func Delete(db *gorm.DB) error {
	var student Student
	err := db.Delete(&student, 1)
	if err.Error != nil {
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
	err = Delete(db)
	if err != nil {
		fmt.Printf("delete data err:%v", err)
	} else {
		fmt.Printf("delete data successful.")
	}
}
