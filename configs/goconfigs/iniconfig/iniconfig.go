/*
@File   : iniconfig.go
@Author : pan
@Time   : 2023-06-09 16:57:17
*/
package main

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type IniConfig struct {
	Config Config
	Auth   Auth
}

type Config struct {
	Host    string `ini:"host,omitempty"`
	Port    int    `ini:"port,omitempty"`
	Timeout int    `ini:"timeout,omitempty"`
}

type Auth struct {
	Username string `ini:"username,omitempty"`
	Password string `ini:"password,omitempty"`
}

// 读取inis配置文件
func LoadConfigfile(filepath string) (*IniConfig, error) {
	cfg, err := ini.Load(filepath)
	if err != nil {
		fmt.Printf("cfg load err:%v", err)
		return nil, err
	}
	fmt.Println("config:", cfg.Section("Config").Key("host").String()) //
	config := new(IniConfig)
	err = cfg.MapTo(config)
	if err != nil {
		fmt.Printf("ini config parser err:%v", err)
		return nil, err
	}
	fmt.Println(config)
	return config, nil
}

// 将结构体映射成配置文件
type ConfigWrite struct {
	Title  string `ini:"title"`
	Name   string `ini:"name,omitempty" comment:"name's hello"`
	Status int    `ini:"status"`
}

func WriteIni(filepath string) {
	config := &ConfigWrite{
		Title:  "hello world",
		Name:   "pansor",
		Status: 200,
	}
	cfg := ini.Empty()
	err := cfg.Section("WriteIni").ReflectFrom(config)
	if err != nil {
		fmt.Printf("ini empty err:%v", err)
	}
	// 写入ini文件
	err = cfg.SaveTo(filepath)
	if err != nil {
		fmt.Printf("failed to save file err:%v", err)
	}
}

func main() {
	// LoadConfigfile("../../config.ini")
	WriteIni("../../writeconfig.ini")
}
