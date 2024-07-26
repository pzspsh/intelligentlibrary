/*
@File   : main.go
@Author : pan
@Time   : 2024-07-26 14:45:53
*/
package main

import (
	"fmt"
)

// 循环用户密码字典

func main() {
	username := []string{"root", "root1", "root12", "admin", "admin12", "root123", "root@123", "admin123", "admin@123"}
	password := []string{"tech", "root", "root123", "admin", "admin123", "root@123", "admin123", "admin@123", "root1"}
	for _, user := range username {
		for _, pass := range password {
			fmt.Println(user, pass)
		}
	}
}
