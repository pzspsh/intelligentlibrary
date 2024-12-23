/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 10:11:31
*/
package main

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
)

func main() {
	value, err := json.Marshal("hello world")
	if err != nil {
		fmt.Println(err)
	}

	fnv128 := fnv.New128a()
	_, err = fnv128.Write(value)
	if err != nil {
		fmt.Println(err)
	}
	hash := fnv128.Sum([]byte{})
	fmt.Println(hash)
}
