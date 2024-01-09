/*
@File   : main.go
@Author : pan
@Time   : 2023-06-07 14:27:58
*/
package main

import (
	"fmt"
	"os"
	"pangin/configs"
	"pangin/pkg/api"
	"pangin/pkg/database"
	"pangin/pkg/models"
)

const defaultconfig = "./conf/configs.ini"

func main() {
	configfile := defaultconfig
	if len(os.Args) > 2 {
		fmt.Println("user specified config file", os.Args[1])
		configfile = os.Args[1]
	} else {
		fmt.Println("no configuration file was specied, using default config file")
	}
	if err := configs.Init(configfile); err != nil {
		fmt.Println("error reading config file:", err)
		return
	}
	// 创建数据库
	// sql: CREATE DATABASE pangin;
	// 连接数据库
	err := database.InitMysql(configs.Conf.MySQLConfig)
	if err != nil {
		fmt.Println("init mysql failed err:", err)
		return
	}
	defer database.Close() // 程序退出关闭数据库连接
	// 模型绑定
	database.DB.AutoMigrate(&models.PanDemo{})
	// 注册路由
	r := api.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", configs.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
