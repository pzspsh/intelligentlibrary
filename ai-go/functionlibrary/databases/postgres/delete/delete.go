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

func Delete(db *sql.DB) (sql.Result, error) {
	stmt, err := db.Prepare("delete from table where column1=$1")
	if err != nil {
		return nil, err
	}
	result, err := stmt.Exec("column1_value")
	if err != nil {
		return nil, err
	}
	fmt.Printf("result = %v", result)
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
	result, err := Delete(db)
	if err != nil {
		fmt.Printf("删除失败：%v", err)
	} else {
		fmt.Printf("删除成功：%v", result)
	}
}
