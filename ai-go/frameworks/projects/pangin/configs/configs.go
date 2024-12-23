/*
@File   : configs.go
@Author : pan
@Time   : 2023-06-11 22:26:16
*/
package configs

import (
	"github.com/go-ini/ini"
)

var Conf = new(Configs)

type Configs struct {
	Port         int  `ini:"port"`
	Release      bool `ini:"release"`
	*MySQLConfig `ini:"mysql"`
}

type MySQLConfig struct {
	Username string `ini:"username"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	DB       string `ini:"db"`
}

func Init(file string) error {
	return ini.MapTo(Conf, file)
}
