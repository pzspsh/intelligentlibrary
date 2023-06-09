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

func Insert(db *sql.DB, table string) (sql.Result, error) {
	stmt, err := db.Prepare(fmt.Sprintf("insert into %v(column1,column2) values($1,&2)", table))
	if err != nil {
		return nil, err
	}
	result, err := stmt.Exec("column1_value", "column2_value")
	if err != nil {
		return nil, err
	}
	fmt.Printf("result = %d", result)
	return result, err
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
	result, err := Insert(db, "table")
	if err != nil {
		fmt.Printf("更新失败:%v", err)
	} else {
		fmt.Printf("更新成功：%v", result)
	}
}
