/*
@File   : insert.go
@Author : pan
@Time   : 2023-06-13 11:55:36
*/
package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConfig struct {
	Host     string `json:"host,omitempty"` // 可以用于json解析获取配置对应信息
	Port     string `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Database string `json:"database,omitempty"`
}

func (m *MysqlConfig) MysqlConn() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.Port, m.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	} else {
		return db, err
	}
}

type StudentInfo struct {
	ID       int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Number   string `json:"number,omitempty" orm:"column:title;type:varchar(32);default:null"`
	Title    string `json:"title,omitempty" orm:"column:title;type:varchar(255);default:null"`
	Username string `json:"username,omitempty" orm:"column:username;type:varchar(255);default:null"`
	Password string `json:"password,omitempty" orm:"column:password;type:varchar(255);default:null"`
	Desc     string `json:"description,omitempty" orm:"column:description;type:text;default:null"`
	Phone    string `json:"phone,omitempty" orm:"column:phone;type:varchar(255);default:null"`
}

func (s *StudentInfo) Insert(db *gorm.DB) error {
	// slist := []*StudentInfo{
	// 	&StudentInfo{Number: "100005"},
	// 	&StudentInfo{Number: "100006s"},
	// }
	// err := db.Create(slist)
	err := db.Create(s)
	if err.Error != nil {
		return err.Error
	} else {
		return nil
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
		fmt.Printf("mysql conn err:%v", m)
	} else {
		fmt.Printf("mysql conn successful:%v", db)
	}
	s := &StudentInfo{
		Number:   "100001",
		Title:    "hello world!",
		Username: "stu01",
		Password: "123456",
		Desc:     "该学生。。。",
		Phone:    "12312342123",
	}
	err = s.Insert(db)
	if err != nil {
		fmt.Printf("insert data err:%v", err)
	} else {
		fmt.Printf("insert data successful.")
	}
}
