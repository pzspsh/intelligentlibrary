/*
@File   : utils.go
@Author : pan
@Time   : 2023-09-18 14:17:05
*/
package test

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomNumber(prefix string, count int) string {
	sb := strings.Builder{}
	sb.Grow(count)
	for i := 0; i < count; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	number := fmt.Sprintf("%v-%v-%v", prefix, time.Now().Unix(), sb.String())
	return number
}
