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

func Search(db *sql.DB) {

}

func main() {

}
