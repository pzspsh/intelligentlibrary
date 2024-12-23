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
	Email       string `json:"email,omitempty" gorm:"column:email;type:text;default:null"`
}

func (t *Table_demo) Update(db *gorm.DB, id int64) error {
	// db.Model(&Table_demo{}).Where("type = ?", "value").Update("name", "hello") // UPDATE demo SET name='hello' WHERE active=true;
	err := db.Where("id = ?", id).Updates(t)
	if err.Error != nil {
		fmt.Printf("update data err:%v", err.Error)
		return err.Error
	} else {
		fmt.Printf("update data successful")
		return nil
	}
}

func (t *Table_demo) Updates(db gorm.DB, id string) {
	//  实现更新某个字段的执行，其它字段的值不变
	result := db.Model(&Table_demo{}).Where("id = ?", id).Updates(&Table_demo{Email: "new_email@example.com"})
	if result.Error != nil {
		panic(result.Error)
	}
}

func Update(db *gorm.DB, obj, target string, table interface{}) error {
	err := db.Where(obj+" = ?", target).Updates(table)
	if err.Error != nil {
		return err.Error
	}
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
	t := &Table_demo{
		Title:       "hello world1",
		Description: "你好世界1",
		Name:        "hw1",
		Type:        "随便",
	}
	err = t.Update(db, 2)
	if err != nil {
		fmt.Printf("数据插入失败：%v", err)
	} else {
		fmt.Printf("数据插入唱歌")
	}
}
