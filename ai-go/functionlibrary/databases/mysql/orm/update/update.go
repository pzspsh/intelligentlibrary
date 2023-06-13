/*
@File   : update.go
@Author : pan
@Time   : 2023-06-13 11:56:16
*/
package main

import (
	"fmt"
	"time"

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
	ID       int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Number   string    `json:"number,omitempty" orm:"column:title;type:varchar(32);default:null"`
	Title    string    `json:"title,omitempty" orm:"column:title;type:varchar(255);default:null"`
	Username string    `json:"username,omitempty" orm:"column:username;type:varchar(255);default:null"`
	Password string    `json:"password,omitempty" orm:"column:password;type:varchar(255);default:null"`
	Desc     string    `json:"description,omitempty" orm:"column:description;type:text;default:null"`
	Phone    string    `json:"phone,omitempty" orm:"column:phone;type:varchar(255);default:null"`
	Updated  time.Time `json:"updated,omitempty" orm:"column:updated;default:null"`
	Age      int       `json:"age:omitempty" orm:"column:age,default:null"`
}

func (s *StudentInfo) Update(db *gorm.DB, id string) error {
	// err := db.Model(s).Where("password = ?", "admin123").Updates(&StudentInfo{Number: "100002", Username: "pan123"})
	// if err.Error != nil {
	// 	return err.Error
	// } else {
	// 	return nil
	// }
	err := db.Where("id = ?", id).Updates(s)
	if err.Error != nil {
		fmt.Printf("update data err:%v", err.Error)
		return err.Error
	} else {
		fmt.Printf("update data successful")
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
		Number: "10009",
	}
	err = s.Update(db, "3")
	if err != nil {
		fmt.Printf("update data err:%v", err)
	} else {
		fmt.Printf("update data successful.")
	}
}
