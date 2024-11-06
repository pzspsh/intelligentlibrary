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

// 查询单个
func Search(db *gorm.DB, column, searchv string) (bool, *Table_demo) {
	demo := new(Table_demo)
	err := db.Where(column+" = ?", searchv).First(demo)
	if err.Error != nil {
		fmt.Printf("search data err:%v", err.Error)
		return false, nil
	} else {
		fmt.Println("search data successfule.")
		return true, demo
	}
}

// 查询所有
func SearchList(db *gorm.DB) (bool, *[]Table_demo) {
	tabledemo := new([]Table_demo)
	err := db.Find(tabledemo)
	if err.Error != nil {
		fmt.Printf("search list err:%v", err.Error)
		return false, nil
	} else {
		return true, tabledemo
	}
}

func Select(db *gorm.DB, obj interface{}, selectobj, target string) (interface{}, error) {
	err := db.Where(selectobj+" = ?", target).First(obj).Error
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func Search1(target, number string, db *gorm.DB) {
	tabledemo := new(Table_demo)
	if err := db.Where(target+" = ?", number).First(tabledemo); err.RowsAffected == 0 {
		fmt.Println("AAAAA", err.RowsAffected)
	} else if err.Error != nil {
		fmt.Println("err: ", err.Error)
	} else {
		fmt.Println("aaa: ", err)
	}
}


// HasData 函数用于判断表是否有数据
func HasData(db *gorm.DB, model interface{}) bool {
	var count int64
	db.Model(model).Count(&count)
	return count > 0
}

// 多条件查询
func Search2(value1, value2 string, db *gorm.DB) error {
	var err error
	tabledemo := new(Table_demo)
	if err = db.Where("title = ? and name = ?", value1, value2).First(tabledemo).Error; err != nil {
		return err
	}
	return err
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
	ok, tabledata := Search(db, "column", "search value")
	if ok {
		fmt.Println(tabledata)
	}
	ok, tablelist := SearchList(db)
	if ok {
		fmt.Println(tablelist)
	}
}
