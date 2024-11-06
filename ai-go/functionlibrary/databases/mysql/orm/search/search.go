/*
@File   : search.go
@Author : pan
@Time   : 2023-06-13 11:55:58
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

func Search(db *gorm.DB, goal string) (*StudentInfo, error) {
	// 获取第一条记录（主键升序）
	// db.First(&StudentInfo{}) 	// SELECT * FROM studentinfo ORDER BY id LIMIT 1;

	// LIKE
	// db.Where("username LIKE ?", "%jin%").Find(&StudentInfo{}) // SELECT * FROM studentinfo WHERE username LIKE '%jin%';

	// IN
	// db.Where("username IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&StudentInfo{}) // SELECT * FROM studentinfo WHERE username IN ('jinzhu','jinzhu 2');

	// AND
	// db.Where("username = ? AND age >= ?", "jinzhu", "22").Find(&StudentInfo{}) // SELECT * FROM studentinfo WHERE username = 'jinzhu' AND age >= 22;

	// Time
	// db.Where("updated > ?", lastWeek).Find(&StudentInfo{}) // SELECT * FROM studentinfo WHERE updated > '2000-01-01 00:00:00';

	// Map
	// db.Where(map[string]interface{}{"username": "jinzhu", "age": 20}).Find(&StudentInfo{}) // SELECT * FROM studentinfo WHERE usersname = "jinzhu" AND age = 20;

	// Slice of primary keys
	// db.Where([]int64{20, 21, 22}).Find(&StudentInfo{}) // SELECT * FROM studentinfo WHERE id IN (20, 21, 22);

	// Struct
	// db.Where(&StudentInfo{username: "jinzhu", Age: 20}).First(&StudentInfo{}) // SELECT * FROM studentinfo WHERE username = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;

	// db.Where(&StudentInfo{username: "jinzhu"}, "username", "Age").Find(&users) // SELECT * FROM studentinfo WHERE username = "jinzhu" AND age = 0;

	// db.Where("password = ?", "admin").Or("password = ?", "super_admin").Find(&users) // SELECT * FROM studentinfo WHERE password = 'admin' OR password = 'super_admin';

	var studentinfo StudentInfo
	err := db.Where("username = ?", goal).First(&studentinfo)
	if err.Error != nil {
		return nil, err.Error
	} else {
		return &studentinfo, nil
	}
}

// 查询所有
func SearchList(db *gorm.DB) (bool, *[]StudentInfo) {
	student := new([]StudentInfo)
	err := db.Find(student)
	if err.Error != nil {
		fmt.Printf("search list err:%v", err.Error)
		return false, nil
	} else {
		return true, student
	}
}

func SearchRun() {
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
	result, err := Search(db, "stu01")
	if err != nil {
		fmt.Printf("search data err:%v", err)
	} else {
		fmt.Printf("search data successful:%v", result)
	}
}

type YourModel struct {
	ID uint `gorm:"primaryKey"`
	// 其他字段
}

func checkIfExists(db *gorm.DB, ids []uint) {
	var existingIDs []YourModel

	// 批量查询存在的 ID
	db.Select("ID").Where("ID IN ?", ids).Find(&existingIDs)

	// 将查询结果转换为 ID 切片
	existingIDSet := make(map[uint]struct{}, len(existingIDs))
	for _, model := range existingIDs {
		existingIDSet[model.ID] = struct{}{}
	}

	// 检查哪些 ID 存在，哪些不存在
	for _, id := range ids {
		if _, exists := existingIDSet[id]; exists {
			// ID 存在
			fmt.Println("ID", id, "exists")
		} else {
			// ID 不存在
			fmt.Println("ID", id, "does not exist")
		}
	}
}

func SearcRun2() {
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 确保数据库连接成功
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database connection")
	}
	defer sqlDB.Close()

	// 示例 ID 列表
	ids := []uint{1, 2, 3, 4, 5 /* 其他 ID... */, 100}

	// 检查是否存在
	checkIfExists(db, ids)
}

func main() {

}
