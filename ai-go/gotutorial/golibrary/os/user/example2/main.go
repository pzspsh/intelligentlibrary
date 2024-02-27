/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 19:12:14
*/
package main

import (
	"fmt"
	"os/user"
	"reflect"
)

func main() {
	fmt.Println("== 测试 Current 正常情况 ==")
	if u, err := user.Current(); err == nil {
		fmt.Println("用户ID: " + u.Uid)
		fmt.Println("主组ID: " + u.Gid)
		fmt.Println("用户名: " + u.Username)
		fmt.Println("主组名: " + u.Name)
		fmt.Println("家目录: " + u.HomeDir)
	}

	fmt.Println("== 测试 Lookup 正常情况 ==")
	if u, err := user.Lookup("root"); err == nil {
		fmt.Println("用户ID: " + u.Uid)
		fmt.Println("主组ID: " + u.Gid)
		fmt.Println("用户名: " + u.Username)
		fmt.Println("主组名: " + u.Name)
		fmt.Println("家目录: " + u.HomeDir)
	}
	fmt.Println("== 测试 Lookup 异常情况 ==")
	if _, err := user.Lookup("roo"); err == nil {
	} else {
		fmt.Println("错误信息: " + err.Error())
		fmt.Print("错误类型: ")
		fmt.Println(reflect.TypeOf(err))
	}

	fmt.Println("== 测试 LookupId 正常情况 ==")
	if u, err := user.LookupId("0"); err == nil {
		fmt.Println("用户ID: " + u.Uid)
		fmt.Println("主组ID: " + u.Gid)
		fmt.Println("用户名: " + u.Username)
		fmt.Println("主组名: " + u.Name)
		fmt.Println("家目录: " + u.HomeDir)
	}
	fmt.Println("== 测试 LookupId 异常情况 ==")
	if _, err := user.LookupId("10000"); err == nil {
	} else {
		fmt.Println("错误信息: " + err.Error())
		fmt.Print("错误类型: ")
		fmt.Println(reflect.TypeOf(err))
	}
}
