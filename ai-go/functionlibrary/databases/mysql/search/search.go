/*
@File   : search.go
@Author : pan
@Time   : 2023-06-12 22:14:27
*/
package main

import (
	"database/sql"
	"fmt"
	"time"

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

type Users struct {
	ID       int64     `db:"uid"`
	Username string    `db:"username"`
	Password string    `db:"password"`
	Gender   string    `db:"gender"`
	Created  time.Time `db:"created"`
}

func Search(db *sql.DB) (*Users, error) {
	s := "select * from users where uid = ?"
	var u Users
	err := db.QueryRow(s, 1).Scan(&u.ID, &u.Username, &u.Password, &u.Gender, &u.Created)
	if err != nil {
		return nil, err
	} else {
		return &u, nil
	}
}

func SearchList(db *sql.DB) ([]*Users, error) {
	s := "select * from users"
	stmt, err := db.Query(s)
	var u Users
	var us []*Users
	if err != nil {
		return nil, err
	} else {
		for stmt.Next() {
			stmt.Scan(&u.ID, &u.Username, &u.Password, &u.Gender, &u.Created)
			us = append(us, &u)
		}
	}
	return us, nil
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
	result, err := Search(db)
	if err != nil {
		fmt.Printf("search data err:%v", err)
	} else {
		fmt.Printf("search data successful:%v", result)
	}
	resultlist, err := SearchList(db)
	if err != nil {
		fmt.Printf("search list data err:%v", err)
	} else {
		fmt.Printf("search list data successful:%v", resultlist)
	}
}
