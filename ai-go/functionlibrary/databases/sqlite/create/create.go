/*
@File   : create.go
@Author : pan
@Time   : 2023-06-12 16:09:41
*/
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func SqliteConn(dbpath string) (*sql.DB, error) {
	// 打开数据库,dbpath 表示数据库路径
	db, err := sql.Open("sqlite3", dbpath) //打开数据库，如果不存在，则创建
	if err != nil {
		fmt.Printf("sqlite3 open db err:%v", err)
		return nil, err
	}
	// defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Printf("sqlite3 conn err:%v", err)
		return nil, err
	}
	return db, err
}

func Create(db *sql.DB, sqltable string) (sql.Result, error) {
	// 创建表
	result, err := db.Exec(sqltable)
	if err != nil {
		fmt.Printf("create table err:%v\n", err)
		return nil, err
	} else {
		return result, nil
	}
}

func main() {
	sql_table := `
    CREATE TABLE IF NOT EXISTS users(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username VARCHAR(64) NULL,
        password VARCHAR(64) NULL,
        created DATE NULL
    );
    `
	db, err := SqliteConn("../demo.db")
	if err != nil {
		fmt.Printf("sqlite conn err:%v\n", err)
	} else {
		fmt.Printf("sqlite conn successful:%v\n", db)
	}
	result, err := Create(db, sql_table)
	if err != nil {
		fmt.Printf("create table fault:%v\n", err)
	} else {
		fmt.Printf("create table successful:%v\n", result)
	}
}
