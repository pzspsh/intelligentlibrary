/*
@File   : postgres.go
@Author : pan
@Time   : 2023-12-28 15:23:53
*/
package database

type PgOptions struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Db       string `json:"db,omitempty"`
}

func Postgres() {

}
