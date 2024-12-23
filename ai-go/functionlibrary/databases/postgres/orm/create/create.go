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
	Name        string `json:"name,omitempty" gorm:"column:name;type:varchar(255);default:null"`
	Description string `json:"description,omitempty" gorm:"column:description;type:text;default:null"`
	Type        string `json:"type,omitempty" gorm:"column:type;type:text;default:null"`
}

// 设置表名称
func (Table_demo) TableName() string {
	return "demo"
}

func Create(db *gorm.DB) {
	dbMig := db.Migrator()
	if !dbMig.HasTable(&Table_demo{}) { // 查看表是否存在，不存在这则执行创建表操作
		db.AutoMigrate(&Table_demo{})
		fmt.Println("创建表 demo成功。。。")
	}
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
	// 创建表
	Create(db)
}
