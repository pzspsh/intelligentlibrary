/*
@File   : update.go
@Author : pan
@Time   : 2023-06-12 16:11:10
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

func Update(db *sql.DB, username string, id int64) (sql.Result, error) {
	stmt, err := db.Prepare("update users set username=? where id=?")
	if err != nil {
		fmt.Printf("stmt prepare update err:%v", err)
		return nil, err
	}
	result, err := stmt.Exec(username, id)
	if err != nil {
		fmt.Printf("stmt exec update err:%v", err)
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
	result, err := Update(db, "go", 3)
	if err != nil {
		fmt.Printf("update data err:%v", err)
	} else {
		fmt.Printf("update data successful:%v", result)
	}
}
