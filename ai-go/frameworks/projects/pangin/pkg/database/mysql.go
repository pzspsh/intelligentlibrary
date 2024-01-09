/*
@File   : mysql.go
@Author : pan
@Time   : 2024-01-08 17:20:45
*/
package database

import (
	"fmt"
	"pangin/configs"

	// "github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMysql(cfg *configs.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%v:%v@(%v:%v)/%v", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return
}

func Close() {
	// DB.Close()
}
