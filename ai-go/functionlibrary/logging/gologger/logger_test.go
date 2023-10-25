/*
@File   : logger_test.go
@Author : pan
@Time   : 2023-10-25 16:17:25
*/
package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	log := LoggerSet(`..\data\log\log1.txt`, 10, "10mb")
	log2 := LoggerSet(`..\data\log\log2.txt`, 10, "10mb")
	Info("hello world")
	log.Error("你好！！！！！！！")
	log2.Error("你好！！！！！！！")
}
