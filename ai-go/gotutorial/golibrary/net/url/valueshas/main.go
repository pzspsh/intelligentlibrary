/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 17:30:16
*/
package main

import (
	"fmt"
	"net/url"
)

func main() {
	v := url.Values{}
	v.Add("cat sounds", "meow")
	v.Add("cat sounds", "mew")
	v.Add("cat sounds", "mau")
	fmt.Println(v.Has("cat sounds"))
	fmt.Println(v.Has("dog sounds"))

}
