/*
@File   : tomlconfig.go
@Author : pan
@Time   : 2023-06-09 16:58:17
*/
package main

import (
	"fmt"
	"os"

	// "github.com/fsnotify/fsnotify"
	"github.com/BurntSushi/toml"
	"github.com/spf13/viper"
)

type TomlConfig struct {
	Config Config `toml:"config"`
	Auth   Auth   `toml:"Auth"`
}

type Config struct {
	Host    string `toml:"host,omitempty"`
	Port    int    `tome:"port,omitempty"`
	Timeout int    `tome:"timeout,omitempty"`
}

type Auth struct {
	Username string `toml:"username,omitempty"`
	Password string `toml:"password,omitempty"`
}

func ParseToml(filepath string) (*TomlConfig, error) {
	viper.AddConfigPath(filepath) // "D:/GoProjects/src/intelligentlibrary/configs"
	viper.SetConfigType("toml")
	viper.SetConfigName("configs")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	config := &TomlConfig{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}
	// viper.WatchConfig()
	// viper.OnConfigChange(func(in fsnotify.Event) {
	// 	_ = viper.Unmarshal(config)
	// })
	return config, nil
}

func WriteToml(filename string, config *TomlConfig) error {
	// 判断filename是否存在，不存在则创建再执行一下操作(Check whether filename exists. If no, create a file and perform operations)
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	if err := toml.NewEncoder(file).Encode(config); err != nil {
		return err
	}
	return nil
}

func main() {
	data, err := ParseToml("../../")
	if err != nil {
		fmt.Printf("parse toml err:%v", err)
	}
	fmt.Println(data)
	err = WriteToml("../../writeconfig.toml", data)
	if err != nil {
		fmt.Printf("write toml err:%v", err)
	}
}
