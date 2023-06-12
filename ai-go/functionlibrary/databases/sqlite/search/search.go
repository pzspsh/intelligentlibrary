/*
@File   : search.go
@Author : pan
@Time   : 2023-06-12 16:10:46
*/
package main

import (
	"database/sql"
	"fmt"
	"time"

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

func Search(db *sql.DB) ([][]interface{}, error) {
	result := make([][]interface{}, 0)
	data := make([]interface{}, 0)
	rows, err := db.Query("select * from users")
	if err != nil {
		fmt.Printf("search query data err:%v", err)
		return result, err
	}
	var uid int64
	var username string
	var password string
	var created time.Time
	for rows.Next() {
		err = rows.Scan(&uid, &username, &password, &created)
		if err != nil {
			fmt.Printf("rows scan err:%v", err)
			return result, err
		}
		data = append(data, uid)
		data = append(data, username)
		data = append(data, password)
		data = append(data, created)
		result = append(result, data)
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
	result, err := Search(db)
	if err != nil {
		fmt.Printf("result search err:%v", err)
	} else {
		fmt.Printf("result search successful:%v", result)
	}
}
