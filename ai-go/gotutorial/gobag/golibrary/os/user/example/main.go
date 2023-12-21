/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 19:12:14
*/
package main

import (
	"log"
	"os/user"
)

func main() {
	u, _ := user.Current()
	log.Println("用户名：", u.Username)
	log.Println("用户id", u.Uid)
	log.Println("用户主目录：", u.HomeDir)
	log.Println("主组id：", u.Gid)
	// 用户所在的所有的组的id
	s, _ := u.GroupIds()
	log.Println("用户所在的所有组：", s)
}
