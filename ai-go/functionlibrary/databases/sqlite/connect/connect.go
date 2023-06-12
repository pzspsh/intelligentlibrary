/*
@File   : connect.go
@Author : pan
@Time   : 2023-06-12 16:08:34
*/
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func SqliteConn(dbpath string) (*sql.DB, error) {
	// 打开数据库
	db, err := sql.Open("sqlite3", dbpath) //打开数据库，如果不存在，则创建
	if err != nil {
		fmt.Printf("sqlite3 open db err:%v\n", err)
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Printf("sqlite3 conn err:%v\n", err)
		return nil, err
	}
	return db, err
}

func main() {
	db, err := SqliteConn("../demo.db")
	if err != nil {
		fmt.Printf("sqlite conn err:%v", err)
	} else {
		fmt.Printf("sqlite conn successful:%v", db)
	}
}
