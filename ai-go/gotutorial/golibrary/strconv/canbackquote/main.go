/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 22:32:49
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.CanBackquote("Fran & Freddie's Diner ☺"))
	fmt.Println(strconv.CanBackquote("`can't backquote this`"))
}
