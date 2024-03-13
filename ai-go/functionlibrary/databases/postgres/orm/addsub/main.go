/*
@File   : main.go
@Author : pan
@Time   : 2024-03-13 10:41:23
*/
package main

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Host     string
	Port     string
	Username string
	Password string
	DB       string
}

func (p *Postgres) PostgresConn() (*gorm.DB, error) {
	dbdsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", p.Username, p.Password, p.Host, p.Port, p.DB)
	sqldb, err := sql.Open("postgres", dbdsn)
	if err != nil {
		fmt.Println("failed to open a db conn:", err)
		return nil, err
	}
	err = sqldb.Ping()
	if err != nil {
		return nil, err
	}
	defer sqldb.Close()
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn:                 sqldb,
		PreferSimpleProtocol: true,
	}))
	if err != nil {
		return nil, err
	}
	return db, nil
}

type Table_demo struct {
	ID          int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Title       string `json:"title,omitempty" gorm:"column:title;type:varchar(255);default:null"`
	Name        string `json:"name,omitempty" gorm:"column:name;type:varchar(255);default:null"`
	Description string `json:"description,omitempty" gorm:"column:description;type:text;default:null"`
	Counter     int    `json:"counter,omitempty" gorm:"column:counter;default:null"`
}

// 设置表名称
func (Table_demo) TableName() string {
	return "demo"
}

func (t *Table_demo) AddCount(db *gorm.DB, count int) error {
	// 对主键ID为1进行更新
	// err := db.Model(t).Where("id = ?", 1).UpdateColumn("counter", gorm.Expr("counter + ?", count))
	err := db.Model(t).Where("id = ?", 1).Update("counter", gorm.Expr("counter + ?", count))
	if err.Error != nil {
		fmt.Println("update counter error:", err)
		return err.Error
	}
	return nil
}

func (t *Table_demo) SubCount(db *gorm.DB, count int) error {
	// 对主键ID为1进行更新
	// err := db.Model(t).Where("id = ?", 1).UpdateColumn("counter", gorm.Expr("counter - ?", count))
	err := db.Model(t).Where("id = ?", 1).Update("counter", gorm.Expr("counter - ?", count))
	if err.Error != nil {
		fmt.Println("update counter error:", err)
		return err.Error
	}
	return nil
}

func main() {
	var pg = Postgres{
		Host:     "localhost",
		Port:     "5432",
		Username: "root",
		Password: "root",
		DB:       "test",
	}

	db, err := pg.PostgresConn()
	if err != nil {
		fmt.Println("pg connect error:", err)
	}
	var tb *Table_demo
	err = tb.AddCount(db, 5)
	if err != nil {
		fmt.Println("add count error:", err)
	}
	err = tb.SubCount(db, 3)
	if err != nil {
		fmt.Println("sub count error:", err)
	}
}
