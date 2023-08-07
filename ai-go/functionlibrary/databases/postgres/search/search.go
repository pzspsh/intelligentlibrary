package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Postgres struct {
	Host     string
	Port     string
	Username string
	Password string
	DB       string
}

func (p *Postgres) PostgresConn() (*sql.DB, error) {
	dbdsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", p.Username, p.Password, p.Host, p.Port, p.DB)
	db, err := sql.Open("postgres", dbdsn)
	if err != nil {
		fmt.Println("failed to open a db conn:", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	db.Close()
	return db, nil
}

type Cv struct {
	column1 string
	column2 string
}

func (c *Cv) Search(db *sql.DB) (sql.Result, error) {
	rows, err := db.Query("select column1, column2 from table where column3=$1", "column3_value")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&c.column1, &c.column2)
		if err != nil {
			return nil, err
		}
		fmt.Printf("column1 = %v, column2 = %v", c.column1, c.column2)
	}
	return nil, nil
}

// select 查询字段为空 rows.Scan()会报错,
// 解决办法：
//     1、那个字段会有Null，就将此字段对应的Golang结构体的字段类型改为指针类型
//     2、 SQL语句格式, 例如：SELECT name, COALESCE(name, '') FROM person 或 SELECT id, IFNULL(id, 0) FROM person

func SearchCount(db *sql.DB) (int64, error) {
	var count int64
	err := db.QueryRow("select count(*) from table where column1=$1", "column1_value").Scan(&count)
	if err != nil || count == 0 {
		fmt.Printf("count could not query err:%v", err)
		return count, err
	}
	// db.Close()
	return count, err
}

func main() {
	p := &Postgres{
		Host:     "",
		Port:     "",
		Username: "",
		Password: "",
		DB:       "",
	}
	db, err := p.PostgresConn()
	if err != nil {
		fmt.Println("连接失败：", err)
	} else {
		fmt.Println("连接成功：", db)
	}
	cv := Cv{}
	_, err = cv.Search(db)
	if err != nil {
		fmt.Printf("查询失败：%v", err)
	} else {
		fmt.Println("查询成功")
	}
}
