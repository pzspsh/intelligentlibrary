/*
@File   : mysql.go
@Author : pan
@Time   : 2024-01-08 17:20:45
*/
package database

import (
	"pangin/configs"

	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

func InitMysql(cfg *configs.MySQLConfig) (err error) {
	dsn := "root:root@(xxx.xxx:3306)/gormDemo?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}
