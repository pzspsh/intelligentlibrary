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

// 创建表
func Create(db *sql.DB) {

}

func main() {
	p := &Postgres{
		Host:     "ip",
		Port:     "port",
		Username: "user",
		Password: "pass",
		DB:       "dbname",
	}
	db, err := p.PostgresConn()
	if err != nil {
		fmt.Println("连接失败：", err)
	} else {
		fmt.Println("连接成功：", db)
	}
}
