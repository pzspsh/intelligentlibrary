/*
@File   : connect.go
@Author : pan
@Time   : 2023-06-05 14:16:23
*/
package main

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string `json:"host,omitempty"` // 可以用于json解析获取配置对应信息
	Port     string `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Database string `json:"database,omitempty"`
}

func (c *DBConfig) PostgresConn() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai", c.Host, c.Username, c.Password, c.Database, c.Port)
	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		fmt.Printf("sql open err:%v", err)
		return nil, err
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn:                 sqlDB,
		PreferSimpleProtocol: true,
	}))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	c := &DBConfig{
		Host:     "ip",
		Port:     "port",
		Username: "user",
		Password: "pass",
		Database: "dbname",
	}
	db, err := c.PostgresConn()
	if err != nil {
		fmt.Printf("postgres conn err:%v", err)
	} else {
		fmt.Printf("postgres conn successful:%v", db)
	}
}
