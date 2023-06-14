/*
@File   : confconfig.go
@Author : pan
@Time   : 2023-06-09 17:01:57
*/
package main

import (
	"fmt"

	"github.com/Unknwon/goconfig"
)

var cfg *goconfig.ConfigFile

func init() {
	config, err := goconfig.LoadConfigFile("../../config.conf")
	if err != nil {
		fmt.Printf("get config error")
	}
	cfg = config
}

// 全局配置
func GlobalConfig() {
	glob, _ := cfg.GetSection("Config") //读取[Config]全部配置
	fmt.Println(glob)
}

func main() {
	host, _ := cfg.GetValue("Config", "host")
	fmt.Println(host)
	err := cfg.Reload() //重载配置
	if err != nil {
		fmt.Printf("reload config file error: %s", err)
	}
	GlobalConfig()
}
