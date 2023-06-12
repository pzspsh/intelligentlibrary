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

type Table_demo struct {
	ID          int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Title       string `json:"title,omitempty" gorm:"column:title;type:varchar(255);default:null"`
	Name        string `json:"name,omitempty" gorm:"column:name;type:varchar(255);default:null"` // `gorm:"uniqueIndex:name"`
	Description string `json:"description,omitempty" gorm:"column:description;type:text;default:null"`
	Type        string `json:"type,omitempty" gorm:"column:type;type:text;default:null"`
}

func Delete(db *gorm.DB) error {
	// db.Delete(&Table_demo{}, 11) // DELETE FROM demo WHERE id = 10;
	// db.Where("name = ?", "hw").Delete(&Table_demo{}) // DELETE from demo where name = "hw";
	// db.Delete(&Table_demo{}, []int{1, 2, 3}) // DELETE FROM demo WHERE id IN (1,2,3);
	// db.Where("hw LIKE ?", "%hw%").Delete(&Table_demo{}) // DELETE from deme where name LIKE "%hw%"; // 批量删除
	// db.Delete(&Table_demo{}, "name LIKE ?", "%hw%") // DELETE from demo where name LIKE "%hw%";
	// db.Unscoped().Delete(&Table_demo{}) // DELETE FROM orders WHERE id=10; // 永久删除

	return nil
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
	err = Delete(db)
	if err != nil {
		fmt.Printf("delete date err:%v", err)
	} else {
		fmt.Printf("delete data successful.s")
	}
}
