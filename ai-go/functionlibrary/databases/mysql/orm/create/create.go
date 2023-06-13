/*
@File   : create.go
@Author : pan
@Time   : 2023-06-13 11:53:42
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

// 设置表名称
func (StudentInfo) TableName() string {
	return "studentinfo"
}

func (s *StudentInfo) Create(db *gorm.DB) {
	dbMig := db.Migrator()
	if !dbMig.HasTable(&s) { // 查看表是否存在，不存在这则执行创建表操作
		db.AutoMigrate(&s)
		fmt.Println("创建表 studentinfo成功。。。")
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
	s := &StudentInfo{}
	// 创建数据库表
	s.Create(db)
}
