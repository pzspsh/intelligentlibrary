/*
@File   : insert.go
@Author : pan
@Time   : 2023-06-12 22:03:35
*/
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	Host     string `json:"host,omitempty"` // 可以用于json解析获取配置对应信息
	Port     string `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Database string `json:"database,omitempty"`
}

func (m *MysqlConfig) MysqlConn() (*sql.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", m.Username, m.Password, m.Host, m.Port, m.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Insert(db *sql.DB, insertdata string) (int64, error) {
	stmt, err := db.Exec(insertdata, "stu001", "man", "admin@123", "2023-06-12")
	if err != nil {
		return -1, err
	}
	result, err := stmt.LastInsertId()
	if err != nil {
		return -1, err
	}
	return result, nil
}

func main() {
	m := &MysqlConfig{
		Host:     "ip",
		Port:     "port",
		Username: "user",
		Password: "pass",
		Database: "dbname",
	}
	db, err := m.MysqlConn()
	if err != nil {
		fmt.Printf("mysql conn err:%v", err)
	} else {
		fmt.Printf("mysql conn successful:%v", db)
	}
	insertdata := "insert into users(username, gender, password, created)values(?, ?, ?, ?)"
	result, err := Insert(db, insertdata)
	if err != nil {
		fmt.Printf("insert data err:%v", err)
	} else {
		fmt.Printf("insert data successful:%v", result)
	}
}
