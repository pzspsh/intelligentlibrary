/*
@File   : update.go
@Author : pan
@Time   : 2023-06-13 10:37:14
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

func Update(db *sql.DB) (int64, error) {
	s := "update users set username=?,password=? where uid=?"
	r, err := db.Exec(s, "pan01", "admin@1234", "3")
	if err != nil {
		return -1, err
	} else {
		i, _ := r.RowsAffected()
		return i, nil
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
	result, err := Update(db)
	if err != nil {
		fmt.Printf("update data err:%v", err)
	} else {
		fmt.Printf("update data successful:%v", result)
	}
}
