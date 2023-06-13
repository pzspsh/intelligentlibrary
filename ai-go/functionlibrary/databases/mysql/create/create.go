/*
@File   : create.go
@Author : pan
@Time   : 2023-06-12 21:56:30
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

func Create(db *sql.DB, newtable string) (sql.Result, error) {
	result, err := db.Exec(newtable)
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
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
	//定义一个创建数据库表的变量,后续进行exec的调用.
	createTable := "CREATE TABLE `users` (`uid` int(10) NOT NULL AUTO_INCREMENT,`username` varchar(64) NOT NULL DEFAULT '1', `gender` varchar(32) DEFAULT NULL,`password` varchar(64) DEFAULT NULL,`created` date DEFAULT NULL,PRIMARY KEY (`uid`)) ;"
	result, err := Create(db, createTable)
	if err != nil {
		fmt.Printf("create table err:%v", err)
	} else {
		fmt.Printf("create table successful:%v", result)
	}
}
