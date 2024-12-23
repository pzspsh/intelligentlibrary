/*
@File   : example.go
@Author : pan
@Time   : 2023-05-17 14:14:44
*/
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Number(prefix string) string {
	var count int = 6
	sb := strings.Builder{}
	sb.Grow(count)
	for i := 0; i < count; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	number := fmt.Sprintf("%v-%v-%v", prefix, time.Now().Unix(), sb.String())
	return number
}

func main() {
	number := Number("PAN")
	fmt.Println(number)
}
