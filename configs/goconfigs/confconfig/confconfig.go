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

func ParseConf(filepath string) (*goconfig.ConfigFile, error) {
	config, err := goconfig.LoadConfigFile("../../config.conf")
	if err != nil {
		fmt.Printf("get config error")
		return nil, err
	}
	glob, err := config.GetSection("Config")
	if err != nil {
		fmt.Printf("config GetSection err:%v", err)
		return nil, err
	}
	fmt.Println(glob)
	host, err := config.GetValue("Config", "host")
	if err != nil {
		fmt.Printf("config GetValue err:%v", err)
		return nil, err
	}
	fmt.Println(host)
	return config, nil
	// err = config.Reload() //重载配置
	// if err != nil {
	// 	fmt.Printf("reload config file error: %s", err)
	// }
}

func WriteConf(filepath string) {

}

func main() {
	config, err := ParseConf("../../config.conf")
	if err != nil {
		fmt.Printf("parse conf err:%v", err)
	}
	fmt.Println(config)
}
