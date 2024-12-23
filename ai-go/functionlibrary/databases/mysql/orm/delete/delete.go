/*
@File   : delete.go
@Author : pan
@Time   : 2023-06-13 11:55:11
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

func (s *StudentInfo) Delete(db *gorm.DB, id string) error {
	// db.Delete(s, []int{1,2,3}) // DELETE FROM studentinfo WHERE id IN (1,2,3);

	// db.Delete(s, "10") // DELETE FROM studentinfo WHERE id = 10;

	// db.Where("username LIKE ?", "%jinzhu%").Delete(s) // DELETE from studentinfo where username LIKE "%jinzhu%";

	// db.Delete(&[]StudentInfo{{username: "jinzhu1"}, {username: "jinzhu2"}}).Error // gorm.ErrMissingWhereClause

	// db.Exec("delete from studentinfo") // 删除表

	// 永久删除
	// db.Unscoped().Delete(&StudentInfo{}) // DELETE FROM studentinfo WHERE id=10;

	err := db.Delete(s, id)
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
	s := &StudentInfo{}
	err = s.Delete(db, "id_value")
	if err != nil {
		fmt.Printf("delete data err:%v", err)
	} else {
		fmt.Printf("delete data successful.")
	}
}
