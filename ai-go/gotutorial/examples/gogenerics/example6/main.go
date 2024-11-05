/*
@File   : main.go
@Author : pan
@Time   : 2024-11-05 16:26:43
*/
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Demo1 struct {
	Fiel1 string
	Fiel2 string
}

type Demo2 struct {
	Fiel1 string
	Fiel2 string
}

type Demo3 struct {
	Fiel1 string
	Fiel2 string
}

type PgConfig struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Database string `json:"database,omitempty"`
}

func (c *PgConfig) PGConn() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai", c.Host, c.Username, c.Password, c.Database, c.Port)
	sqlDB, _ := sql.Open("pgx", dsn)
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn:                 sqlDB,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				IgnoreRecordNotFoundError: true,
			}),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

type LdkGen[T []*Demo1] func(data T) T

// func Insert[T *Demo | *Demo2 | *Demo3](data []T, db *gorm.DB) error {
func Insert[T any](data []T, db *gorm.DB) error {
	var err error
	if len(data) > 0 {
		nlen := len(data)
		if nlen > 1000 {
			splits := nlen / 1000
			if nlen%1000 != 0 {
				splits++
			}
			for i := 0; i < splits; i++ {
				start := i * 1000
				end := start + 1000
				if end > len(data) {
					end = len(data)
				}
				part := data[start:end]
				if err = db.Create(&part).Error; err != nil {
					return err
				}
			}
		} else {
			if err = db.Create(&data).Error; err != nil {
				return err
			}
		}
	}
	return err
}

func main() {
	var err error
	var demolist = []Demo1{}
	demo1 := Demo1{
		Fiel1: "hello",
		Fiel2: "world",
	}
	demolist = append(demolist, demo1)
	dbconfig := PgConfig{}
	db, _ := dbconfig.PGConn()
	if err = Insert(demolist, db); err != nil {
		fmt.Println(err)
	}
}
