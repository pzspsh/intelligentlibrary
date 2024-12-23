/*
@File   : insert.go
@Author : pan
@Time   : 2023-06-12 16:10:22
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

func Insert(db *sql.DB) (sql.Result, error) {
	stmt, err := db.Prepare("INSERT INTO users(username, password, created) values(?,?,?)")
	if err != nil {
		fmt.Printf("stmt prepare err:%v", err)
		return nil, err
	}
	result, err := stmt.Exec("pan", "123456", "2023-06-11")
	if err != nil {
		fmt.Printf("stmt exec err:%v", err)
		return nil, err
	}
	return result, nil
}

func main() {
	db, err := SqliteConn("../demo.db")
	if err != nil {
		fmt.Printf("sqlite conn err:%v\n", err)
	} else {
		fmt.Printf("sqlite conn successful:%v\n", db)
	}
	result, err := Insert(db)
	if err != nil {
		fmt.Printf("insert data err:%v", err)
	} else {
		fmt.Printf("insert data successfule:%v", result)
	}
}
