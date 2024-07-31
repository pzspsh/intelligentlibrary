/*
@File   : main.go
@Author : pan
@Time   : 2024-07-26 14:45:53
*/
package main

import "fmt"

// 循环用户密码字典

type UserPass struct {
	User string
	Pass string
}

func main() {
	username := []string{"root", "root1", "root12", "admin", "admin12", "root123", "root@123", "admin123", "admin@123"}
	password := []string{"tech", "root", "root123", "admin", "admin123", "root@123", "admin123", "admin@123", "root1"}
	dicts := UserPassDict(username, password)
	for _, dict := range dicts {
		fmt.Println(dict.User, dict.Pass)
	}
}

func UserPassDict(users []string, passs []string) []UserPass {
	dicts := []UserPass{}
	for _, user := range users {
		for _, pass := range passs {
			dict := UserPass{user, pass}
			dicts = append(dicts, dict)
		}
	}
	return dicts
}

func OneUserPass(users []string, passs []string) {

}
