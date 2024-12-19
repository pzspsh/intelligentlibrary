/*
@File   : main.go
@Author : pan
@Time   : 2024-12-17 16:31:27
*/
package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

type Postgres struct {
	Host     string
	Port     string
	Username string
	Password string
	DB       string
}

type QueryData struct {
	Id        uint
	TableName string
	Select    string
}

func (p *Postgres) PostgresConn() (*sql.DB, error) {
	dbdsn := fmt.Sprintf(`postgres://%v:%v@%v:%v/%v?sslmode=disable`, p.Username, p.Password, p.Host, p.Port, p.DB)
	db, err := sql.Open("postgres", dbdsn)
	if err != nil {
		fmt.Println("failed to open a db conn:", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	// db.Close()
	return db, nil
}

func GetData(q QueryData, db *sql.DB) (map[string]string, error) {
	var err error
	var result = make(map[string]string)
	query := fmt.Sprintf(`select %s FROM %s where id=$1`, q.Select, q.TableName)
	row, err := db.Query(query, q.Id)
	if err != nil {
		return result, err
	}
	defer row.Close()
	columns, err := row.Columns()
	if err != nil {
		return result, err
	}
	values := make([]interface{}, len(columns))
	valueptrs := make([]interface{}, len(columns))
	for i := range columns {
		valueptrs[i] = &values[i]
	}
	for row.Next() {
		if err = row.Scan(valueptrs...); err != nil {
			return result, err
		}
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			result[col] = fmt.Sprintf("%v", v)
		}
	}
	return result, err
}

func UpdateData(db *sql.DB, table string, id int, updates map[string]string) (sql.Result, error) {
	// 构建更新字段的 SQL 片段和参数切片
	var setClause []string
	var args []interface{}
	args = append(args, id) // 首先添加 ID 参数
	for field, value := range updates {
		setClause = append(setClause, fmt.Sprintf("%s = $%d", field, len(args)+1))
		args = append(args, value)
	}
	// 构建完整的 SQL 更新语句
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", table, strings.Join(setClause, ", "), 1)
	// 执行更新操作
	fmt.Println(query)
	fmt.Println(args...)
	result, err := db.Exec(query, args...)
	fmt.Printf("result = %d", result)
	return result, err
}

func main() {
	var err error
	p := &Postgres{
		Host:     "",
		Port:     "",
		Username: "",
		Password: "",
		DB:       "",
	}
	db, err := p.PostgresConn()
	if err != nil {
		fmt.Println("数据库连接失败", err)
	}
	q := QueryData{
		Id:        1,
		TableName: "",
		Select:    "",
	}
	result, err := GetData(q, db)
	if err != nil {
		fmt.Println("BBBBBBBB", err)
	}
	fmt.Println(result)
	UpdateData(db, "", 1, result)
}
