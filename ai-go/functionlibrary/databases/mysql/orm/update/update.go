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

// 定义一个结构体表示数据表的模型
type User struct {
	ID   uint
	Name string
	Age  int
}

// 定义一个UpdateMap结构体表示需要更新的字段和对应的值
type UpdateMap struct {
	Field string
	Value interface{}
}

// 执行批量更新操作
func BatchUpdate(db *gorm.DB, table string, condition interface{}, updateMapList []UpdateMap) error {
	return db.Table(table).Where(condition).Updates(updateMapList).Error
}

func BatchMain() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动创建数据表
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to migrate database")
	}

	// 批量更新数据
	updateMapList := []UpdateMap{
		{"Name", "Alice"},
		{"Age", 18},
	}
	err = BatchUpdate(db, "users", "id > 0", updateMapList)
	if err != nil {
		panic("failed to batch update")
	}

	fmt.Println("batch update success")
}

func UpdateMain() {
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

func main() {

}
