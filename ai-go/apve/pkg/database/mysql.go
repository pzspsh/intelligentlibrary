/*
@File   : mysql.go
@Author : pan
@Time   : 2023-12-28 15:24:38
*/
package database

type MysqlOption struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Db       string `json:"db,omitempty"`
}

func Mysql() {

}
