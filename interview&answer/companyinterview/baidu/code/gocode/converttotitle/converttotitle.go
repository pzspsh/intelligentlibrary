/*
@File   : converttotitle.go
@Author : pan
@Time   : 2023-05-24 11:06:13
*/
package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(28))
}

// Excel表列名称
func convertToTitle(n int) string {
	str := ""

	for n > 0 {
		n--
		str = string(byte(n%26)+'A') + str
		n /= 26
	}
	return str
}
