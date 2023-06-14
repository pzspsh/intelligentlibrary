/*
@File   : tomlconfig.go
@Author : pan
@Time   : 2023-06-09 16:58:17
*/
package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
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
	viper.AddConfigPath("D:/GoProjects/src/intelligentlibrary/configs")
	viper.SetConfigType("toml")
	viper.SetConfigName("configs")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	config := &TomlConfig{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		_ = viper.Unmarshal(config)
	})
	return config, nil
}

func main() {
	data, err := ParseToml("../../")
	if err != nil {
		fmt.Printf("parse toml err:%v", err)
	}
	fmt.Println(data)
	// fmt.Println(config)
}
