/*
@File   : delete.go
@Author : pan
@Time   : 2023-06-12 16:10:02
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

func Delete(db *sql.DB, goal string) (sql.Result, error) {
	stmt, err := db.Prepare("delete from users where id = ?")
	if err != nil {
		return nil, err
	}
	result, err := stmt.Exec(goal)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteTable(db *sql.DB) (sql.Result, error) {
	result, err := db.Exec("drop table userinfo")
	if err != nil {
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
	result, err := Delete(db, "goalvalue")
	if err != nil {
		fmt.Printf("delete data err:%v", err)
	} else {
		fmt.Printf("delete data successful:%v", result)
	}
}
